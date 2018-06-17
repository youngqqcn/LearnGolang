package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/17  14:13
*描述: 空接口 和 comma-ok 断言
*/

type Element interface {}

type Book struct {
	name string
	author string
}

func main() {

	var a Element = 9
	//var a int = 9
	fmt.Printf("%T", a) //打印int, 而不是 interface
	fmt.Println(a)


	//comma-ok断言
	// element.(Type)      //其中 element是 Interface 变量
	if _ , ok := a.(int) ; ok{
		fmt.Println("int: ", a)
	}else if value, ok := a.(string); ok{
		fmt.Println("string: ", value)
	}

	var b Element = Book{"Gone with the wind", "Michalle"}

	if value , ok := b.(Book); ok{
		fmt.Println(value.name, " creat  by ", value.author)
	}

	// 可以体现Value的用途
	//if _, ok := b.(Book); ok {
	//	fmt.Println(b.name, " creat  by ", b.author)
	//}


}
