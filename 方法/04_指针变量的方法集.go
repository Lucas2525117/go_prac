package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetValue() {
	fmt.Println("SetValue")
}

func (p *Person) SetPointer() {
	fmt.Println("SetPointer")
}

func main() {
	//结构体变量是一个指针变量，它能够调用哪些方法，这些方法就是一个集合，简称方法集
	p := &Person{"koko", 'm', 13}
	p.SetPointer()
	(*p).SetPointer() //把(*p)转换成p

	//内部做的转换:先把指针p转换成*p再调用
	(*p).SetValue()
	p.SetValue()
}
