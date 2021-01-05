package main

import "fmt"

func main() {
	//支持比较
	//1. 只支持"=="或"!="
	//2. 比较的是:是不是每个元素都一样
	//3. 2个数组比较,数组的类型要一样
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}

	fmt.Println("a == b ", a == b)
	fmt.Println("a == c ", a == c)

	//1. 同类型的数组可以赋值
	var d [5]int
	d = a
	fmt.Println("d = ", d)
}
