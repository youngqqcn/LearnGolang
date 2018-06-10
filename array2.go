package main

import "fmt"

//值传递
//func modify(array [5]int) {
//	array[0] = 10 // 试图修改数组的第一个元素
//	//In modify(), array values: [10 2 3 4 5]
//	fmt.Println("In modify(), array values:", array)
//}
//
//func main() {
//	array := [5]int{1, 2, 3, 4, 5} // 定义并初始化一个数组
//	modify(array)                  // 传递给一个函数，并试图在函数体内修改这个数组内容
//	//In main(), array values: [1 2 3 4 5]
//	fmt.Println("In main(), array values:", array)
//}


//数组指针做参数,  地址传递
func modify(array *[5]int) {
	(*array)[0] = 10
	//In modify(), array values: [10 2 3 4 5]
	fmt.Println("In modify(), array values:", *array)
}

func main() {
	array := [5]int{1, 2, 3, 4, 5} // 定义并初始化一个数组
	modify(&array)                 // 数组指针
	//In main(), array values: [10 2 3 4 5]
	fmt.Println("In main(), array values:", array)
}

