/*
延时的3种方式
sleep.go
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	//方式1
	time.Sleep(2 * time.Second)

	//方式2 定时2s 阻塞2s 2s后会产生一个事件 往channel写数据
	<-time.After(2 * time.Second)

	//方式3
	//延时2s后打印一句话
	timer := time.NewTimer(2 * time.Second)
	<-timer.C //时间未到会阻塞
	fmt.Println("时间到")

}
