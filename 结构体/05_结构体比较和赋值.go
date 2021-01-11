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
	//结构体比较(同数组比较)
	var s Student = Student{1, "jack", 'm', 18, "beijing"}
	var s1 Student = Student{1, "jack", 'm', 19, "beijing"}
	fmt.Println("s = s1", s == s1)

	//同类型的两个结构体变量可以相互赋值(同数组赋值)
	var t Student
	t = s1
	fmt.Println("t = ", t)
}
