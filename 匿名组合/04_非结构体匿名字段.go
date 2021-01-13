package main

import "fmt"

type mystr string //自定义类型

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //结构体匿名字段
	int    //基础类型的匿名字段
	id     int
	mystr
}

func main() {
	var s Student
	s.int = 1
	s.mystr = "gzd"

	s.Person = Person{name: "fofo"}
	fmt.Println("s = ", s)
	fmt.Printf("s = %+v\n", s)
}
