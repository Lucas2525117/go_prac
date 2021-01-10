package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateNum(p *int) {
	//设置种子
	rand.Seed(time.Now().UnixNano())

	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}
	}

	*p = num
}

func GetNum(s []int, num int) {
	s[0] = num / 1000       //千位
	s[1] = num % 1000 / 100 //百位
	s[2] = num % 100 / 10   //十位
	s[3] = num % 10         //个位
}

func PlayOneGame(s []int) {
	var num int
	keySlice := make([]int, 4)

	for {
		for {
			fmt.Printf("请输入一个四位数: ")
			fmt.Scan(&num)

			if num > 999 && num < 10000 {
				break
			}

			fmt.Printf("输入的数不符合要求,请重新输入\n")
		}

		GetNum(keySlice, num)

		index := 0
		for i := 0; i < 4; i++ {
			if keySlice[i] > s[i] {
				fmt.Printf("第%d位大了一点\n", i+1)
			} else if keySlice[i] < s[i] {
				fmt.Printf("第%d位小了一点\n", i+1)
			} else {
				fmt.Printf("第%d位猜对了\n", i+1)
				index++
			}
		}

		if index == 4 {
			fmt.Println("全部猜对")
			break
		}
	}
}

func main() {
	var randNum int

	//产生一个4位的随机数
	CreateNum(&randNum)
	randSlice := make([]int, 4)

	//获取切片
	GetNum(randSlice, randNum)
	//猜切片的值
	PlayOneGame(randSlice)
}
