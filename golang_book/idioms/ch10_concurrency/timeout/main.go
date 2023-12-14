package main

import (
	"errors"
	"time"
)

func main() {
	
}

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err := doWork()
		close(done)
	}()
	select {
	case <- done:
		return result, err
	case <- time.After(2 * time.Second):
		return 0, errors.New("timed out")
	}
} 