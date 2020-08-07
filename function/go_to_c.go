package function

/*
// see https://pubs.opengroup.org/onlinepubs/009695399/basedefs/errno.h.html
#include <errno.h>

int sum(int num1, int num2){
	return num1 + num2;
}

int div(int num1, int num2){
	if (num2 == 0) {
		errno = EINVAL;
		return 0;
	}
	return num1 / num2;
}

void do_something(){
	errno = EACCES;
}

*/
import "C"
import (
	"fmt"
)

func InvokeFunction() {
	var sum = C.sum(1, 2)
	// print: 3
	fmt.Println(sum)

	sum, err := C.sum(1, 2)
	// print: 3 <nil>
	fmt.Println(sum, err)
}

func InvokeFunctionGetErr() {
	var div, err = C.div(1, 0)
	// print: 0
	fmt.Println(div, err)
	// print: Invalid argument.
	fmt.Println(err)
}

func InvokeFunctionRetVoid() {
	var ret, err = C.do_something()
	// print: []
	fmt.Println(ret)
	// print: Permission denied.
	fmt.Println(err)
}
