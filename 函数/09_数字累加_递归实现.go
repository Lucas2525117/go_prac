package main

import "fmt"

//普通实现
func test0() (sum int) {
	for i := 0; i <= 100; i++ {
		sum += i
	}

	return
}

//递归实现
func test1(i int) int {
	if i == 1 {
		return 1
	}

	return i + test1(i-1)
}

func test2(i int) int {
	if i == 100 {
		return 100
	}

	return i + test2(i+1)
}

func main() {
	sum := test0()
	fmt.Println("sum = ", sum)

	sum1 := test1(100)
	fmt.Println("sum1 = ", sum1)

	sum2 := test2(1)
	fmt.Println("sum2 = ", sum2)
}
