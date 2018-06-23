package main

import (
	"os"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/18  13:30
*描述: 文件操作
*/





func main() {

	fout, err := os.OpenFile("./test.txt", os.O_CREATE, 0666)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer fout.Close()

	for i := 0; i < 5; i++{
		outStr := fmt.Sprintf("%s%d\n", "hello ", i)
		fout.WriteString(outStr)
	}

}
