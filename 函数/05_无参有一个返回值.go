package main

import "fmt"

func MyFunc() int {
	return 1
}

func MyFunc1() (result int) {
	return 1
}

//最常用
func MyFunc2() (result int) {
	result = 1
	return
}

func main() {
	a := MyFunc()
	fmt.Println("a = ", a)

	b := MyFunc1()
	fmt.Println("b = ", b)

	c := MyFunc2()
	fmt.Println("c = ", c)
}
