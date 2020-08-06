package type_conversion

/*
#include <stdlib.h>
#include <stdio.h>
#include <float.h>
#include <string.h>

char arr_char[3] = {1,1,1};
int arr_int[3] = {1,1,1};
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func GetSetArrayTypeVariables() {
	// ------ get value ------
	// 克隆一份数组到Go内存空间中
	var arrByte []byte = C.GoBytes(unsafe.Pointer(&C.arr_char[0]), 3)
	// print: [1 1 1]
	fmt.Println(arrByte)

	// 复用C内存空间的数组
	// 需要保证内存不会被C端提前释放掉。
	// todo 验证用完后该内存会不会自动回收
	var arrByte2 []byte
	var arrByte2Hdr = (*reflect.SliceHeader)(unsafe.Pointer(&arrByte2))
	arrByte2Hdr.Data = uintptr(unsafe.Pointer(&C.arr_char[0]))
	arrByte2Hdr.Len = 3
	arrByte2Hdr.Cap = 3
	// print: [1 1 1]
	fmt.Println(arrByte2)

	// ------ set value ------
	// 克隆一份数组到C内存空间中
	arrByte = []byte{2, 2, 2}
	C.arr_char = *(*[3]C.char)(C.CBytes(arrByte))
	// print: [2 2 2]
	fmt.Println(C.arr_char)
	// 复用Go内存中的数组
	// todo 需遵守以下规范：
	// Go code may pass a Go pointer to C provided the Go memory to which
	// it points does not contain any Go pointers. The C code must preserve
	// this property:
	// it must not store any Go pointers in Go memory, even temporarily.
	// see https://golang.org/cmd/cgo/#hdr-C_references_to_Go
	arrByte = []byte{3, 3, 3}
	var arrByteHdr = (*reflect.SliceHeader)(unsafe.Pointer(&arrByte))
	C.arr_char = *(*[3]C.char)(unsafe.Pointer(arrByteHdr.Data))
	// print: [2 2 2]
	fmt.Println(C.arr_char)
}
