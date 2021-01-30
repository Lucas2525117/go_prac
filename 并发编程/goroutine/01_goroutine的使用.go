/*
协程的创建: go关键字 创建新协程
*/

package main

import (
	"fmt"
	"time"
)

func NewTask() {
	for {
		fmt.Println("This is a new task")
		time.Sleep(time.Second)
	}
}

func main() {
	go NewTask() //新建一个协程(子协程) 新建一个任务

	for {
		fmt.Println("This is a main goroutine")
		time.Sleep(time.Second)
	}
}
