package main

import "reflect"

func filter(slice interface{}, filter interface{}) interface{} {
	sv := reflect.ValueOf(slice)
	fv := reflect.ValueOf(filter)
	sliceLen := sv.Len()
	out := reflect.MakeSlice(sv.Type(), 0, sliceLen)
	for i := 0; i < sliceLen; i++ {
		curval := sv.Index(i)
		values := fv.Call([]reflect.Value{curval})
		if values[0].Bool() {
			out = reflect.Append(out, curval)
		}
	}
	return out.Interface()
}