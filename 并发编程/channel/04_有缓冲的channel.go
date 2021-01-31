/*
格式: make(chan type, capacity)
如果给定缓冲区容量，通道是异步的。只要缓冲区有未使用的空间用于发送数据，或者还包括
可以接收的数据，那么其通信就会无阻塞的进行
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	//创建有缓冲的通道
	ch := make(chan int, 3)

	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch)) //len(ch):0 cap(ch):3

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i //ch满之后不能再写
			fmt.Printf("子协程[%d]:len(ch) = %d, cap(ch) = %d\n", i, len(ch), cap(ch))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch
		fmt.Println("num = ", num)
	}
}
