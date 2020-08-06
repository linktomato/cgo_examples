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

func GetSetIntArrayTypeVariables() {
	// ------ get value ------
	// 1.克隆一份数组到Go内存空间中
	// 除了byte类型的数组之外，cgo没有提供类似GoBytes的辅助函数，
	// 但是原理都是一样的，我们可以手动实现内存拷贝。
	var arr []int32
	for i := 0; i < 3; i++ {
		arr = append(arr, int32(C.arr_int[i]))
	}
	// [1 1 1]
	fmt.Println(arr)

	// 2.复用C内存空间的数组（参考GetSetCharArrayTypeVariables函数实现)

	// ------ set value ------
	// 1.克隆一份数组到C内存空间中
	var cArr = [3]C.int{0}
	for i := 0; i < 3; i++ {
		cArr[i] = C.int(2)
	}
	C.arr_int = cArr
	// print: [2 2 2]
	fmt.Println(C.arr_int)
	// 2.设置元素值（参考GetSetCharArrayTypeVariables函数实现)
	// 3.复用Go内存中的数组（参考GetSetCharArrayTypeVariables函数实现)
}

func GetSetCharArrayTypeVariables() {
	// ------ get value ------
	// 1.克隆一份数组到Go内存空间中
	var arrByte []byte = C.GoBytes(unsafe.Pointer(&C.arr_char[0]), 3)
	// print: [1 1 1]
	fmt.Println(arrByte)

	// 2.复用C内存空间的数组
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
	// 1.克隆一份数组到C内存空间中
	arrByte = []byte{2, 2, 2}
	C.arr_char = *(*[3]C.char)(C.CBytes(arrByte))
	// print: [2 2 2]
	fmt.Println(C.arr_char)

	// 2.设置元素值
	for i := 0; i < 3; i++ {
		C.arr_char[i] = C.char(3)
	}
	// print: [3 3 3]
	fmt.Println(C.arr_char)

	// 3.复用Go内存中的数组
	// todo 需遵守以下规范：
	// Go code may pass a Go pointer to C provided the Go memory to which
	// it points does not contain any Go pointers. The C code must preserve
	// this property:
	// it must not store any Go pointers in Go memory, even temporarily.
	// see https://golang.org/cmd/cgo/#hdr-C_references_to_Go
	arrByte = []byte{4, 4, 4}
	var arrByteHdr = (*reflect.SliceHeader)(unsafe.Pointer(&arrByte))
	C.arr_char = *(*[3]C.char)(unsafe.Pointer(arrByteHdr.Data))
	// print: [4 4 4]
	fmt.Println(C.arr_char)
}
