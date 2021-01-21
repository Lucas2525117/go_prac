/*
要点:空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口。
*/
package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "go"
	i[2] = Student{"koko", 12}

	//类型查询/断言
	//第一个返回下标，第二个返回下标对应的值
	for index, data := range i {
		//第一个返回的是值，第二个返回判断结果的真假
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d]类型为int,内容为%d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d]类型为string,内容为%s\n", index, value)
		} else if value, ok := data.(Student); ok == true {
			fmt.Printf("x[%d]Student,内容为name = %s, age = %d\n", index, value.name, value.age)
		}
	}
}
