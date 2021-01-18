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

	pFunc := s.SetPointer // 方法值 不要带"()"
	pFunc()               //s.SetPointer()

	vFunc := s.SetValue //方法值
	vFunc()             //等价于s.SetValue()
}
