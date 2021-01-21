/*
要点:空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口。
*/
package main

import "fmt"

func main() {
	//空接口 万能类型 可以保存任意类型的值(类似c++中的纯虚函数)
	var i interface{} = 1
	fmt.Println("i = ", i)

	i = "abc"
	fmt.Println("i = ", i)

	i = false
	fmt.Println("i = ", i)

	//从空接口获取值
	a := 1
	var m interface{} = a
	fmt.Println("m = ", m)

	//n := m

	//空接口比较
	var p interface{} = 10
	var q interface{} = 9

	fmt.Println("p == q is", p == q)

	//不能比较空接口中的动态值
	// var c interface{} = []int{10} //[]int为不可比较的类型
	// var d interface{} = []int{20}

	// fmt.Println("c == d is", c == d)

}
