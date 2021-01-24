package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "2.13 12.12 2 1.2 5.6 12 acb"

	//1. 解析正则表达式
	reg := regexp.MustCompile(`\d+\.\d+`) //+表示前一个字符1次或多次
	if reg == nil {
		fmt.Println("reg = ", reg)
		return
	}

	//2. 提取关键信息
	s := reg.FindAllStringSubmatch(buf, -1) //-1表示匹配所有
	fmt.Println("s = ", s)
}
