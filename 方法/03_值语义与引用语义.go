package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//值传递:一份拷贝
func (p Person) SetValue(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age

	//%p为打印地址
	fmt.Printf("SetValue &p = %p\n", &p)
}

//引用传递:传递自身地址
func (p *Person) SetPointer(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age

	fmt.Printf("SetPointer p = %p\n", p)
}

func main() {
	var s Person = Person{"koko", 'w', 16}
	fmt.Printf("&s = %p\n", &s)

	//s.SetValue("lucas", 'm', 12)
	//fmt.Println("s = ", s)

	s.SetPointer("lucas", 'm', 25)
	fmt.Println("s = ", s)
}
