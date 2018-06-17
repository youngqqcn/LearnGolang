package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/16  12:22
*描述: map
*/

//空指针错误
func TestMap1()  {

	//在一个map里所有的键都是唯一的，
	// 而且必须是支持==和!=操作符的类型，
	// 切片、函数以及包含切片的结构类型这些类型由于具有引用语义，
	// 不能作为映射的键，
	var  m1  map[string]int
	m1["yqq"] = 5   //空指针
	m1["mesii"] = 10
	m1["cr7"] = 7

	fmt.Println(m1)

}

func TestMap2()  {
	//在一个map里所有的键都是唯一的，
	// 而且必须是支持==和!=操作符的类型，
	// 切片、函数以及包含切片的结构类型这些类型由于具有引用语义，
	// 不能作为映射的键，

	//var m1 = make( map[string]int, 10) //这种方式也行?
	m1 := make( map[string]int, 10)

	m1["yqq"] = 5
	m1["mesii"] = 10
	m1["cr7"] = 7

	for k, v := range m1{
		fmt.Println(k, ":", v)
	}
}

func TestDefine()  {

	{
		var x = 9  //其实也是自动推到类型   x := 9 类似
		fmt.Println(x)
	}

	{
		x := 9   //ok的    相当于   var x int
		fmt.Println(x)

		//var x string = "hello"  //error
		//fmt.Println(x)
	}

}

func TestMap3()  {

	//var m map[string]int{"yqq":5, "messi":10} //error
	var m map[string]int = map[string]int{"yqq":5, "messi":10}

	for k , v := range m{
		fmt.Println(k, "=", v)
	}

	//如果ok 存在, 则返回值 和 true, 否则返回 零值(字符串是空), 和false
	value, ok:= m["ok"]
	fmt.Println(value, ok)

	delete(m, "yqq")
	delete(m, "ok")//如果不存在, 若无其事

	for k , v := range m{
		fmt.Println(k, "=", v)
	}
}


func main() {

	//TestMap2()

	//TestDefine()

	TestMap3()

}
