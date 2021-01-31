package main

import (
	"fmt"
	"time"
)

func Print(s string) {
	for _, data := range s {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Println("\n")
}

func Person1() {
	Print("hello")
}

func Person2() {
	Print("go")
}

func main() {
	//新建两个协程
	go Person1()
	go Person2()

	for {

	}
}
