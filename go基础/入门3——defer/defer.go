package main

import (
	"fmt"
)

func printA(a int) {
	fmt.Println("value of a in deferred function", a) //后打印
}

type person struct {
	firstName string
	lastName  string
}

//调用方法
func (p person) fullName() {
	fmt.Printf("%s %s", p.firstName, p.lastName)
}

func main() {
	p := person{
		firstName: "John",
		lastName:  "Smith",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")

	// a := 5
	// defer printA(a) //执行时即会传值
	// a = 10
	// fmt.Println("value of a before deferred function call", a) //先打印

	// name := "Naveen"
	// fmt.Printf("Orignal String: %s\n", string(name))
	// fmt.Printf("Reversed String: ")
	// for _, v := range []rune(name) {
	// 	defer fmt.Printf("%c", v)
	// }
}
