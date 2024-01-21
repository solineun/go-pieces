package main

import (
	"encoding/binary"
	"math/bits"
	"unsafe"
)

func main() {
	
}

type Data struct {
	Value uint32 // 4 bytes
	Label [10]byte // 10 bytes
	Active bool // 1 byte
	// go rounds this data 1 byte more 
}

// example of data were read from net
// [0 132 95 237 80 104 111 110 101 0 0 0 0 0 1 0]

// secure way to read that kind of data
func DataFromBytes(b [16]byte) Data {
	d := Data{}
	d.Value = binary.BigEndian.Uint32(b[:4])
	copy(d.Label[:], b[4:14])
	d.Active = b[14] != 0
	return d
}

func DataFromBytesUnsafe(b [16]byte) Data {
	data := *(*Data)(unsafe.Pointer(&b))
	if isLE {
		data.Value = bits.ReverseBytes32(data.Value)
	}
	return data
}

var isLE bool

func init() {
	var x uint16 = 0xFF00
	xb := *(*[2]byte)(unsafe.Pointer(&x))
	isLE = (xb[0] == 0x00)
}

func BytesFromData(d Data) [16]byte {
	out := [16]byte{}
	binary.BigEndian.PutUint32(out[:4], d.Value)
	copy(out[4:14], d.Label[:])
	if d.Active {
		out[14] = 1
	}
	return out
}

func BytesFromDataUnsafe(d Data) [16]byte {
	if isLE {
		d.Value = bits.ReverseBytes32(d.Value)
	}
	out := *(*[16]byte)(unsafe.Pointer(&d))
	return out
}