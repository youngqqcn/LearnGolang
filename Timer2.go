package main

import (
	"time"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/23  12:13
*描述:
*/

func Hello(timer *time.Timer)  {
	<- timer.C

	fmt.Println("hello")

}


func main() {

	//for i := 0; i < 5; i++ {
	//	timer1 := time.NewTimer(time.Second * 1)
	//	if i == 3 {
	//		timer1.Stop() //停止
	//		//break  //如果不break  会发生死锁
	//	}
	//	fmt.Println(<-timer1.C)
	//}

	//每隔 1s 打印一个hello
	for i := 0; i < 5; i++ {
		//Hello(time.NewTicker(time.Second * 1))
		timer1 := time.NewTimer(time.Second * 1)
		//fmt.Printf("%T\n", timer1)
		Hello(timer1)
	}
}

