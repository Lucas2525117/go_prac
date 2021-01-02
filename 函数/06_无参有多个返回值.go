package main

import "fmt"

//方法1: 变量名不是必须的
func MyFunc() (int, int, int) {
	return 1, 2, 3
}

//方法2
func MyFunc1() (k int, j int, m int) {
	k = 1
	j = 2
	m = 3
	return
}

//方法3
func MyFunc2() (k, j, m int) {
	k, j, m = 1, 2, 3
	return
}

func main() {
	a, b, c := MyFunc()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)

	k, j, m := MyFunc2()
	fmt.Printf("k = %d, j = %d, m = %d\n", k, j, m)
}
