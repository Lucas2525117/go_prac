package main

import (
	"fmt"
)

var a byte

func test() {
	fmt.Printf("%T\n", a) //无局部变量 只能用全部变量
}

func main() {
	var a int = 10
	test() //uint8

	//1. 不同作用域 允许使用同名变量
	//2. 使用变量的原则 就近原则
	fmt.Printf("%T\n", a) //int

	{
		var a float32
		fmt.Printf("%T\n", a) //float32
	}

	test() //uint8
}
