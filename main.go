package main

import "fmt"

type Doer struct {
	num int
}

func (d *Doer) method() {
	d.num *= 2
}

func (d Doer) do() {
	d.method()
	fmt.Println(d.num)
}

func main() {
	d := Doer{8}
	d.do()
}