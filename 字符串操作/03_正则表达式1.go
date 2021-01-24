package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "abc azc a2c qwe 888 ip9 bn5"

	//1.解析规则:会解析正则表达式 如果成功返回解析器
	//reg := regexp.MustCompile(`a.c`) //a开头 中间任意字符 c结尾
	//reg := regexp.MustCompile(`a[0-9]c`)
	reg := regexp.MustCompile(`a\dc`)
	if reg == nil {
		fmt.Println("reg = ", reg)
		return
	}

	//2.根绝规则提取关键信息
	s := reg.FindAllStringSubmatch(buf, -1) //-1表示匹配所有
	fmt.Println("s = ", s)
}
