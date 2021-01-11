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
	//方式1
	var s Student
	var s1 *Student = &s //取地址

	s1.addr = "shanghai"
	(*s1).id = 2
	fmt.Println("s1 = ", *s1)
	fmt.Println("s = ", s)

	//方式2
	s2 := new(Student) //通过new
	s2.addr = "beijing"
	(*s2).id = 2
	fmt.Println("s2 = ", *s2)
}
