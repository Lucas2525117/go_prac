package main

import (
	"fmt"
	"time"
)

func main() {
	//创建无缓冲的通道
	ch := make(chan int, 0)

	//len(ch)缓冲区剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("子协程:i = ", i)
			ch <- i //ch中的数据没有被读出去之前会阻塞在这
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch
		fmt.Println("num = ", num)
	}
}
