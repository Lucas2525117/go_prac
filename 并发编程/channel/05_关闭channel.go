/*
1. 关闭channel后，无法向channel再发送数据(引发panic错误后导致接收立即返回0)
2. 关闭channel后，可以继续向channel接收数据
3. 对于nil channel,无论收发都会别阻塞
*/
package main

import (
	"fmt"
)

func main() {
	//创建无缓冲的通道
	ch := make(chan int, 0)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //写数据
		}

		//不需要再写数据时，关闭channel
		close(ch) //屏蔽此行会死锁 无法跳出下面的死循环
		//ch <- 1 //关闭channel后无法再发送数据
	}()

	//方法1
	/*

		for {
				//如果ok为true,说明ch没有关闭
				num, ok := <-ch //读数据
				if ok == true {
					fmt.Println("num = ", num)
				} else {
					break
				}
			}

	*/

	//方法2
	for num := range ch {
		fmt.Println("num = ", num)
	}
}
