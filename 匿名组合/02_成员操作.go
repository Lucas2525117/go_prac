package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型 没有名字:匿名字段 继承了Person里面的成员
	id     int
	addr   string
	name   string
}

func main() {
	var s Student = Student{Person{"jack", 'n', 18}, 1, "shanghai", "koko"}
	s.name = "lucas"

	s.Person.age = 12

	fmt.Println(s.addr, s.age, s.id, s.name, s.sex)
}
