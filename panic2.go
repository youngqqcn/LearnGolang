package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/17  15:38
*描述: 内置 panic 异常
*/

func InsidePanic()  {
	//a := [2]int //错误的定义方式
	//var a [2]int{1, 2} //错误的定义方式

	//var a = [2]int{1, 2} //ok
	//var a [2]int = [2]int{1, 2} //ok
	fmt.Println("hello")

	a := [2]int{1, 2} //ok
	fmt.Println(a)

	i := 100
	a[i] = 99  //引发内置的 panic

}

func main() {
	InsidePanic()

}
