package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/16  16:21
*描述: 面向对象编程
*/

type Int int



//func (a int)ADD(b int) int  { //error??
//	return  a + b
//}


func (a Int)Add(b Int) Int  {
	return  a + b
}


func main() {

	var a Int = 1
	var b Int = 2
	fmt.Println( a.Add(b))

}
