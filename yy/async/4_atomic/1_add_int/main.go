package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var count int64

	for i := 0; i < 1e5; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(count)
}