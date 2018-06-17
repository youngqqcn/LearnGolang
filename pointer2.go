package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/16  10:39
*描述: 数组指针
*/

func foo(pa  *[5]int/*len int*/){
	fmt.Println(len(*pa))   //  使用 len()  获取数组的长度
	fmt.Println(cap(*pa))   //cap() 获取容量

	for i := 0; i < len(*pa); i++{
		fmt.Printf("%d ", (*pa)[i])
	}
	fmt.Println()

	fmt.Println((*pa)[2:])  //切片
}



func main() {

	//array := [5]int{1, 2, 3, 4, 5}
	//array := [...]int{1, 2, 3, 4, 5}
	array := [5]int{3:5, 4:6}
	foo(&array)

}
