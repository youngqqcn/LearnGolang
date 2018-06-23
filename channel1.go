package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/23  11:03
*描述: 无缓冲的channel
*/

func main() {

	ch := make(chan int, 0) //无缓冲管道

	go func() {

		for i := 0; i < 3; i++{
			ch <- i  //向管道中写数据
			fmt.Println(i, "\t", len(ch), "\t", cap(ch))
		}
	}()


	for i := 0; i < 3; i++{
		number := <- ch  //从管道中读取数据
		fmt.Println("number = ", number)
	}





}
