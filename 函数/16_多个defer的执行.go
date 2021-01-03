/*
defer的规则:先进后出
*/
package main

import "fmt"

func test(x int) {
	tmp := 100 / x
	fmt.Println("tmp = ", tmp)
}

func main() {
	defer fmt.Println("test1")

	defer fmt.Println("test2")

	test(0)

	defer fmt.Println("test3")
}
