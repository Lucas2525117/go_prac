/*
GOMAXPROCS: 用来设置可以并行计算的CPU核数的最大值
会影响运行的速度
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.GOMAXPROCS(2) //指定单核运算
	fmt.Println("num = ", num)
	
	for {
		gp fmt.Print(1)
		
		fmt.Print(0)
	}
}
