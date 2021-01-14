/*
重点:
1. 带有接收者的函数叫方法
2. func (receiver ReceiverType) funcname (parameters) (results)
   1) 参数receiver可以任意命名。
   2) 参数receiver类型可以是T或T*。基类型T不能是接口或指针。
   3) 只要接收者类型不同 函数名相同也是不同的方法
*/
package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) PrintInfo() {
	fmt.Println("p = ", p)
}

//通过一个函数给成员赋值
func (p *Person) SetInfo(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age
}

// pointer为接收者类型，本身不能为指针
// type pointer *int

// func (p Pointer) Add01() {

// }

//1.只要接收者类型不同 函数名相同也是不同的方法
type long int

func (a long) test1() {

}

//2.
type char byte

func (b char) test1() {

}

func main() {
	//定义同时初始化
	p := Person{"koko", 'n', 18}
	p.PrintInfo()

	//定义结构体变量
	var p1 Person
	(&p1).SetInfo("lucas", 'n', 14)
	p1.PrintInfo()
}
