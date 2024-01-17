package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	A int		`myTag:"value"`
	B string 	`myTag:"value2"`
}

func intpt() {
	type myint int
	var myx myint = 123
	mt := reflect.TypeOf(myx)
	fmt.Println(mt.Name(), mt.Kind())
	
	var x int
	xpt := reflect.TypeOf(&x)
	fmt.Println(xpt.Name())
	fmt.Println(xpt.Kind())
	fmt.Println(xpt.Elem().Name())
	fmt.Println(xpt.Elem().Kind())
}

func strctype() {
	var f Foo
	ft := reflect.TypeOf(f)
	fmt.Println(ft.Name(), ft.Kind())
	for i := 0; i < ft.NumField(); i++ {
		curField := ft.Field(i)
		fmt.Println(curField.Name, curField.Type.Name(), curField.Tag.Get("myTag"))
	}
}

func val() {
	v := 123
	vv := reflect.ValueOf(&v)
	fmt.Println(vv.Type(), vv.Kind(), vv.Elem())
	vv.Elem().SetInt(20)
	fmt.Println(vv.Type(), vv.Kind(), vv.Elem())
}