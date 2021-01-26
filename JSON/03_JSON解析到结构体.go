/*
Unmarshal: 第二个参数地址传递
*/
package main

import (
	"encoding/json"
	"fmt"
)

type MyJson struct {
	Company  string   `json:"company"`  //"-" 此字段不会输出到屏幕
	Subjects []string `json:"subjects"` //二次编码
	IsOk     bool     `json:"isok"`     //以字符串输出
	Price    float64  `json:"price"`
}

func main() {
	jsonBuf := `
	{
	"company":"Google",
	"subjects":[
		"Go",
		"C++",
		"C",
		"PHP"
	],
	"isok":true,
	"price":3.1415
	}
	`
	var myJson MyJson
	err := json.Unmarshal([]byte(jsonBuf), &myJson) //注意:第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Printf("myJson = %+v\n", myJson) //"%+v"
}
