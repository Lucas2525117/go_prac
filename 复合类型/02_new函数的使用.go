/*
go语言中有自动垃圾回收
*/
package main

import "fmt"

func main() {
	//a ：= 10

	//p是*int类型,指向int
	var p *int
	p = new(int)
	*p = 10
	fmt.Println("*p = ", *p)

	q := new(int) //自动推导类型
	*q = 12
	fmt.Println("*q = ", *q)
}
