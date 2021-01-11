/*
要点:注意取地址符号 "&"
*/
package main

import "fmt"

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//指针
	var s *Student = &Student{1, "jack", 'm', 18, "beijing"}
	fmt.Println("*s = ", *s)

	//部分初始化
	s1 := &Student{name: "koko"}
	fmt.Printf("s1 = %T\n", s1)
	fmt.Println("s1 = ", *s1)
	fmt.Println("s1 = ", s1)
}
