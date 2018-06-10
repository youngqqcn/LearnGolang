package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/10  16:36
*描述: 数组
*/

func main() {

	var a [100]int

	for i := 0; i < 100; i++{
		a[i] = i * 100
	}

	//for i := range a{
	//	fmt.Println(a[i])
	//}
	//
	//for i, v := range a {
	//	fmt.Println("a[", i, "]=", v)
	//}

	fmt.Println(len(a), cap(a))



}
