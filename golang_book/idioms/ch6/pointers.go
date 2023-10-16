package ch6

import "fmt"

func Run() {
	y := 111
	py := &y
	ppy := &py
	fmt.Println(py)
	update(ppy)
	fmt.Println(py)
}

func update(p **int) {
	x := 12
	*p = &x
	fmt.Println(*p)
}