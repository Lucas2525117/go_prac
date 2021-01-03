package main

import "fmt"

func main() {
	a := 10
	str := "lucas"

	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ", str)
	}

	f1()

	type FuncType func()
	//声明
	var f2 FuncType
	f2 = f1
	f2()

	//1. 定义匿名函数
	func() {
		fmt.Printf("a = %d, str = %s\n", a, str)
	}() //()代表调用此匿名函数

	//2. 定义匿名函数
	func(i, j int) {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}(1, 2)

	//3. 定义匿名函数 有参有返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}

		return
	}(10, 20)

	fmt.Printf("x = %d, y = %d\n", x, y)
}
