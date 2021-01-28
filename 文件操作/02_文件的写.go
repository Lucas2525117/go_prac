package main

import (
	"fmt"
	"os"
)

func WriteFile(path string) {
	//打开文件 新建文件
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//文件关闭
	defer file.Close()

	var buf string

	for i := 0; i < 10; i++ {
		buf = fmt.Sprintf("i = %d\n", i) //buf作为返回值 返回具体内容
		num, err := file.WriteString(buf)
		if err != nil {
			fmt.Println("err = ", err)
			return
		}
		fmt.Println("num = ", num)
	}
}

func main() {
	path := "D:/test.txt"

	WriteFile(path)
}
