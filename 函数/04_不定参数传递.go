/*功能: 不定参数传递给另一个不定参数*/
package main

import "fmt"

func test1(args ...int) {
	for _, data := range args {
		fmt.Println("data = ", data)
	}
}

func test0(args ...int) {
	test1(args...) //传递全部

	//传递后两个参数
	test1(args[:2]...) //传递args[0]~args[1]
	test1(args[2:]...) //从args[2]开始传递
}

func main() {
	test0(1, 2, 3, 4)
}
