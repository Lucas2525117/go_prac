package main

import "fmt"

//结构体定义
type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//1. 顺序初始化,每一个成员必须初始化
	var s Student = Student{1, "jack", 'm', 18, "beijing"}
	fmt.Println("s = ", s)

	//2. 部分成员初始化 没有初始化的成员自动赋值为0 字符串则为空
	s1 := Student{name: "lucas", addr: "shandong"}
	fmt.Println("s1 = ", s1)

	s2 := Student{age: 12, id: 2}
	fmt.Println("s2 = ", s2)
}
