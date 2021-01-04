package main

import "fmt"

//不同作用域
func swap(a, b int) {
	a, b = b, a
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func main() {
	a, b := 10, 20
	swap(a, b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}
