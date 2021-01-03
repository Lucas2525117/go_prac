package main

import "fmt"

func main() {
	a := 1
	str := "abc"

	func() {
		//闭包以"引用"的方式捕获外部变量
		a = 2
		str = "def"
		fmt.Printf("内部: a = %d, str = %s\n", a, str)
	}() //()代表直接调用
	fmt.Printf("外部: a = %d, str = %s\n", a, str)
}
