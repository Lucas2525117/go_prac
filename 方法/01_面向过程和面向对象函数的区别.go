package main

import "fmt"

//面向过程
func Add01(a, b int) int {
	return a + b
}

//面向对象，方法:给某个类型绑定一个函数
type long int

//a为接收者，接收者就是传递的一个参数
func (a long) Add01(b long) long {
	return a + b
}

func main() {
	result := Add01(1, 2)
	fmt.Println("result = ", result)

	//定义变量
	var a long = 2
	result1 := a.Add01(3)
	fmt.Println("result1 = ", result1)
}
