package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//创建一个map
	m := make(map[string]interface{}, 4)

	m["company"] = "google"
	m["subjects"] = []string{"Go", "C++", "C", "PHP"} //切片
	m["isok"] = true
	m["price"] = 12.2

	//编码成json
	//buf, err := json.Marshal(m)
	buf, err := json.MarshalIndent(m, "", " ")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
}
