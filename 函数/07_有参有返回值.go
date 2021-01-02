package main

import "fmt"

func MyFunc(a, b int) (min, max int) {
	if a > b {
		min = b
		max = a
	} else {
		min = a
		max = b
	}

	return //有返回值的必须添加return
}

func main() {
	min, max := MyFunc(1, 2)
	fmt.Printf("min = %d, max = %d\n", min, max)

	//通过匿名变量"_"丢弃小的值
	_, a := MyFunc(1, 2)
	fmt.Printf("a = %d\n", a)
}
