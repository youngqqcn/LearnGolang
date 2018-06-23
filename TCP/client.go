package main

import (
	"net"
	"log"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/23  16:01
*描述: 客户端
*/

func main() {

	connect , err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil{
		fmt.Println("连接服务器失败")
		log.Fatal(err)
		return
	}
	fmt.Println("连接服务器成功")

	defer connect.Close()

	buf := make([]byte, 1024)

	for{
		fmt.Printf("请输入发送内容: ")
		fmt.Scan(&buf)
		fmt.Println("发送的内容是:", string(buf) )

		//发送数据
		sendCount, err := connect.Write(buf)
		if err != nil{
			fmt.Println("发送数据失败")
			//fmt.Println(err)
			return
		}
		fmt.Println("发送了 ", sendCount, "字节数据")

		//等待读取服务器的回复数据
		recvCount, err := connect.Read(buf)
		if err != nil{
			//fmt.Println(err)
			if err.Error() == "EOF" {
				fmt.Println("退出成功")
				return
			}else {
				fmt.Println(err)
			}
		}

		fmt.Println("来自服务器的数据:", string(buf[ :recvCount]))

	}

}
