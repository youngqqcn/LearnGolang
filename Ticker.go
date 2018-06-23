package main

import (
	"time"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/23  13:51
*描述: Ticker  定时器
*/

func main() {

	//每隔 1s 想channel发送一个数据
	ticker := time.NewTicker(time.Second * 1)

	go func() {
		for i := 0; i < 10; i++{
			<-ticker.C
			fmt.Println("i = ", i)

			if i == 5{
				ticker.Stop()
			}
		}

	}()

	for{
	}

}
