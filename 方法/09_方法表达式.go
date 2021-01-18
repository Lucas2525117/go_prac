package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetValue() {
	fmt.Printf("SetValue: %p, %v\n", &p, p)
}

func (p *Person) SetPointer() {
	fmt.Printf("SetPointer: %p, %v\n", p, p)
}

func main() {
	s := Person{"koko", 'm', 14}
	fmt.Printf("main: %p, %v\n", &s, s)

	//方法值
	//pFunc := s.SetPointer // 方法值 不要带"()"
	//pFunc()               //s.SetPointer()

	//方法表达式
	f := (*Person).SetPointer
	f(&s) //显式把接收者传递过去 == s.SetPointer

	f1 := (Person).SetValue
	f1(s) //s.SetValue()
}
