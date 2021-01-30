/*
重点: 主协程退出了，其他子协程也会跟着退出
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() { //调用匿名函数
		j := 0
		for {
			j++
			fmt.Println("j = ", j)
			time.Sleep(time.Second)
		}
	}()

	i := 0
	for {
		i++
		fmt.Println("i = ", i)
		time.Sleep(time.Second)

		if i == 5 {
			break
		}
	}
}
