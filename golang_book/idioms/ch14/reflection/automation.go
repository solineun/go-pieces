package main

import (
	"fmt"
	"reflect"
	"time"
)

func MakeTimedFunction(f interface{}) interface{} {
	ft := reflect.TypeOf(f)
	fv := reflect.ValueOf(f)
	wrapperF := reflect.MakeFunc(ft, func(in []reflect.Value) []reflect.Value {
		start := time.Now()
		out := fv.Call(in)
		end := time.Now()
		fmt.Println(end.Sub(start))
		return out
	})
	return wrapperF.Interface()
}

func timeMe(a int) int {
	time.Sleep(time.Duration(a) * time.Second)
	result := a * 2
	return result
}
	
// func main() {
// 	timed:= MakeTimedFunction(timeMe).(func(int) int)
// 	fmt.Println(timed(2))
// }