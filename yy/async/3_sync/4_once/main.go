package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	for i := 0; i < 1e6; i++ {
		go func ()  {
			once.Do(func() {
				fmt.Println("once")
			})
		}()
	}
}