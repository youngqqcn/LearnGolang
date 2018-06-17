package main

import (
	"LearnGolang/test"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/16  15:44
*描述: 可见性
*/

func main() {

	stu1 := test.Stu{1, "yqq"}
	fmt.Println(stu1)
	fmt.Println(stu1.Id)
	//fmt.Println(stu.name) // error

}

