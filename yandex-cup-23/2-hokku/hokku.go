package main

import (
	"fmt"
	"os"
)

func main() {
	var input string
	fmt.Fscan(os.Stdin, &input)
	fmt.Println(Format(input))
}

func Format(input string) string {
	
}