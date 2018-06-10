package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/10  16:32
*描述: 指针的使用
*/

//func main() {
//
//	var a int = 10              //声明一个变量，同时初始化
//	fmt.Printf("&a = %p\n", &a) //操作符 "&" 取变量地址
//
//	var p *int = nil //声明一个变量p, 类型为 *int, 指针类型
//	p = &a
//	fmt.Printf("p = %p\n", p)
//	fmt.Printf("a = %d, *p = %d\n", a, *p)
//
//	*p = 111 //*p操作指针所指向的内存，即为a
//	fmt.Printf("a = %d, *p = %d\n", a, *p)
//
//}


func swap01(a, b int) {
	a, b = b, a
	fmt.Printf("swap01 a = %d, b = %d\n", a, b)
}

func swap02(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	a := 10
	b := 20

	//swap01(a, b) //值传递
	swap02(&a, &b) //变量地址传递
	fmt.Printf("a = %d, b = %d\n", a, b)
}
