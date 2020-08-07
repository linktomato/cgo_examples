package type_conversion

/*
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

char *str = "hello golang";
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

// 字符串类型的转化
// conversion of string type
// see https://chai2010.cn/advanced-go-programming-book/ch2-cgo/ch2-03-cgo-types.html
func GetSetStringTypeVariables() {
	// ------ get value ------
	// 将C字符串(字节数组)克隆一份到Go内存空间中
	// print: hello golang
	fmt.Println(C.GoString(C.str))

	// 复用C内存字符串(字节数组)
	// 因为Go语言的字符串是只读的，用户需要自己保证Go字符串在使用期间，
	// 底层对应的C字符串内容不会发生变化、内存不会被提前释放掉。
	// todo 验证用完后该内存会不会自动回收
	var s string
	var sHdr = (*reflect.StringHeader)(unsafe.Pointer(&s))
	sHdr.Data = uintptr(unsafe.Pointer(C.str))
	sHdr.Len = int(C.strlen(C.str))
	// print: hello golang
	fmt.Println(s)

	// ------ set value ------
	C.str = C.CString("hello c")
	// CString函数从C端堆内存中分配了内存，调用方有责任释放这部分内存。
	// The C string is allocated in the C heap using malloc.
	// It is the caller's responsibility to arrange for it to be
	// freed, such as by calling C.free.
	// see https://golang.org/cmd/cgo/#hdr-C_references_to_Go
	defer C.free(unsafe.Pointer(C.str))
	// print: hello c
	fmt.Println(C.GoString(C.str))
	// print: hello
	fmt.Println(C.GoStringN(C.str, 5))
}
