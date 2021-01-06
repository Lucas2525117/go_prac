package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1. 设置种子,只需一次
	rand.Seed(time.Now().UnixNano())

	var a [10]int
	n := len(a)

	//2. 产生100以内随机数
	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
	}

	fmt.Println("排序前 a = ", a)

	//3. 冒泡排序
	for i := 0; i < n-1; i++ {
		for j := 1; j < n-i; j++ {
			if a[j-1] > a[j] {
				a[j-1], a[j] = a[j], a[j-1]
			}
		}
	}

	fmt.Println("排序后 a = ", a)
}
