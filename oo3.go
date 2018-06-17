package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/17  10:19
*描述: 方法继承
*/


type BasePeople struct {
	Id int
	Name string
}

func (bp *BasePeople )ShowInfo()  {
	fmt.Println(bp.Id)
	fmt.Println(bp.Name)
}

func (bp *BasePeople)SetInfo(id int, name string)  {
	bp.Id = id
	bp.Name = name
}


type XPeople struct {
	BasePeople//继承 people的属性, 包括方法
	XValue int
}

//重定义 父类的 ShowInfo 方法
func (xp *XPeople)ShowInfo()  {
	xp.BasePeople.ShowInfo() //调用父类的
	fmt.Println(xp.XValue)
}


func main() {

	xp := XPeople{BasePeople{5, "yqq"}, 99}
	xp.BasePeople.ShowInfo() //调用父类(BasePeople) 的ShowInfo方法

	fmt.Println("================")

	xp.ShowInfo() //调用重定义的方法


}
