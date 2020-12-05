// +build ignore

package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
    Name  string
    Order string
}

func main() {
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll好的",    "Order": "Dasyuromorphia"},
		{"Name": "中文名字",    "Order": "Dasyuromorphia"}
	]`)

	var animals []Animal
	
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("%+v\n\n", animals)

	// ▲上边是英文的情况，▼中文

	type StuRead struct {
		Name  interface{} `json:"name"`
		Age   interface{}
		HIgh  interface{}
		sex   interface{} //非大写字母开头，所以sex一定不会生成键值对
		Class interface{} `json:"class"`
		Test  interface{} //字符没有Test，所以值为nil 
	}
	
	// 好像没用，再观察
	// type Class struct {
	// 	Name  string
	// 	Grade int
	// }

    str:=[]byte(`{
		"name": "张三"
	}`)
	// "Age": 18,
	// "high": true,
	// "sex": "男",
	// "CLASS": {"naME": "1班", "GradE":3}
	// ▼此处可以直接 stu := new(StuRead), 此时的stu自身就是指针
	var stu StuRead // 也可 stu:=StuRead{}
    // ▼ Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构（2参必须是指针
    err2 := json.Unmarshal(str, &stu)

    //解析失败会报错，如json字符串格式不对，缺"号，缺}等。
    if err2 != nil{
		fmt.Println(err2)
		return 
    }
    fmt.Printf("%+v \n\n", stu)
}