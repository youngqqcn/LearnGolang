package main

import (
	"fmt"
	"time"
)

/**
*作者: yqq
*日期: 2018/6/23  15:46
*描述: 超时
*/

func main() {

	ch1 := make(chan int)
	quit := make(chan bool)

	go func() {

		for{
			select {
			case x := <-ch1:
				fmt.Println(x)
			case <-time.After(3 * time.Second): //3s
				fmt.Println("timeout...")
				quit <- true
			}
		}

	}()

	for i := 0; i < 5; i++ {
		ch1 <- i
	}

	<-quit  //等待子协程的 退出信号
}
