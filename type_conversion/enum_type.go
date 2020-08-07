package type_conversion

// see https://chai2010.cn/advanced-go-programming-book/ch2-cgo/ch2-03-cgo-types.html

/*
enum e {
	ONE,
	TWO
};
enum e v_e = ONE;
*/
import "C"
import "fmt"

func GetSetEnumTypeVariable()  {
	// ------ get value ------
	// print: 0(ONE)
	fmt.Println(C.v_e)

	// ------ set value ------
	C.v_e = C.TWO
	// print: 1(TWO)
	fmt.Println(C.v_e)
}