package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/10  15:03
*描述: none
*/


//形如...type格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数
// args 是一个 []int
//func Test(args ...int) {
//	fmt.Printf("%T\n", args)
//	for _, n := range args { //遍历参数列表
//		fmt.Println(n)
//	}
//}
//
//func main() {
//	//函数调用，可传0到多个参数
//	Test()
//	Test(1)
//	Test(1, 2, 3, 4)
//}
//
//func MyFunc01(args ...int) {
//	fmt.Println("MyFunc01")
//	for _, n := range args { //遍历参数列表
//		fmt.Println(n)
//	}
//}
//
//func MyFunc02(args ...int) {
//	fmt.Println("MyFunc02")
//	for _, n := range args { //遍历参数列表
//		fmt.Println(n)
//	}
//}
//
//func Test(args ...int) {
//	MyFunc01(args...) //类似C++的完美转发    //按原样传递, Test()的参数原封不动传递给MyFunc01
//	MyFunc02(args[1:]...) //Test()参数列表中，第1个参数及以后的参数传递给MyFunc02
//}
//
//func main() {
//	Test(1, 2, 3) //函数调用
//}




func Test01() (int, string) { //方式1
	return 250, "sb"
}

func Test02() (a int, str string) { //方式2, 给返回值命名
	a = 250
	str = "sb"
	return
}

func main() {
	v1, v2 := Test01() //函数调用
	_, v3 := Test02()  //函数调用， 第一个返回值丢弃
	v4, _ := Test02()  //函数调用， 第二个返回值丢弃
	fmt.Printf("v1 = %d, v2 = %s, v3 = %s, v4 = %d\n", v1, v2, v3, v4)
}


