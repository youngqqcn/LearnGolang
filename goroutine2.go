package main

import (
	"runtime"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/23  10:43
*描述: cpu 的个数
*/


func main() {

	cpuN := runtime.NumCPU()
	fmt.Println(cpuN)

	runtime.GOMAXPROCS(2) //设置用来计算的 cpu 核数
	//runtime.GOMAXPROCS(1) //设置用来计算的 cpu 核数


	for i := 0; i < 100000; i++{
		go fmt.Print(1)
		fmt.Print(0)
	}

}
