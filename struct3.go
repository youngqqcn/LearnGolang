package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/16  15:56
*描述: 匿名组合
*/

type Person struct {

	name string
	sex byte
	age int
}

type Stu struct {
	Person //相当于公有继承 Person
	id 	int
	addr string
	name string //与Person.name同名
}


func main() {

	stu1 := Stu{Person{"yqq", 'm', 5}, 55, "sz", "yqq" }

	fmt.Println(stu1)

	fmt.Println(stu1.Person.name) //显示调用Person中的name, 相当于调用父类的中成员变量
	fmt.Println(stu1.name) //代用本类的成员变量

}
