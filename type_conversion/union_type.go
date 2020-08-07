package type_conversion

/*
union u {
	int i;
	long long ll;
};

union u v_u = {
	.ll = 202008070000
};
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

// 联合类型在cgo中会被转化为相应长度的字节数组
func GetSetUnionTypeVar() {
	// print: [8]uint8
	fmt.Printf("%T\n", C.v_u)
	// print: [112 135 158 8 47 0 0 0]
	fmt.Println(C.v_u)

	// ------ get value ------
	// 1.通过unsafe.Pointer强转类型 (性能最好 best performance)
	var i64 = *(*int64)(unsafe.Pointer(&C.v_u))
	// print: 202008070000
	fmt.Println(i64)
	var i32 = *(*int32)(unsafe.Pointer(&C.v_u))
	// print: 144607088 (truncated)
	fmt.Println(i32)

	// 2.手动解码二进制数据
	i64 = 0
	var buf = bytes.NewBuffer([]byte{})
	_ = binary.Write(buf, binary.LittleEndian, C.v_u)
	fmt.Println(buf.Bytes())
	_ = binary.Read(buf, binary.LittleEndian, &i64)
	// print: 202008070000
	fmt.Println(i64)

	// ------ set value ------
	// 1.手动编码二进制数据
	i64 = 123456
	buf = bytes.NewBuffer([]byte{})
	_ = binary.Write(buf, binary.LittleEndian, i64)
	var bs = buf.Bytes()
	for i := 0; i < 8; i++ {
		if i < len(bs) {
			C.v_u[i] = bs[i]
		} else {
			C.v_u[i] = byte(0)
		}
	}
	// [64 226 1 0 0 0 0 0]
	fmt.Println(C.v_u)

	// 2.创建新的联合
	C.v_u = C.union_u{ //使用union_前缀加类型名访问C端联合
		112, 135, 158, 8, 47, // 16进制数据
	}
	// print: [112 135 158 8 47 0 0 0] (202008070000)
	fmt.Println(C.v_u)
}
