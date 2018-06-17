package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/17  11:08
*描述: 方法表达式
*/


type Worker struct {
	Id int
	Name string
}

func (w *Worker )ShowInfo()  {
	fmt.Println(w.Id)
	fmt.Println(w.Name)
}

func (w Worker )ShowInfoByValue()  {
	fmt.Println(w.Id)
	fmt.Println(w.Name)
}

func (w *Worker)SetWorkerInfo(id int , name string)  {
	w.Id = id
	w.Name = name
}


func (w Worker)SetWorkerInfoByValue(id int , name string)  {
	w.Id = id
	w.Name = name
}

func main() {

	pFunc := (&Worker{5, "yqq"}).ShowInfo
	pFunc()

	fmt.Println("===============")
	w := Worker{6, "tom"}
	w.ShowInfo()

	fmt.Println("===============")
	pShow := (*Worker).ShowInfo
	pShow(&w)

	fmt.Println("===============")
	//(*Worker).SetWorkerInfo(&w, 5, "jerry") //ok
	pSetFunc := (*Worker).SetWorkerInfo//(&w, 5, "jerry")   //(&w, 5, "jerry")
	pSetFunc(&w, 5, "jerry")
	w.ShowInfo()

}
