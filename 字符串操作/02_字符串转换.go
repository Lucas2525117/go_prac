package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1. 字符串追加
	//转换位字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//第二个数为要追加的数，第三个数指定十进制方式追加
	slice = strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "go")
	fmt.Println("slice = ", string(slice))

	//bool转换为字符串
	var str string
	str = strconv.FormatBool(false)

	//float转换为字符串
	str = strconv.FormatFloat(3.1415, 'f', -1, 64)

	//整型转字符串
	str = strconv.Itoa(1234)
	fmt.Println("str = ", str)

	//字符串转其他类型
	flag, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}

	//字符串转换为整型
	a, err := strconv.Atoi("123")
	if err == nil {
		fmt.Println("a = ", a)
	} else {
		fmt.Println("err = ", err)
	}
}
