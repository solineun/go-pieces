package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	z := 1.0
	eps := 0.001
	for math.Abs(x - z*z ) > eps {
		z -= (z*z - x) / (2*z)
	}
	return fmt.Sprint(z)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func boo() {
	defer fmt.Println(sqrt(2))
	defer fmt.Println("------------")

	fmt.Println(pow(2, 3, 20))
	fmt.Println(pow(3, 3, 20))
}