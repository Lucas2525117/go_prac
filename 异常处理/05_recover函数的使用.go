package main

import (
	"fmt"
)

func testA() {
	fmt.Println("testA")
}

func testB(i int) {
	defer func() {
		//recover()
		//fmt.Println(recover())

		if err := recover(); err != nil { //恢复
			fmt.Println(err)
		}
	}() //调用匿名函数

	var a [10]int
	a[i] = 10 //当i大于10时，数组越界，触发panic
}

func testC() {
	fmt.Println("testC")
}

func main() {
	testA()
	testB(11)
	testC()
}
