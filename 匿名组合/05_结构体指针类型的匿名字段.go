package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	*Person //指针类型
	id      int
}

func main() {
	//指针取地址
	s := Student{&Person{"koko", 'n', 12}, 1} //注意取地址符'&'
	fmt.Println(s.name, s.age, s.id, s.sex)

	//分配空间
	var s1 Student
	s1.Person = new(Person)
	s1.name = "lucas"

	fmt.Printf("s1 = %+v\n", s1)
	fmt.Println(s1.name)
}
