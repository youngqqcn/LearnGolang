package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/23  14:36
*描述: switch case下有多行语句
*/

func SwitchTest(a int)  {

	switch a {
	case 0:
		fmt.Println("0")
		fmt.Println("hello")
	case 1:
		fmt.Println("1")
		fmt.Println("world")
		b := 9
		if a  < b{
			fmt.Println("a < b")
		}
	default:
		fmt.Println("default")
	}

}

func main() {

	SwitchTest(1)
}
