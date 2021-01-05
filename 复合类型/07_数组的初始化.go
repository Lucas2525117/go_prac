package main

import "fmt"

func main() {
	//1. 常规初始化
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)

	//2. 部分初始化 没有赋值的元素自动赋值为0
	b := [5]int{1, 2, 3}
	fmt.Println("b = ", b)

	//3. 特定元素初始化
	c := [5]int{2: 10, 4: 12}
	fmt.Println("c = ", c)
}
