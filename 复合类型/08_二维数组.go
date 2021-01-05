package main

import "fmt"

func main() {
	var a [3][4]int

	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d, ", i, j, a[i][j])
		}

		fmt.Println("\n")
	}

	//可以直接打印
	fmt.Println("a = ", a)

	//多维数组 未定义的为0
	b := [3][4]int{1: {1, 2, 3, 4}}
	fmt.Println("b = ", b)

	c := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	fmt.Println("c = ", c)
}
