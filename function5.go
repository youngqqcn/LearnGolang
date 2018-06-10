package main

//匿名函数 会修改外部变量的值

import "fmt"

func main() {
	i := 10
	str := "mike"
	func() {
		i = 100
		str = "go"
		//内部：i = 100, str = go
		fmt.Printf("内部：i = %d, str = %s\n", i, str)
	}() //别忘了后面的(), ()的作用是，此处直接调用此匿名函数

	//外部：i = 100, str = go
	fmt.Printf("外部：i = %d, str = %s\n", i, str)
}

