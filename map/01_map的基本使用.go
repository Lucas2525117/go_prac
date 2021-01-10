package main

import "fmt"

func main() {
	//map的创建: key为int,value为string
	var m map[int]string
	fmt.Println("m = ", m) //map[]

	//map只有len,没有cap
	fmt.Printf("len(m) = %d\n", len(m))
	//fmt.Printf("cap = %d\n", cap(m)) //invalid argument m (type map[int]string) for cap

	//可以通过make创建
	m1 := make(map[int]string)
	fmt.Println("m1 = ", m1)
	fmt.Printf("len(m1) = %d\n", len(m1))

	//通过make创建 可以指定长度 这里的10相当于指定了容量 长度还是0
	//map容量会自动扩充
	//map是无序的
	m2 := make(map[int]string, 2)
	m2[0] = "lucas"
	m2[0] = "wawa"
	m2[11] = "lili"
	m2[-1] = "yoku"
	m2[-1] = "koko" //当有相同键值时,会替换成最新的
	fmt.Println("m2 = ", m2)
	fmt.Printf("len(m2) = %d\n", len(m2)) //3

	//注意:初始化时键值唯一
	m3 := map[int]string{1: "wawa", 2: "koko", 3: "luis"}
	fmt.Println("m3 = ", m3)
}
