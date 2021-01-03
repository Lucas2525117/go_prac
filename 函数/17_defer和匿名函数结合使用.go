package main

import "fmt"

func main() {
	a := 10
	b := 20

	defer func(a, b int) {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}(a, b) //执行到这个函数的时候，a,b已经传进去了10、20,只是要等到main函数执行完才能进行打印

	a = 30
	b = 40
	fmt.Printf("a = %d, b = %d\n", a, b)
}
