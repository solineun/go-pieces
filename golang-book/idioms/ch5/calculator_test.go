package ch5

import "testing"

func Test_calculate(t *testing.T) {
	expr := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"5"},
	}
	if got, _ := Calculate(expr[0]); got != 5 {
		t.Errorf("Calculate(%v) = %v; want 5", expr[0], got)
	}
}
