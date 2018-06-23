package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/23  11:27
*描述: 关闭 管道,  防止死锁
*/

func main() {

	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 10; i++{
			ch <- i
		}
		//close(ch) //子协程关闭channel, 死锁
	}()

	for{
		if value, ok :=  <- ch; ok{
			fmt.Println(value)
		}else{
			break
		}
	}


}
