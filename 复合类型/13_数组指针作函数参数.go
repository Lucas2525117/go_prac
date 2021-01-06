package main

import (
	"fmt"
)

//p指向实现数组a,他是指向数组,是数组指针
//*p代表指针所指向的内存,就是实参a
func modify(p *[5]int) {
	(*p)[0] = 10
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	modify(&a)             //地址传递
	fmt.Println("a = ", a) //执行结果:a =  [10 2 3 4 5]
}
