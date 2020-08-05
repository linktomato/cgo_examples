package go_ref_to_c

/*
 int add(int foo,int bar){
	return foo + bar;
 }
*/
import "C"

func CAdd(foo int, bar int) int {
	// 错误示例：
	// 调用c函数时及获取返回值时，需要将参数进行强制转化类型。
	// 在cgo中，c语言的int类型是32位的，与go语言的int32类型对应。
	// go语言的int类型在32位和64位系统中分别是32位和64位，无法直接传给c函数。
	// int类型无法直接传
	// wrong : return C.add(foo, bar)

	return int(C.add(C.int(foo), C.int(bar)))
}
