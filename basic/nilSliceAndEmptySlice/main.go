// reference to https://mp.weixin.qq.com/s?__biz=Mzg5NDY2MDk4Mw==&mid=2247486369&idx=1&sn=cc950a5fa0bf84fa917e4dbf1bff29b0&source=41#wechat_redirect
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var nilSlice []int
	emptySlice1 := make([]int, 0)
	emptySlice2 := make([]int, 0)

	fmt.Printf("nilSlice pointer:%+v, emptySlice1 pointer:%+v, emptySlice2 pointer:%+v, \n", *(*reflect.SliceHeader)(unsafe.Pointer(&nilSlice)), *(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice1)), *(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice2)))
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&nilSlice))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice1))).Data)
	fmt.Printf("%v\n", (*(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice1))).Data == (*(*reflect.SliceHeader)(unsafe.Pointer(&emptySlice2))).Data)
}
