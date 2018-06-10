package main

import "fmt"

func main() {
	i := 0
	str := "yqq"

	//方式1
	f1 := func() { //匿名函数，无参无返回值
		//引用到函数外的变量
		fmt.Printf("方式1：i = %d, str = %s\n", i, str)
	}

	f1() //函数调用

	//方式1的另一种方式
	type FuncType func() //声明函数类型, 无参无返回值
	var f2 FuncType = f1
	f2() //函数调用

	//方式2
	var f3 FuncType = func() {
		fmt.Printf("方式2：i = %d, str = %s\n", i, str)
	}
	f3() //函数调用

	//方式3
	func() { //匿名函数，无参无返回值
		fmt.Printf("方式3：i = %d, str = %s\n", i, str)
	}() //别忘了后面的(), ()的作用是，此处直接调用此匿名函数

	//方式4, 匿名函数，有参有返回值
	v := func(a, b int) (result int) {
		result = a + b
		return
	}(1, 1) //别忘了后面的(1, 1), (1, 1)的作用是，此处直接调用此匿名函数， 并传参
	fmt.Println("v = ", v)

}

