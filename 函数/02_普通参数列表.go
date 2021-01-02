package main

import "fmt"

func MyFunc(a int) {
	fmt.Println("a = ", a)
}

func MyFunc1(a, b int) {
	fmt.Printf("a = %d, b = %d\n", a, b)
}

func MyFunc2(a int, b byte, c string) {
	fmt.Printf("a = %d, b = %c, c = %s\n", a, b, c)
}

func main() {
	MyFunc(11)

	MyFunc1(12, 13)

	MyFunc2(14, 'c', "abc")
}
