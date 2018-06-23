package main

import (
	"net"
	"log"
	"fmt"
	"strings"
)

/**
*作者: yqq
*日期: 2018/6/23  16:00
*描述: 服务器
*/

//处理客户端连接
func dealConnect( conn net.Conn)  {

	defer conn.Close()

	ipAddr := conn.RemoteAddr().String() //ip地址
	fmt.Println("来自", ipAddr, "连接成功")

	buf := make([]byte, 1024)

	for{
		n , err := conn.Read(buf) //阻塞等待客户端的数据
		if err != nil{
			fmt.Println(err)
		}

		fmt.Println("收到来自", ipAddr, "的数据", string(buf[:n]))
		if "exit" == string(buf[:n]){
			fmt.Println(ipAddr, "退出连接")
			return
		}

		//向客户端发送数据
		conn.Write([]byte(strings.ToUpper(string(buf[:n])))) //转成大写

	}

}




func main() {

	listener, err  := net.Listen("tcp", "127.0.0.1:9999")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("监听成功...")

	defer listener.Close() //延时关闭

	for{
		connect , err := listener.Accept() //阻塞等待客户端连接
		if err != nil{
			log.Println(err)
			continue
		}

		go dealConnect(connect)
	}



}



