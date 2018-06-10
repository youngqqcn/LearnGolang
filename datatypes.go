// datatypes.go
//数据类型
package main

import (
	"fmt"
)

func main() {

	var a string = "hello"

	fmt.Printf("%T\n", a)

	b := 4234.343242342 //定义
	fmt.Printf("%T", b)
	fmt.Printf("%.2f\n", b) //只保留两位小数
	fmt.Printf("%p\n", &b)
	fmt.Printf("%f\n", *(&b)) //指针

	c := 4324
	fmt.Printf("%v\n", c) // %v 使用默认格式输出

	d := false
	fmt.Printf("%t\n", d) // bool类型

	type bigint int64
	var bi bigint = 983244234324234
	fmt.Printf("%T\n", bi) //main.bigint
	fmt.Printf("%v\n", bi)

	//fmt.Println("Hello World!")
}
