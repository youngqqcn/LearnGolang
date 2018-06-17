package test

//只能在本文件中使用
type stu struct {
	id int
	name string
}

//可以在其他包中使用
type Stu struct {
	Id int
	name string  //这个变量只能在本文件中使用
}


