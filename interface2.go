package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/17  13:57
*描述: 接口嵌套
*/

type Human interface {
	Say() //说话
}
type Man interface {
	Human //嵌套
	Sing() //唱歌
}

type Singer struct {
	name string
	song string
}

func (s *Singer)Say()  {
	fmt.Println("My name is ", s.name)
}

func (s *Singer)Sing()  {
	fmt.Println("I sing ", s.song)
}


func main() {

	//可以将子类的指针作为 父类的
	var a Man = &Singer{"yqq", "ABC"}
	a.Say()
	a.Sing()

	var b Human = &Singer{"Tom", "DEF"}
	b.Say()
	//b.Sing() // Human 没有 Sing() 方法

}
