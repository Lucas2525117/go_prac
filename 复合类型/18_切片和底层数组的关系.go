package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	s := slice[2:5]
	fmt.Println("s = ", s)
	fmt.Printf("len(s) = %d, cap(s) = %d\n", len(s), cap(s))

	s[1] = 10 //会改变底层数组
	fmt.Println("s = ", s)
	fmt.Println("slice = ", slice)

	s1 := s[2:7]
	fmt.Println("s1 = ", s1)

	s1[2] = 17
	fmt.Println("s1 = ", s1)
	fmt.Println("slice = ", slice)
}
