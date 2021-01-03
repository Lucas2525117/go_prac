/*
闭包的特点:
不关心捕获的变量或常量时候超出作用域，只要闭包还在使用它，这些变量就会存在
*/
package main

import "fmt"

//test0的返回值是一个匿名函数，返回一个函数类型
func test0() func() int {
	var x int

	return func() int {
		x++
		return x * x
	}
}

func main() {
	//返回值是一个匿名函数，返回一个函数类型，通过f来调用返回的匿名函数，f来调用闭包函数
	f := test0()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
