package book

import (
	"fmt"
)

func main()  {
	a, b, s := 4.0 / 2, 32132 * 42452, "Hello"
	fmt.Printf("%T\n", a)
	fmt.Println(s[1])
	fmt.Println(true && false)

	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(len(s))
}