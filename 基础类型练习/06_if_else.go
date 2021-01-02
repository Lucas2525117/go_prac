package main

import "fmt"

func main() {
	if a := 8; a == 10 {
		fmt.Println("a ==10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else {
		fmt.Println("a < 10")
	}
}
