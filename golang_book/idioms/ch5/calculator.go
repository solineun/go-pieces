package ch5

import (
	"errors"
	"strconv"
)

type opFuncType func(int, int) (int, error)

func add(a, b int) (int, error) {return a + b, nil}
func sub(a, b int) (int, error) {return a - b, nil}
func mul(a, b int) (int, error) {return a * b, nil}
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

var opMap = map[string]opFuncType {
	"+" : add,
	"-" : sub,
	"*" : mul,
	"/" : div,
}

func Calculate(input []string) (int, error) {
	if len(input) != 3 {
		return 0, errors.New("invalid expression")
	}
	n1, err := strconv.Atoi(input[0])
	if err != nil {
		return 0, err
	}
	op := input[1]
	n2, err := strconv.Atoi(input[2])
	if err != nil {
		return 0, err
	}
	opFunc, ok := opMap[op]
	if !ok {
		return 0, errors.New("invalid operation")
	}
	return opFunc(n1, n2)
}



