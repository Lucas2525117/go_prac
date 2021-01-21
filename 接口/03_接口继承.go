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
	//定义一个接口类型的变量
	var p Personer
	s := &Student{"koko", 18}
	p = s
	p.sayhi() //继承的方法
	p.sing("Lonely")
}
