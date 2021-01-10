/*
map作函数参数为引用传递
*/
package main

import "fmt"

func ChangeMap(m map[int]string) {
	m[1] = "fofo"
	//delete(m, 1)
}

func main() {
	m := map[int]string{1: "lucas", 2: "wawa", 3: "jack"}
	fmt.Println("m = ", m)

	ChangeMap(m)
	fmt.Println("m = ", m)
}
