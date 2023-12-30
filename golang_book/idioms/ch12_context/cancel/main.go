package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"time"
)

func main() {
	ss := slowServer()
	defer ss.Close()
	fs := fastServer()
	defer fs.Close()
	
	ctx := context.Background()
	callBoth(ctx, os.Args[1], ss.URL, fs.URL)
}

func slowServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
			time.Sleep(2 * time.Second)
			w.Write([]byte("slow ok"))
	}))
	return s
}

func fastServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, 
		r *http.Request) {
			if r.URL.Query().Get("error") == "true" {
				w.Write([]byte("error"))
				return
			}
			w.Write([]byte("ok"))
		}))
	return s
}

var client = http.Client{}

func callBoth(ctx context.Context, errVal, slowUrl, fastUrl string) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := callServer(ctx, "slow", slowUrl)
		if err != nil {
			cancel()
		}
	}()
	go func() {
		defer wg.Done()
		err := callServer(ctx, "fast", fastUrl+"?error="+errVal)
		if err != nil {
			cancel()
		}
	}()
	wg.Wait()
	fmt.Println("done both")
}

func callServer(ctx context.Context, label, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(label, "req err: ", err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(label, "resp err: ", err)
		return err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result := string(data)
	if result != "" {
		fmt.Println(label, "result: ", result)
	}
	if result == "error" {
		fmt.Println("cancelling from", label)
		return fmt.Errorf("error happened")
	}
	return nil
}

func longRunningThingManager(ctx context.Context, data string) (string, error) {
	type wrapper struct {
		result string
		err 	error
	}
	ch := make(chan wrapper, 1)
	go func() {
		res, err := longRunningThing(ctx, data)
		ch <- wrapper{res, err}
	}()
	select {
	case data := <-ch:
		return data.result, data.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}