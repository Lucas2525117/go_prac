package main

import "fmt"

func main() {
	str := "abc"

	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	//range 默认返回两个元素 一个是元素的位置 一个是元素本身
	for i, data := range str {
		fmt.Printf("str[%d] = %c\n", i, data)
	}

	for i := range str { //第2个返回值 默认丢弃
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}

	for i, _ := range str {
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}
}
