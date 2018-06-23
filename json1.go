package main

import (
	"encoding/json"
	"fmt"
)

/**
*作者: yqq
*日期: 2018/6/18  10:44
*描述: json解析
*/

type IT struct {
	//name string  //小写开头的不会被输出
	Name string
	Subjects []string
	IsOk bool
	Price float64
}

func main() {


	//1.通过结构体生成 json
	it1 := IT{"yqq", []string{"C++", "C", "Python", "Go"}, true, 66.5}

	outByteArray , err := json.Marshal(it1)
	if err == nil{
		for _, ch := range(outByteArray){
			fmt.Printf("%c", ch)
		}
	}else{
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println("===================")

	outByteArray2 , err2 := json.MarshalIndent(it1, "", "    ")
	if err2 == nil{
		//for _, ch := range(outByteArray2){
		//	fmt.Printf("%c", ch)
		//}
		fmt.Println( string(outByteArray2) )
	}else{
		fmt.Println(err)
	}

	//2.通过map生成 json
	t2 := make(map[string]interface{})
	t2["Name"] = "yqq"
	t2["Subjects"] = []string{"C++", "C", "Python", "Go"}
	t2["IsOk"] = true
	t2["Price"] = 66666.55

	out2 , err3  := json.Marshal(t2)
	if err3 != nil{
		fmt.Println(err3)
	}else{
		fmt.Println(string(out2))
	}

	//3.解析json到结构体
	b := []byte(`{
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": true,
    "price": 666.666
	}`)

	var it4 IT
	err4 := json.Unmarshal(b, &it4)
	if err4 != nil{
		fmt.Println(err4)
	}else{
		fmt.Println(it4)
	}

	//4.解析到interface
	var it5 interface{}
	err5 := json.Unmarshal(b, &it5)
	if err5 != nil{
		fmt.Println(err5)
		fmt.Println("err")
	}else{
		fmt.Println(it5)
		fmt.Printf("%T", it5) // map[string]interface{}
	}



}
