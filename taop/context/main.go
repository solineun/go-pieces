package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		err := cancelReq(ctx)
		if err != nil {
			cancel()
		}
	}()
	
	doReq(ctx, "https://google.com")
}

func cancelReq(ctx context.Context) error {
	time.Sleep(100 * time.Millisecond)
	return fmt.Errorf("fail request")
}

func doReq(ctx context.Context, reqStr string) {
	req, _ := http.NewRequest(http.MethodGet, reqStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err) 
	}

	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("response completed, status code, %d\n", res.StatusCode)
	case <-ctx.Done():
		fmt.Println("request too long")
	}
}