package main

import "fmt"

func main() {
	m := map[int]string{1: "lucas", 2: "wawa"}
	fmt.Println("修改前 m = ", m)

	//如果已存在key值,则修改
	m[1] = "jack"
	fmt.Println("修改后 m = ", m)

	m[3] = "koko" //m追加时，map底层会自动扩容(类似于append)
	fmt.Println("m = ", m)
}
