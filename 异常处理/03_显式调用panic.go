package main

import (
	"fmt"
)

func testA() {
	fmt.Println("testA")
}

func testB() {
	//fmt.Println("testB")
	panic("this is a panic")
}

func testC() {
	fmt.Println("testC")
}

func main() {
	testA()
	testB()
	testC()
}
