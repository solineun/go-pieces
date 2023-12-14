package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	pg := New(10)
	http.HandleFunc("/req", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("too many requests"))
		}
	})
	go spawn()
	http.ListenAndServe("localhost:8080", nil)
}

type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no capacity")
	}
}

func doLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func spawn() {
	for i := 0; i < 15; i++ {
		go func() {
			r, err := http.Get("http://localhost:8080/req")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(r.Status)
		}()
	}
}