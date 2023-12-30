package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()
	b := time.Now().Equal(t1)
	fmt.Println(b)
}