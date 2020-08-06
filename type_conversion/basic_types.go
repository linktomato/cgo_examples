package type_conversion

/*
 #include <stdlib.h>
#include <stdio.h>
#include <float.h>

// ------------基本数据类型------------
char v_char = 1 << 7;
signed char v_schar = 1 << 7;
unsigned char v_uchar = 1 << 7;
//在cgo中short类型的长度固定为2个字节
short v_short = 1 << 15;
unsigned short v_ushort = 1 << 15;
//在cgo中int类型的长度固定为4个字节
int v_int = 1 << 31;
unsigned int v_uint = 1 << 31;
long v_long = 1 << 31;
unsigned long v_ulong = 1 << 31;
long long int v_longlong = 1LL << 63;
unsigned long long int v_ulonglong = 1LL << 63;
float v_float = FLT_MAX;
double v_double = DBL_MAX;
//在64位系统中，size_t类型的长度为8个字节
size_t v_size_t = 1LL << 63;
*/
import "C"
import "fmt"

// 基本数据类型的转化
// see https://chai2010.cn/advanced-go-programming-book/ch2-cgo/ch2-03-cgo-types.html
func GetSetBasicTypeVariables() {
	// ------ get value ------
	var vChar = byte(C.v_char)
	var vSChar = int8(C.v_schar)
	var vUChar = uint8(C.v_uchar)
	var vShort = int16(C.v_short)
	var vUShort = uint16(C.v_ushort)
	var vInt = int32(C.v_int)
	var vUInt = uint32(C.v_uint)
	var vLong = int32(C.v_long)
	var vULong = uint32(C.v_ulong)
	var vLongLong = int64(C.v_longlong)
	var vULongLong = uint64(C.v_ulonglong)
	var vFloat = float32(C.v_float)
	var vDouble = float64(C.v_double)
	var vSizeT = uint(C.v_size_t)

	// ------ set value ------
	C.v_char = C.char(0)
	C.v_schar = C.schar(0)
	C.v_uchar = C.uchar(0)
	C.v_short = C.short(0)
	C.v_ushort = C.ushort(0)
	C.v_int = C.int(0)
	C.v_uint = C.uint(0)
	C.v_long = C.long(0)
	C.v_ulong = C.ulong(0)
	C.v_longlong = C.longlong(0)
	C.v_ulonglong = C.ulonglong(0)
	C.v_float = C.float(0)
	C.v_double = C.double(0)
	C.v_size_t = C.size_t(0)

	fmt.Println(vChar, C.v_char)
	fmt.Println(vSChar, C.v_schar)
	fmt.Println(vUChar, C.v_uchar)
	fmt.Println(vShort, C.v_short)
	fmt.Println(vUShort, C.v_ushort)
	fmt.Println(vInt, C.v_int)
	fmt.Println(vUInt, C.v_uint)
	fmt.Println(vLong, C.v_long)
	fmt.Println(vULong, C.v_ulong)
	fmt.Println(vLongLong, C.v_longlong)
	fmt.Println(vULongLong, C.v_ulonglong)
	fmt.Println(vFloat, C.v_float)
	fmt.Println(vDouble, C.v_double)
	fmt.Println(vSizeT, C.v_size_t)
}
