// go 1.18.1
// пользователь вводит кольчество элементов массива, затем сами элементы
// вывести количество цифр в массиве

package main

import (
	"fmt"
	book "mymodule/golang-book/idioms/ch7"
	"mymodule/yy"
	"time"
)

func main() {
	
}

func yySet() {
	set := yy.NewSet[string]()
	set.Add("a", "b", "b")
	fmt.Println(set.Contains("a"))
}

func yyClient() {
	t := yy.Timeout(time.Millisecond * 2)
	c := yy.NewClient("dasd", 31, t)
	fmt.Println(c.String())
}

func runCounter() {
	var c book.Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())
}

func runTree() {
	var it *book.IntTree
	it = it.Insert(10)
	fmt.Println(it)

	it = it.Insert(20)
	fmt.Println(it)
	it = it.Insert(5)
	fmt.Println(it)

	fmt.Println(it.Contains(10))
	fmt.Println(it.Contains(20))
	fmt.Println(it.Contains(2))
}

func runAdder() {
	myAdder := book.Adder{Start: 10}
	fmt.Println(myAdder.AddTo(5))

	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	f2 := book.Adder.AddTo
	fmt.Println(f2(myAdder, 15))
}

func runTests() {
	fmt.Println(test()) //2
	fmt.Println(anotherTest()) //1
}

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}
func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}

func runComposition() {
	m := book.Manager{
		Employee: book.Employee{
			Person: book.Person{
				FirstName: "Allen",
				LastName: "Ilurpen",
				Age: 23,
			},
			ID: "412987",
		},
		Reports: []book.Employee{},
	}
	fmt.Println(m.ID)
	fmt.Println(m.Description())

	o := book.Outer{
		Inner: book.Inner{
			A: 10,
		},
		S: "Hi",
	}
	fmt.Println(o.Double())
}