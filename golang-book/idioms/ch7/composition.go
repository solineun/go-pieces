package ch7

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
}

type Employee struct {
	Person
	ID string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.String(), e.ID)
}

type Manager struct	{
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	return nil
}

var m Manager = Manager{}
var e Employee = m.Employee

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)	
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)	
}