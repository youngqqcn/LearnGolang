package main

import (
	"fmt"
	"time"
)

/**
*作者: yqq
*日期: 2018/6/23  14:58
*描述:
*/


func main() {

	ch := make(chan int)

	//ch <-  9  // fatal error
	go func() {
		fmt.Println(<-ch)
	}()
	ch <-  9  //ok

	time.Sleep(1 * time.Second)

}
