package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var count int64

	for i := 0; i < 1e5; i++ {
		go func(i int) {
			atomic.SwapInt64(&count, int64(i))
		}(i)
	}

	time.Sleep(time.Second * 3 / 2)
	fmt.Println(count)
}