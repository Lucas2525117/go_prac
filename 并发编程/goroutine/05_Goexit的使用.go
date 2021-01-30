/*
GoExit: 终止当前协程
*/
package main

import (
	"fmt"
	"runtime"
)

func test1() {
	defer fmt.Println("test1 end")

	//return
	runtime.Goexit() //终止所在的协程

	fmt.Println("test1 begin")
}

func main() {
	//创建新的协程
	go func() { //匿名函数
		fmt.Println("test begin")

		test1()

		fmt.Println("test end")
	}()

	for {

	}
}
