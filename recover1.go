package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/17  15:49
*描述: recover  的使用
*/

/*

注意：recover只有在defer调用的函数中有效。

如果调用了内置函数recover，并且定义该defer语句的函数发生了panic异常，
recover会使程序从panic中恢复，并返回panic value。
导致panic异常的函数不会继续运行，但能正常返回。
在未发生panic时调用recover，recover会返回nil。

 */


func TestA()  {
	fmt.Println("TestA()")
}

func TestB() (err error) {
	defer func() { ////在发生异常时，设置恢复
		if x := recover(); x != nil{
			err = fmt.Errorf("从以下内部错误中恢复:%v", x)
		}
	}()

	fmt.Println("TestB():1")
	panic("TestB()发生了Panic异常")

	fmt.Println("TestB():2") //panic 已经中断了 TestB() , 所以此句执行不到
	return

	//err = nil
	//return
}

func TestC()  {
	fmt.Println("TestC()")
}


func main() {

	TestA()

	err := TestB()

	if err != nil{
		fmt.Println(err)
	}

	TestC()

}
