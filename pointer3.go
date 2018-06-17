package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/16  10:54
*描述: 函数指针
*/

func Foo(a, b int) int {
	return a + b
}

type FType func(int, int) int
//type PFType *func(int, int) int     //这样不行

func main() {

	var f FType
	f = Foo
	fmt.Println(f(5, 7))

	//var pf PFType
	//pf = &Foo
	//fmt.Println((*pf)(5, 7))
}
