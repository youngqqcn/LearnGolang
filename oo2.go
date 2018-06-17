package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/17  10:00
*描述: 给结构体绑定方法
*/

type People struct{
	id int
	name string
}

// 绑定方法
func (p1 People )ShowInfo()  {
	fmt.Println(p1.id)
	fmt.Println(p1.name)
}

//引用语义
func (pp *People)ShowInfo2()  {
	fmt.Println("ShowInfo2():", pp.id)
	fmt.Println("ShowInfo2()", pp.name)
}

//引用语义
func (pp *People)SetInfoByPointer(id int, name string)  {
	pp.id = id
	pp.name = name
}

//值语义
func (p People)SetInfoByValue(id int , name string)  {
	p.id = id
	p.name = name
}

func main() {

	//var p People = People{5, "yqq"}
	//p.ShowInfo()
	//p.ShowInfo2()

	var p People  =  People{5, "yqq"}


	//值传递
	p.SetInfoByValue(6, "Tom")
	p.ShowInfo()
	fmt.Println("======================")

	//引用传递
	p.SetInfoByPointer(6, "Tom")
	p.ShowInfo()

}
