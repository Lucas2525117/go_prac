/*
重点: copy函数的使用
*/
package main

import (
	"fmt"
)

func main() {
	srcSlice := []int{1, 2, 3, 4, 9, 10}
	dstSlice := []int{5, 6, 7, 8}

	copy(dstSlice, srcSlice)
	fmt.Println("srcSlice = ", srcSlice)
	fmt.Println("dstSlice = ", dstSlice) //结果: [1,2,3,4]

	srcSlice1 := []int{1, 2}
	dstSlice1 := []int{5, 6, 7, 8}

	copy(dstSlice1, srcSlice1)
	fmt.Println("srcSlice1 = ", srcSlice1)
	fmt.Println("dstSlice1 = ", dstSlice1) //结果: [1,2,7,8]
}
