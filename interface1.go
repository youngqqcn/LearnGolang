package main

import (
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/17  11:43
*描述: 接口
*/

type Humaner interface {
	SayHello()
}

type Player struct {
	id 	int
	name string
	degree int
}

type Teacher struct {
	id int
	name string
	subject string
}

//为 Player 绑定 SayHello 方法
func (s *Player)SayHello()  {
	fmt.Println("Hi , I'm a player! ", s.id,  s.name, s.degree)
}

func (t *Teacher)SayHello()  {
	fmt.Println("Hi, I'm a teacher!", t.id, t.name, t.subject)
}

func SayHi(m Humaner)  {
	m.SayHello()
}


func main() {

	p := Player{5, "yqq", 6}
	t := Teacher{999, "Tom", "Math"}

	p.SayHello()
	t.SayHello()

	//多态实现
	SayHi(&p)
	SayHi(&t)

	//多态
	a := make([]Humaner, 2)

	//a[0], a[1] = p, t //error
	a[0], a[1] = &p, &t  //子类指针转为父类


	for _, item := range a{

		item.SayHello()
	}


}
