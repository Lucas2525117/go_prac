package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型，通过type给一个函数起名
//FuncType是一个函数类型
//多态的思想
type FuncType func(int, int) int

func main() {
	result := Add(1, 2)
	fmt.Println("result = ", result)

	var f FuncType
	f = Add
	fmt.Println("f(1, 2) = ", f(1, 2))

	f = minus
	fmt.Println("f(12, 6) = ", f(12, 6))
}
