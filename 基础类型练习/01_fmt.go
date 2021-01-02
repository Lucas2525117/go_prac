package main

import "fmt"

func main() {
	a := 10
	b := 'a' //97
	c := 2.1
	d := "abc"
	fmt.Printf("%T, %T, %T, %T\n", a, b, c, d)

	fmt.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d)
}
