package main

import "fmt"

func main() {
	m := map[int]string{1: "lucas", 2: "wawa", 3: "jack"}

	for key, value := range m {
		fmt.Printf("key = %d, value = %s\n", key, value)
	}

	//如何判断一个key值是否存在 第一个返回值是key对应的value(不存在返回空) 第二个是是否存在
	value, ok := m[3]
	fmt.Printf("value = %s, ok = %v\n", value, ok)
}
