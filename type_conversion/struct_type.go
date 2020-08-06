package type_conversion

/*
typedef struct {
	char *name;
	int age;
} programer;

programer me = {
	"tomato",30
};
*/
import "C"
import "fmt"

func GetSetStructTypeVariables() {
	// ------ get value ------
	// print: tomato
	fmt.Println(C.GoString(C.me.name))
	// print: 30
	fmt.Println(C.me.age)

	
	// ------ set value ------
	// 修改结构体字段值
	// set field value
	C.me.name = C.CString("old tomato")
	C.me.age = C.int(60)
	// print: old tomato
	fmt.Println(C.GoString(C.me.name))
	// print: 60
	fmt.Println(C.me.age)

	// 创建新的结构体变量
	// set a new struct variable
	var me = C.programer{
		name: C.CString("super tomato"),
		age:  C.int(120),
	}
	C.me = me
	// print: super tomato
	fmt.Println(C.GoString(C.me.name))
	// print: 120
	fmt.Println(C.me.age)
}
