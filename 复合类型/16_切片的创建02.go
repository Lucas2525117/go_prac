package main

import (
	"fmt"
)

func main() {
	//自动推导类型
	s := []int{1, 2, 3, 4}
	fmt.Println("s = ", s)

	//借助make函数 格式: make(type, len, cap)
	s1 := make([]int, 5, 10)
	fmt.Printf("len(s1) = %d, cap(s1) = %d\n", len(s1), cap(s1))

	//没有指定容量 则长度=容量
	s2 := make([]int, 5)
	fmt.Printf("len(s2) = %d, cap(s2) = %d\n", len(s2), cap(s2))
}
