/*
Marshal
MarshalIndent:格式化编码
切片转换为结构体
*/
package main

import (
	"encoding/json"
	"fmt"
)

//注意: 成员变量首字母必须大写
type MyJson struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

type MyJson1 struct {
	Company  string   `json:"-"`           //"-" 此字段不会输出到屏幕
	Subjects []string `json:"subjects"`    //二次编码
	IsOk     bool     `json:"isok,string"` //以字符串输出
	Price    float64  `json:"price"`
}

func main() {
	//定义一个结构体变量，同时初始化
	s := MyJson1{"Google", []string{"Go", "C++", "C", "PHP"}, true, 3.1415}

	// {"Company":"Google","Subjects":["Go","C++","C","PHP"],"IsOk":true,"Price":3.1415}
	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", " ") //格式化编码
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf)) //buf为切片 需要转换成字符串
}
