package main

import "fmt"

type Humaner interface { //子集
	sayhi()
}

type Personer interface { //超集
	Humaner        //匿名字段 继承了sayhi()
	sing(s string) //接口
}

type Student struct {
	name string
	age  int
}

func (s *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", s.name, s.age)
}

func (stu *Student) sing(s string) {
	fmt.Printf("Student[%s, %d] sings %s\n", stu.name, stu.age, s)
}

func main() {
	//超集可以转换为子集 反过来不可以(可以理解为c++中的父类和子类)
	var p Personer //超集
	s := &Student{"koko", 12}
	p = s

	var h Humaner //子集

	//p = h //error
	h = p
	h.sayhi()
}
