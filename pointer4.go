package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/16  11:02
*描述: 数组指针  指针数组
*/
const MAX int = 9


//数组指针
func arraypointer(){
	a  :=  [MAX]int{1, 2, 3, 4, 5}

	//1.自动推到出  数组指针类型
	//pa := &a  //数组指针

	//2.手工定义
	var pa *[MAX]int
	pa = &a

	for i := 0; i < len(*pa); i++{
		(*pa)[i] = 9 * i
	}

	fmt.Println(*pa)
}

//指针数组
type array  []int
func pointerarray()  {

	//var a [MAX]int
	//for i, _ := range a{
	//	a[i] = i * 9
	//}
	//fmt.Println(a)

	var a [MAX]*int    //指针数组

	for i , _ := range a{
		a[i] = new(int)
		*(a[i]) = 9 * i
		//fmt.Printf("%T\n", a[i])
	}

	for _, value := range a{
		fmt.Printf("%d\t", *value)
	}

}

//数组指针数组   实现二维数组
/*

     a[0]--->[x, x, x, x]
     a[1]--->[x, x, x, x]
     a[2]--->[x, x, x, x]
     ......

 */
func arraypointerarray()  {

	var a [MAX]*[MAX]int     //   [MAX]  *[MAX]int
	                          //     ^         ^
	                          //    |         |
	                          //   数组      数组指针
	                          // 从后往前组合成类型名:
	                          //    数组指针数组

	for i , _ := range a{
		a[i] = new([MAX]int)
		for j , _ := range a[i]{
			a[i][j] = i * j
		}
		//fmt.Printf("%T\n", a[i])
	}

	for i, _  := range a{
		for j, _ := range a[i]{
			fmt.Printf("%d\t", a[i][j])
		}
		fmt.Println()
	}

}

//数组指针数组指针
/*

 pa-->a[0]--->[x, x, x, x]
     a[1]--->[x, x, x, x]
     a[2]--->[x, x, x, x]
     ......

 */
func arraypointerarraypointer()  {

	var a[MAX] *[MAX]int  //数组指针数组:  元素类型是数组指针  的数组
	var pa *[MAX] *[MAX]int //数组指针数组指针:  指向 数组指针数组  的指针
	pa = &a

	for i, _ := range *pa{
		(*pa)[i] = new([MAX]int)  //为数组指针 分配空间
		for j, _ := range (*pa)[i]{
			(*pa)[i][j] = i * j
		}
	}

	for _, a1 := range *pa{
		for _ , a2 := range a1{
			fmt.Printf("%d\t", a2)
		}
		fmt.Println()
	}

}

func main() {

	//arraypointer()
	//pointerarray()

	//arraypointerarray()

	arraypointerarraypointer()

}
