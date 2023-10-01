package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("hi from anon func")
	}()
	time.Sleep(time.Second*2)

	go printHello()

	var p printer
	go p.printHello()
}

func printHello() {
	fmt.Println("hi from named func")
}

type printer struct {}

func (p printer) printHello() {
	fmt.Println("hi from struct method")
}