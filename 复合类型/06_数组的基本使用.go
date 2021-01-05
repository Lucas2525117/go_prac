package main

import "fmt"

func main() {
	var a [10]int
	var b [5]int

	fmt.Printf("len(a) = %d, len(b) = %d\n", len(a), len(b))

	for i := 0; i < len(a); i++ {
		a[i] = i + 1
	}

	//迭代打印
	for j, data := range a {
		fmt.Printf("a[%d] = %d\n", j, data)
	}
}
