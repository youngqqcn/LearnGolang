package main

import "fmt"

func main() {
	a, b := 10, 20
	defer func(x int) { // a以值传递方式传给x
		fmt.Println("defer:", x, b) // b 闭包引用
	}(a)

	a += 10
	b += 100

	fmt.Printf("a = %d, b = %d\n", a, b)

	/*
		运行结果：
		a = 20, b = 120
		defer: 10 120
	*/
}

