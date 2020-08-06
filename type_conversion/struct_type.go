package type_conversion

/*
typedef struct {
	char *name;
	int age;
	// 用Go关键字命名的变量名，在Go中以该变量名前加下划线访问，此处为"_type"。
	// 如果存在成员变量名为"_type"，则此变量会被Go忽略。
	// Member's name conflicts with Go's keyword,
	// Go will access it by using name "_type".
	// If there is some one named "_type" then this
    // member will be ignored in Go.
	char *type;
	char *_type;
} programer;

programer me = {
	"tomato",
	30,
	"ignored",
	"primary engineer"
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
	// 如果成员变量的名字和
	// print: primary engineer
	fmt.Println(C.GoString(C.me._type))

	// ------ set value ------
	// 修改结构体成员变量值
	// set struct member value
	C.me.name = C.CString("old tomato")
	C.me.age = C.int(60)
	// print: old tomato
	fmt.Println(C.GoString(C.me.name))
	// print: 60
	fmt.Println(C.me.age)

	// 创建新的结构体变量
	// set a new struct variable
	var me = C.programer{
		name:  C.CString("super tomato"),
		age:   C.int(120),
	}
	C.me = me
	// print: super tomato
	fmt.Println(C.GoString(C.me.name))
	// print: 120
	fmt.Println(C.me.age)
}
