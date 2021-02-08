package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("usage:xxx file")
		return
	}

	//获取文件属性
	info, err := os.Stat(list[1])
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Println("name = ", info.Name())
	fmt.Println("size = ", info.Size())
}
