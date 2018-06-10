package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/10  14:38
*描述: go 的 循环语句
*/



func main() {

	//for
	for i:=10; i < 100; i++{
		fmt.Println(i)
	}

	s := "hello world"


	// 迭代
	for index := range s{
		//fmt.Println(index) //  index
		fmt.Printf("%c", s[index])
	}

	for _, ch := range s{   //丢弃 索引
		fmt.Printf("%c", ch)
	}

	for index, ch := range s{
		fmt.Printf("%d, %c", index, ch)
		//fmt.Println(s[index], ch)
	}


	for i := 0; i < 100; i++{
		if i % 2 == 0{
			continue
		}
		if i == 55{
			break
		}
		fmt.Println(i)  //打印所有奇数
	}


	






}
