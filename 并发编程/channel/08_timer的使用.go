package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器，设置时间为2s，2s后，往time通道写数据(当前时间)
	//时间到了 timer只会响应一次
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("当前时间:", time.Now())

	//2s后，往timer.C写数据，有数据后，就可以读取
	t := <-timer.C //channel没有数据前会阻塞
	fmt.Println("t = ", t)
}
