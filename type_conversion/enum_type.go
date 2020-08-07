package type_conversion

// see https://chai2010.cn/advanced-go-programming-book/ch2-cgo/ch2-03-cgo-types.html

/*
enum e {
	ONE,
	TWO
};
enum e e_val = ONE;
*/
import "C"
import "fmt"

func GetSetEnumTypeVariable()  {
	// ------ get value ------
	// print: 0(ONE)
	fmt.Println(C.e_val)

	// ------ set value ------
	C.e_val = C.TWO
	// print: 1(TWO)
	fmt.Println(C.e_val)
}