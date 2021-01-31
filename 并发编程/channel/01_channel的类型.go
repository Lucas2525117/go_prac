/*
channel实现同步
*/
package main

import (
	"fmt"
	"time"
)

//创建一个管道
var ch = make(chan int)

func Print(s string) {
	for _, data := range s {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println("\n")
}

//目标:先打印hello 再打印go
func Person1() {
	Print("hello")

	ch <- 1 //给管道写数据(发送) "1"只是用于举例
}

func Person2() {
	<-ch //从管道取数据(接收)，如果通道没有数据，会阻塞

	Print("go")
}

func main() {
	//新建两个协程
	go Person1()
	go Person2()

	for {

	}
}
