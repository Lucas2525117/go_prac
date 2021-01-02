package main

import "fmt"

func main() {
	var a int
	fmt.Printf("请输入变量a: ")
	//fmt.Scanf("%d", &a)
	fmt.Scan(&a)

	//阻塞用户输入
	fmt.Println("a = ", a)
}
