package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

//声明函数类型
type FuncType func(int, int) int

func Calc(a, b int, f FuncType) (result int) {
	fmt.Println("Calc")
	result = f(a, b)
	return
}

func main() {
	result := Calc(1, 1, Add)
	fmt.Println("result = ", result)

	result1 := Calc(2, 1, minus)
	fmt.Println("result1 = ", result1)
}
