/*
要点: 切片和数组的区别
*/
package main

import (
	"fmt"
)

func main() {
	//数组[]里面的长度是一个固定的常量，数组不能修改长度，len和cap为固定值
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("len(a) = %d, cap(a) = %d\n", len(a), cap(a))

	//切片[]里面为空,切片的长度和容量不固定
	s := []int{}
	fmt.Printf("len(s) = %d, cap(s) = %d\n", len(s), cap(s))

	s = append(s, 1) //切片末尾追加成员
	fmt.Printf("len(s) = %d, cap(s) = %d\n", len(s), cap(s))
}
