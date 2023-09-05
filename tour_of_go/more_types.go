package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var p *int

func foo() {
	i := 42
	p = &i
	fmt.Println(p)
	*p = 21
	fmt.Println(i)

	
	v := Vertex{2, 4}
	p := &v
	p.X = 1e9
	fmt.Println(p)
}