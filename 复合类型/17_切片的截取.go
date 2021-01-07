package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:high:max] 取下标从low开始的元素 len = high - low, cap = max - low
	s := slice[:]
	fmt.Println("s = ", s)
	fmt.Printf("len(s) = %d, cap(s) = %d\n", len(s), cap(s))

	s1 := slice[3]
	fmt.Println("s1 = ", s1) //3

	s2 := slice[1:5:8]
	fmt.Println("s2 = ", s2)
	fmt.Printf("len(s2) = %d, cap(s2) = %d\n", len(s2), cap(s2))

	s3 := slice[7:] //cap = 3 等价于slice[7:10:10]
	fmt.Println("s3 = ", s3)
	fmt.Printf("len(s3) = %d, cap(s3) = %d\n", len(s3), cap(s3))

	s4 := slice[:4] //cap = 10 等价于slice[0:4:10]
	fmt.Println("s4 = ", s4)
	fmt.Printf("len(s4) = %d, cap(s4) = %d\n", len(s4), cap(s4))

	s5 := slice[1:4] //cap = 9 等价于slice[1:4:10]
	fmt.Println("s5 = ", s5)
	fmt.Printf("len(s5) = %d, cap(s5) = %d\n", len(s5), cap(s5))
}
