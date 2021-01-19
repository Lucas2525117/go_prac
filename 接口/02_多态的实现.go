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

//定义一个普通函数，函数的参数为接口类型
//多态:一个函数，有多种不同的表现
func WhoSayHi(i Humaner) {
	i.sayhi()
}

func main() {
	s := &Student{"koko", 12}
	t := &Teacher{"shanghai", "go"}
	var str MyStr = "hi go"

	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	m := make([]Humaner, 3)
	m[0] = s
	m[1] = t
	m[2] = &str

	//第一个返回下标，第二个返回下标对应的值
	for _, i := range m {
		i.sayhi()
	}
}
