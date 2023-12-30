package main

import (
	"context"
	"fmt"
	"net/http"
)

func main() {
	
}

func MiddleWare(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data := r.FormValue("data")
	res, err := logic(ctx, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(res))
}

type ServiceCaller struct {
	client *http.Client
}

func (sc ServiceCaller) callAnotherservice(ctx context.Context, data string) (string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		"http://example.com?data="+data,
		nil,
	)
	if err != nil {
		return "", err
	}
	req = req.WithContext(ctx)
	resp, err := sc.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if c := resp.StatusCode; c != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d", c)
	}
	id, err := processResp(ctx, resp.Body)
	return id, err
}