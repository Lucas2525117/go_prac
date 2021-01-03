package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args

	n := len(list)
	fmt.Println("n = ", n)

	//迭代的方式
	for i, data := range list {
		fmt.Printf("list[%d] is %s\n", i, data)
	}
}
