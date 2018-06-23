package main

import (
	"os"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/18  13:35
*描述: 读文件
*/

func main() {

	fin , err := os.OpenFile("./test.txt", os.O_RDONLY, 0222)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer fin.Close()

	buf := make( []byte, 1) //一个字节一个字节读取
	for i := 0; ; i++{
		n, _ := fin.Read(buf)
		if n == 0{
			break
		}
		fmt.Printf("%c", buf[0])
		//fmt.Println(i , "-->", string(buf))
	}



}
