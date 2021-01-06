package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//1. 设置种子,只需一次
	//如果种子参数一样，每次运行程序产生的随机数都一样
	rand.Seed(time.Now().UnixNano())

	//2. 产生随机数
	for i := 0; i < 5; i++ {
		//fmt.Println("rand = ", rand.Int()) //产生随机数 随机数很大
		fmt.Println("rand = ", rand.Intn(100)) //限制在100以内的数
	}
}
