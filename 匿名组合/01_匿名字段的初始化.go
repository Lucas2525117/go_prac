package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型 没有名字:匿名字段 继承了Person里面的成员
	id     int
	addr   string
}

func main() {
	//顺序初始化
	var s Student = Student{Person{"jack", 'n', 18}, 1, "shanghai"}
	fmt.Println("s = ", s)

	//自动推导类型
	s1 := Student{Person{"koko", 'n', 18}, 2, "shanghai"}
	fmt.Println("s1 = ", s1)
	//%+v显示更详细的信息
	fmt.Printf("s1 = %+v\n", s1)

	//指定成员初始化,没有初始化的成员自动赋值为0
	s2 := Student{id: 2}
	fmt.Printf("s2 = %+v\n", s2)

	s3 := Student{Person: Person{name: "lucas"}, id: 2}
	fmt.Printf("s3 = %+v\n", s3)
}
