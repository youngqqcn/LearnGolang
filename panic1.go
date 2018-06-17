package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/17  15:35
*描述: panic 异常
*/

func TestPanic0()  {
	fmt.Println("Test panic0()")
}

func TestPanic1(){
	panic("Test Panic1()")
}

func TestPanic2()  {
	fmt.Println("Test panic2()")
}


func main() {
	TestPanic0()
	TestPanic1()
	TestPanic2()

}
