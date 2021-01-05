package main

import "fmt"

func main() {
	//数组的定义:同一类型的集合
	var id [50]int

	//数组的长度计算方法: len(id)
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d] = %d\n", i, id[i])
	}
}
