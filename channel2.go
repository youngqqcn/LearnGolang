package main

import "fmt"


/**
*作者: yqq
*日期: 2018/6/23  11:03
*描述: 死锁
*/

func main2()  {
	/*
	会发生死锁, main routine 等待 ch中的数据, 但是 ch一直为空
	 */

	ch := make(chan int, 3) //无缓冲管道

	go func() {
		for i := 0; i < 3; i++{
		//for{
			ch <- 1  //向管道中写数据
			fmt.Println(1, "\t", len(ch), "\t", cap(ch))
		}
	}()


	for{
		for i := 0; i < 3; i++{
			number := <- ch  //从管道中读取数据
			fmt.Println("number = ", number)
		}
	}
}




func main1() {

	ch := make(chan int, 3) //无缓冲管道

	go func() {

		for{
			for i := 0; i < 3; i++{
				ch <- i  //向管道中写数据
				fmt.Println(i, "\t", len(ch), "\t", cap(ch))
			}
		}

	}()


	for{
		for i := 0; i < 3; i++{
			number := <- ch  //从管道中读取数据
			fmt.Println("number = ", number)
		}
	}
}

func main()  {

	main2()
	main1()

}


