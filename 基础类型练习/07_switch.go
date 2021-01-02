package main

import "fmt"

func main() {
	switch num := 1; num {
	case 1:
		fmt.Print("1")
		//fallthrough //不跳出switch语句 后面的无条件执行
	case 2:
		fmt.Print("2")
		//fallthrough
	case 3, 4, 5:
		fmt.Print("3")
		//fallthrough
	default:
		fmt.Println("-1")
		//fallthrough
	}

	score := 85
	switch {
	case score > 90:
		fmt.Print("优秀")
	case score > 80:
		fmt.Print("良好")
	case score > 60:
		fmt.Print("合格")
	default:
		fmt.Println("不合格")
	}
}
