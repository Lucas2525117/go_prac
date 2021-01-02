package main

import "fmt"

func main() {
	s := "泰罗"

	if s == "泰罗" {
		fmt.Println("他是泰罗")
	}

	//if支持一个初始化语句
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	}
}
