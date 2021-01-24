package main

import (
	"fmt"
	"strings"
)

func main() {
	//"hellogo"中是否包含"hello",包含返回true, 不包含返回false
	//Contains
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))

	//Join:组合
	s := []string{"abc", "hello", "koko"}
	buf := strings.Join(s, " ") //以" "为间隔
	fmt.Println("buf =", buf)

	//Index:查找字串的位置
	fmt.Println(strings.Index("abchellogo", "hello")) //返回具体的位置
	fmt.Println(strings.Index("abchellogo", "txt"))   //不存在返回-1

	//Repeat:查看重复次数
	buf1 := strings.Repeat("go", 3)
	fmt.Println("buf1 = ", buf1) //"gogogo"

	//Split:拆分 指定分隔符
	buf1 = "hello#go#index#queen"
	s1 := strings.Split(buf1, "#")
	fmt.Println("s1 = ", s1) //放在切片里

	//Trim去掉两头的字符
	s2 := strings.Trim("  are you ok?  ", " ") //去掉两头空格
	fmt.Println("s2 =", s2)

	//Fields 去掉空格 把元素放入切片
	s3 := strings.Fields("  are you ok?  ")
	//fmt.Println("s3 =", s3)
	for i, data := range s3 {
		fmt.Println(i, data)
	}
}
