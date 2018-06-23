package main

import (
	"time"
	"runtime"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/23  10:25
*描述: 并发编程
*/


func testTask(taskId int){
	for i := 0; i < 5; i++ {
		if i == 3{
			runtime.Goexit() //终止当前 goroutine
		}
		fmt.Println(taskId, " -> ", i)
		runtime.Gosched() //主动让出时间片
	}
}


func main() {

	for i := 0; i < 5; i++ {
		go testTask(i)
	}

	time.Sleep(2 * time.Second)

}


