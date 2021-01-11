/*
要点: 首字母小写，只能在同一包里面被使用
*/
package test

import "fmt"

type stu struct {
	id int
}

type Stu struct {
	Id int
}

func myFunc() {
	fmt.Println("this is myFunc")
}

func MyFunc() {
	fmt.Println("this is MyFunc")
}
