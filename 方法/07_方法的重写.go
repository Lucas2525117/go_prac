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

//Student也实现了一个方法，这个方法与Person同名，这种方法叫重写
func (p *Student) Print() {
	fmt.Printf("name = %s, sex = %c, age = %d, id = %d, addr = %s\n", p.name, p.sex, p.age, p.id, p.addr)
}

func main() {
	s := Student{Person{"koko", 'm', 15}, 1, "shanghai"}
	//就近原则:先找本作用域的方法，找不到再用继承的方法
	s.Print()

	//显式调用继承的方法
	s.Person.Print()
}
