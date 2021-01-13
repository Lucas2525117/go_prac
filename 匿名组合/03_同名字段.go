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
	var s Student

	//默认规则(就近原则: 如果能在本作用域找到此成员，则操作此成员;如果没有找到,则找继承的字段)
	s.name = "lucas"
	s.Person.name = "koko"

	fmt.Println(s.name, s.Person.name)
}
