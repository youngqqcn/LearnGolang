package main

import (
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/16  14:12
*描述: 结构体
*/

type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}

func TestStruct(p *Student){
	fmt.Println(p.id)
	fmt.Println(p.name)
	fmt.Println(p.sex)
	fmt.Println(p.age)
	fmt.Println(p.addr)
}

func main() {

	var p *Student = &Student{3, "yqq", 'm', 5, "sz"}

	p2 := p
	(*p2).name = "yqq2"

	fmt.Println((*p).id)
	fmt.Println((*p).name)
	fmt.Println((*p).sex)
	fmt.Println((*p).age)
	fmt.Println((*p).addr)

	fmt.Println(p.id) //C语言中应是  p->id

	fmt.Println((*p) == (*p2))

	fmt.Println("=============")
	TestStruct(p)

}
