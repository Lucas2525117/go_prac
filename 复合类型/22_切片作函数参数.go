/*
重点: 切片是引用传递
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func InitData(s []int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(s); i++ {
		s[i] = rand.Intn(100) //创建100以内的随机数
	}
}

func BubbleSort(s []int) {
	n := len(s)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main() {
	n := 10

	//创建一个切片,len是10
	s := make([]int, n)

	//初始化数组
	InitData(s)
	fmt.Println("排序前: ", s)

	//冒泡排序
	BubbleSort(s)
	fmt.Println("排序后: ", s)
}
