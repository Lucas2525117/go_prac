/*
要点: 切片是一种数据结构
*/
package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 0, 0} //切片初始化
	s := a[0:3:5]             //从下标0开始
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) //长度: 3----(3 - 0)
	fmt.Println("cap(s) = ", cap(s)) //容量: 5----(5 - 0)

	s = a[1:4:5]
	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) //长度: 3----(4 - 1)
	fmt.Println("cap(s) = ", cap(s)) //容量: 4----(5 - 1)
}
