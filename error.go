package main

import (
	"errors"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/17  14:55
*描述: 异常处理
*/

func Div(a, b float64) (result float64, err error)  {

	if 0 == b{
		result = 0.0
		err = errors.New("divide by zero")
		return
	}

	result = a / b
	err = nil
	return
}

func main() {

	result , err := Div(5.0, 0.0 )

	if err != nil{
		fmt.Println(result, err)
	}else{
		//fmt.Println(result, err)
	}


}
