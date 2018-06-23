package main

import (
	"time"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/23  11:59
*描述: 定时器
*/



func main() {

	timer1 := time.NewTimer(time.Second * 2) //2s之后向 C管道 发送一个数据
	t1 := time.Now()
	fmt.Println(t1)

	curTime2 := <- timer1.C //阻塞等待  timer1.C 管道的数据
	fmt.Println(curTime2)

	timer1.Reset(time.Second * 5)

	curTime3 := <- timer1.C
	fmt.Println(curTime3)

	//Timer 只是触发一次, 并不是每隔多少秒 就触发一次, 这一点要注意
	//for i := 0; i < 5; i++  {
	//	curTime3 := <- timer1.C
	//	fmt.Println(curTime3)
	//}



}
