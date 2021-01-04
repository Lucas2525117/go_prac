package main

import "fmt"

func main() {
	var a int = 10

	fmt.Printf("a = %d\n", a)
	fmt.Printf("&a = %p\n", &a)

	//保存某个变量的地址 需要指针类型
	var p *int = nil
	p = &a
	fmt.Printf("p = %v, &a = %v\n", p, &a)

	*p = 12
	fmt.Printf("*p = %v, a = %v\n", *p, a)

	var q *int = nil //不能q := nil
	fmt.Println("q = ", q)

}
