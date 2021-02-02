package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	q := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second): //等3s 超时退出
				fmt.Println("超时")
				q <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i //写数据
		time.Sleep(time.Second)
	}

	<-q //q有数据之前会阻塞
	fmt.Println("程序结束")
}
