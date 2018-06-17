package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/17  14:39
*描述:  switch 测试 类型
*/
type DK interface {}

type  Drink struct {
	name string
	price float32
}

func main() {

	var d DK = Drink{"green tea", 7.5}

	switch value := d.(type) {  //测试 类型
	case int:
		fmt.Println("int", value)
	case float32:
		fmt.Println("float32", value)
	case Drink:
		fmt.Println("Drink", value)
	}
}
