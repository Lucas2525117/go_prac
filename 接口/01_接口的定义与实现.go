package main

import "fmt"

//定义接口类型 一般以"er"结尾
type Humaner interface {
	//方法，只有声明，没有实现，由别的类型(自定义)实现
	sayhi()
}

type Student struct {
	name string
	age  int
}

func (s *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", s.name, s.age)
}

type Teacher struct {
	addr  string
	group string
}

func (t *Teacher) sayhi() {
	fmt.Printf("Teacher[%s, %s] sayhi\n", t.addr, t.group)
}

type MyStr string

func (str *MyStr) sayhi() {
	fmt.Printf("MyStr[%s] sayhi\n", *str)
}

func main() {
	var i Humaner

	//只是实现了此接口方法的类型那么这个类型的变量(接收者类型)就可以给i赋值
	s := &Student{"koko", 12} //指针类型
	i = s
	i.sayhi()

	t := &Teacher{"shanghai", "go"}
	i = t
	i.sayhi()

	var str MyStr = "hi go"
	i = &str
	i.sayhi()
}
