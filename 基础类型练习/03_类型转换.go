package main

import "fmt"

func main() {
	var flag bool
	flag = true
	fmt.Printf("flag = %v\n", flag)

	//flag = 1

	//兼容类型才可以转
	var ch byte
	ch = 'a'
	var t int
	t = int(ch)
	fmt.Println("t = ", t)
}
