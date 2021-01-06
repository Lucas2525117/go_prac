package main

import (
	"fmt"
)

//重点: 数组作函数参数，是值传递
//实参数组的每一个元素会给形参数组拷贝一份
//形参的数组是实参数组的复制品，即拷贝而来
func modify(a [5]int) {
	a[0] = 10
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	modify(a)
	fmt.Println("a = ", a) //执行结果: a =  [1 2 3 4 5]
}
