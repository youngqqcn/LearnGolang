// conditions.go
//条件语句
package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello World!")

	//var a, b int = 9, 3
	//支持一个初始化表达式, 初始化字句和条件表达式直接需要用分号分隔
	if a, b := 9, 99; a < b {
		fmt.Println(a)
		fmt.Println(b)
		fmt.Printf("a < b\n")
	} else { //大括号必须和条件语句或else在同一行
		fmt.Printf("a > b\n")
	}

	//switch可以使用任何类型或表达式作为条件语句
	//C/C++中  switch语句只能是 const int
	var score int = 55

	switch score {
	case 90:
		fmt.Println("优秀")
		//fallthrough
	case 80:
		fmt.Println("良好")
		//fallthrough
	case 55, 60, 70:
		fmt.Println("一般")
		fallthrough //强制执行后面的case语句
	default:
		fmt.Println("差")
	}

	var s2 int = 90
	switch { //这里没有写条件
	case s2 >= 90: //这里写判断语句
		fmt.Println("优秀")
	case s2 >= 80:
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}

	//3
	switch s3 := 90; { //只有初始化语句，没有条件
	case s3 >= 90: //这里写判断语句
		fmt.Println("优秀")
	case s3 >= 80:
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}

}
