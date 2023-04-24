package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str := "hello world"
	stringToSliceHeader := *(*reflect.SliceHeader)(unsafe.Pointer(&str))
	byteSlice := *(*[]byte)(unsafe.Pointer(&stringToSliceHeader))
	fmt.Printf("%#v\n", byteSlice)
}
