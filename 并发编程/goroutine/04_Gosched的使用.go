/*
针对的问题: 主协程由于执行的快,导致主协程执行完了,子协程还没来得及执行
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Hello")
		}
	}()

	for i := 0; i < 2; i++ {
		//让出时间片 先让别的协程执行 别的协程执行完 再来执行此协程
		runtime.Gosched()

		fmt.Println("Hello Go")
	}
}
