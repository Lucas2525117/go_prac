package main

import (
	"encoding/json"
	"fmt"
)

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

	//创建map
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonBuf), &m) //注意:第二个参数要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//fmt.Println("m = ", m)
	//fmt.Printf("m = %+v\n", m)

	// var str string
	// str = string(m["company"])
	// fmt.Println("str = ", str)

	var str string

	//类型断言
	for key, value := range m {
		//fmt.Printf("%v, %v\n", key, value)

		// if key == "company" {
		// 	str = string(value)
		// 	fmt.Println("str = ", str)
		// }

		//类型反推, 通过value的类型反推
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("map[%s]的值类型为string, value = %s\n", key, str)
		case bool:
			fmt.Printf("map[%s]bool, value = %v\n", key, data)
		case float64:
			fmt.Printf("map[%s]float64, value = %v\n", key, data)
		case []interface{}: //利用万能指针 不能直接判断切片
			fmt.Printf("map[%s][]interface{}, value = %v\n", key, data)
		}
	}
}
