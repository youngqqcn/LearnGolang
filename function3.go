package main

import "fmt"

type FuncType func(int, int) int //声明一个函数类型, func后面没有函数名

//函数中有一个参数类型为函数类型：f FuncType
func Calc(a, b int, f FuncType) (result int) {
	result = f(a, b) //通过调用f()实现任务
	return
}

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}

func main() {
	//函数调用，第三个参数为函数名字，此函数的参数，返回值必须和FuncType类型一致
	result := Calc(1, 1, Add)
	fmt.Println(result) //2

	var f FuncType = Minus
	fmt.Println("result = ", f(10, 2)) //result =  8
}

