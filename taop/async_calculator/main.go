package main

import "fmt"

func main() {
	eCh := make(chan expression)
	rCh := make(chan string)

	go func(exprs []expression) {
		defer close(eCh)
		for _, e := range exprs {
			eCh <- e
		}
	}(exprs)

	go func() {
		defer close(rCh)
		calculate(eCh, rCh)
	}()

	for r := range rCh {
		fmt.Println(r)
	}
}

type expression struct {
	val1 int
	val2 int
	operation string
}

func calculate(eCh chan expression, rCh chan string) {
	for e := range eCh {
		switch e.operation {
		case "+": 
			rCh <- fmt.Sprintf("%d %s %d = %d", e.val1, e.operation, e.val2, e.val1 + e.val2)
		case "*":
			rCh <- fmt.Sprintf("%d %s %d = %d", e.val1, e.operation, e.val2, e.val1 * e.val2)
		case "-":
			rCh <- fmt.Sprintf("%d %s %d = %d", e.val1, e.operation, e.val2, e.val1 - e.val2)
		case "/":
			rCh <- fmt.Sprintf("%d %s %d = %d", e.val1, e.operation, e.val2, e.val1 / e.val2)
		default:
			rCh <- fmt.Sprintf("wrong operation: %s", e.operation)
		}
	}
}

var exprs = []expression{
	{
		2,
		2,
		"+",
	},
	{
		4,
		12,
		"*",
	},
	 {
		2,
		2,
		"-",
	},
	{
		14,
		3,
		"/",
	},
	{
		32,
		23,
		"?",
	},
}