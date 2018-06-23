package main

/**
*作者: yqq
*日期: 2018/6/23  11:35
*描述: 单向管道
*/

//可以将 channel 隐式转换为单向队列，只收或只发，
// 不能将单向 channel 转换为普通

func main() {

	//var ch1 chan int  //双向管道
	//var ch2 chan<- int //只写
	//var ch3 <-chan int //只读

	//ch1 <- 9 //向ch1中写数据
	//fmt.Println(<-ch1) //读读取ch1中数据

	//ch2 <- 9 //向ch2中写数据
	//fmt.Println(<-ch2) //试图读 只写管道中的数据

	//ch3 <- 9 //试图向 只读管道写数据
	//fmt.Println(<-ch3) //会导致死锁, 因为管道一直为空

	//ch2 = ch1

}
