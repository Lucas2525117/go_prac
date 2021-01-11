package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

//值传递 形参无法改变实参
func test0(s Student) {
	s.addr = "shanghai"
}

//操作指针指向的内存
func test1(s *Student) {
	s.addr = "shanghai"
}

func main() {
	var s Student = Student{1, "jack", 'm', 18, "beijing"}

	test0(s)
	fmt.Println("值传递: s = ", s)

	test1(&s)
	fmt.Println("地址传递: s = ", s)
}
