package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	return
}

func Same(first, second *tree.Tree) bool {
	fch := make(chan int)
	sch := make(chan int)

	go func() {
		defer close(fch)
		Walk(first, fch)
	}()

	go func() {
		defer close(sch)
		Walk(second, sch)
	}()

	var fsum int
	var ssum int

	for f := range fch {
		fsum += f	
	}

	for s := range sch {
		ssum += s	
	}
	return fsum == ssum
}

func main() {
	ch := make(chan int, 10)
	go func() {
		defer close(ch)
		Walk(tree.New(1), ch)
	}()

	fmt.Println(Same(tree.New(1), tree.New(2)))
}