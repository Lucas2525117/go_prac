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
	p := Person{"koko", 'm', 13}
	p.SetPointer() //先把p转换为&p 相当于(&p).SetPointer()

	p.SetValue()
}
