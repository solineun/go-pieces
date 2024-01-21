package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func slices() {
	s := []int{10, 20, 30}
	sHdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len)
	fmt.Println(sHdr.Cap)
	intByteSize := unsafe.Sizeof(s[0])
	fmt.Println(intByteSize)
	for i := 0; i < sHdr.Len; i++ {
		intVal := *(*int)(unsafe.Pointer(sHdr.Data + intByteSize*uintptr(i)))
		fmt.Println(intVal)
	}
	runtime.KeepAlive(s)
}