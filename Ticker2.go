package main

import (
	"time"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/23  14:04
*描述: Ticker定时器的应用
*/

func HelloTicker(ticker *time.Ticker)  {
	<- ticker.C
	fmt.Println("Hello")
}

func main() {

	ticker := time.NewTicker(time.Second * 1)

	for   {
		HelloTicker(ticker)
	}

}
