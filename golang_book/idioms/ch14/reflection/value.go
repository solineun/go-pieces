package main

import (
	"fmt"
	"reflect"
)

var stringType = reflect.TypeOf((*string)(nil)).Elem()
var stringSliceType = reflect.TypeOf([]string(nil))

func reflectNew() {
	ssv := reflect.MakeSlice(stringSliceType, 0, 10)
	sv := reflect.New(stringType).Elem()
	fmt.Printf("%+v, %T, %+v, %T\n", ssv, ssv, sv, sv)
	sv.SetString("hello")
	ssv = reflect.Append(ssv, sv)
	ss := ssv.Interface().([]string)
	fmt.Println(ss)
}

func hasNoValue(i interface{}) bool {
	iv := reflect.ValueOf(i)
	if !iv.IsValid() {
		return true
	}
	switch iv.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Func, reflect.Interface:
		return iv.IsNil()
	default:
		return false
	}
}