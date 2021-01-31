package main

import (
	"fmt"
	"time"
)

func main() {
	//创建管道
	ch := make(chan string)

	defer fmt.Println("主协程结束")

	go func() {
		defer fmt.Println("子协程调用完毕")

		for i := 0; i < 2; i++ {
			fmt.Println("i = ", i)
			time.Sleep(time.Second)
		}

		ch <- "q"
	}()

	str := <-ch //没有数据前，会阻塞
	fmt.Println("str = ", str)
}
