package main

import (
	"fmt"
	"time"
)

func main() {
	go func() { //调用匿名函数
		j := 0
		for {
			j++
			fmt.Println("j = ", j)
			time.Sleep(time.Second)
		}
	}()
}
