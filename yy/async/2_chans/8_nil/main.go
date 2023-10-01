package main

import "fmt"

func main() {
	var a chan int 

	select {
	case i := <-a:
		fmt.Println(i)
	default: 
	fmt.Println("default")
	}
	close(a)
}