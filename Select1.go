package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/23  14:13
*描述: select 语句
*/

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:  //如果能写入
			fmt.Println("写入")
			x, y = y, x+y
		case <-quit: //如果channel已经满了, 如果是退出信号, 则退出

		 	fmt.Println("quit")
			//if value, ok := <- c; ok{
			//	fmt.Println(">>>", value)
			//}
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		fmt.Println(">>>", <-c)
		quit <- 0 //发送退出信号
	}()

	fibonacci(c, quit)


}

