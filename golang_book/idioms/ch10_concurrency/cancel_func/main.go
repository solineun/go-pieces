package main

import (
	"fmt"
	"sync"
)

func main() {
	ch, cancel := countTo(10)
	for i := range ch {
		if i > 5 {
			break
		}
	}
	cancel()
}

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <- done:
				return
			default:
				ch <- i
			}
		}
		close(ch)
	}()
	return ch, cancel
}

type User struct{
	Name string
}



func foo() {
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	max := 0
	for i := 1000; i >=0; i--{
		wg.Add(1)
		go func(i int){
			mu.Lock()
			defer func ()  {
				mu.Unlock()
				wg.Done()
			}()
			if i%2==0{
				max = max(i, max)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println(max)
}