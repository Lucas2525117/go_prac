/*
map的删除:delete
*/
package main

import "fmt"

func main() {
	m := map[int]string{1: "lucas", 2: "wawa", 3: "jack"}
	fmt.Println("m = ", m)

	//删除key为1的内容
	delete(m, 1)
	fmt.Println("m = ", m)
}
