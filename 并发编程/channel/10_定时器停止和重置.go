package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	b := timer.Reset(1 * time.Second) //定时器重置
	fmt.Println("b = ", b)

	<-timer.C
	fmt.Println("时间到") //3s后打印
}

func main1() {
	timer := time.NewTimer(3 * time.Second)

	go func() { //创建一个匿名函数作为子协程
		<-timer.C //3s过后停止阻塞 执行下面的打印
		fmt.Println("定时器时间到,子协程打印")
	}()

	timer.Stop() //停止定时器

	for { //主协程死循环

	}
}
