package main

import (
	"fmt"
)

func main() {
	s := []int{}
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	fmt.Println("s = ", s)

	//原切片末尾添加
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Printf("len = %d, cap = %d\n", len(s), cap(s))
	fmt.Println("s = ", s)

	s1 := []int{1, 2, 3}
	s1 = append(s1, 4)
	s1 = append(s1, 5)
	s1 = append(s1, 6)
	fmt.Println("s1 = ", s1)
}
