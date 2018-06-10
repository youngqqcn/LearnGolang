package main

import "fmt"

func main() {
	var p1 *int
	p1 = new(int)              //p1为*int 类型, 指向匿名的int变量
	fmt.Println("*p1 = ", *p1) //*p1 =  0

	p2 := new(int) //p2为*int 类型, 指向匿名的int变量
	*p2 = 111
	fmt.Println("*p2 = ", *p2) //*p1 =  111
}

