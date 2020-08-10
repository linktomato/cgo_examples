# 简介
本项目包含一些 cgo 的使用示例。主要由以下几个部分组成：
## 类型转换 (已完成)
类型转换的例子都在 type_conversion 包中，包内有几个文件：
* basic_types 一些基本数据类型的转化例子
* pointer_type 指针类型的转化例子
* string_type 字符串(字符数组)的转化例子
* array_type 数组类型的转化例子
* struct_type 结构体类型的转化例子
* enum_type 枚举类型的转化例子
* union_type 联合类型的转化例子

## 函数调用 (已完成)
函数调用的例子都在 function 包中：
* c_to_go C 端调用 Go 端的例子
* go_to_c Go 端调用 C 端的例子
## 库调用 (未开始)

# 参考
* https://golang.org/src/cmd/cgo/doc.go
* https://golang.org/cmd/cgo/#hdr-C_references_to_Go
* https://chai2010.cn/advanced-go-programming-book/ch2-cgo/readme.html
* https://gcc.gnu.org/onlinedocs/gcc-8.4.0/gccgo/C-Type-Interoperability.html
* https://stackoverflow.com/questions/6125683/call-go-functions-from-c
* https://github.com/golang/go/issues/20639
* https://stackoverflow.com/questions/19910647/pass-struct-and-array-of-structs-to-c-function-from-go
* https://www.marlin.pro/blog/cgo-referencing-c-library-in-go/

# 其他
项目还在进行中，欢迎小伙伴纠错或者提意见。

如果本项目对你有帮助，麻烦给小猿点个小星星，感谢！
