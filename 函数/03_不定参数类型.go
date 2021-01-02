package main

import "fmt"

func MyFunc(args ...int) {
	fmt.Println("len(args) = ", len(args))

	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d] = %d\n", i, args[i])
	}

	for i, data := range args {
		fmt.Printf("args[%d] = %d\n", i, data)
	}
}

//不定参数需要放在形参中的最后一个参数
func MyFunc1(a int, args ...int) {

}

func main() {
	MyFunc(1, 2, 3)
	MyFunc(1, 2)

	//MyFunc1()
}
