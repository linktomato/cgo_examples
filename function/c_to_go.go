package function

/*

extern int add(int num1,int num2);

static int go_add(int num1,int num2){
	return add(num1,num2);
}

// Using //export in a file places a restriction on the preamble:
// since it is copied into two different C output files,
// it must not contain any definitions, only declarations.
//
// 含有 export 导出函数的Go文件会生成两个 .c 文件，如果在 preamble(import "C"的上文)
// 中写任何函数的定义将会出现多次定义函数的错误 (multiple definition of `impl_func')。
//
// see: https://golang.org/cmd/cgo/#hdr-C_references_to_Go
// error: multiple definition of `impl_func'
//
// void impl_func(){}
*/
import "C"
import "fmt"

//export add
func add(num1 C.int, num2 C.int) C.int {
	return num1 + num2
}

func CGoAdd() {
	fmt.Println(C.go_add(C.int(1), C.int(2)))
}