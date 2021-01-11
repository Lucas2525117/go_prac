package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	var s Student

	//注意使用"."符号
	s.name = "jack"
	s.id = 1
	s.addr = "shanghai"
	fmt.Println("s = ", s)
}
