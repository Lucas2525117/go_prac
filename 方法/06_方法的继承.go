package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

//Person类型,实现了一个方法
func (p *Person) Print() {
	fmt.Printf("name = %s, sex = %c, age = %d\n", p.name, p.sex, p.age)
}

//Student继承了Person,则成员和方法都继承
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

func main() {
	s := Student{Person{"koko", 'm', 15}, 1, "shanghai"}
	s.Print()

	s.Person.Print()
}
