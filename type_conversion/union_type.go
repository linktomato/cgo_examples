package type_conversion

/*
union u {
	int i;
	long long ll;
};

union u v_u = {20200807};
*/
import "C"
import "fmt"

func GetSetUnionTypeVar() {
	// print: [8]uint8
	fmt.Printf("%T\n", C.v_u)
	fmt.Println(C.v_u)
}
