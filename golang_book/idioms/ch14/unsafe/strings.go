package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func strings() {
	s := "hello"
	sHdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Println(sHdr.Len)
	for i := 0; i < sHdr.Len; i++ {
		bp := *(*byte)(unsafe.Pointer(sHdr.Data + uintptr(i)))
		fmt.Println(string(bp))
	}
	fmt.Println()
	runtime.KeepAlive(s)
}

