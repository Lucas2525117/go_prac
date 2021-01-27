package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Stdout.Close() //关闭后 无法输出
	fmt.Println("are you ok?")

	//标准设备文件 默认已经打开，用户可以直接使用
	os.Stdout.WriteString("are you ok?\n")

	//os.Stdin.Close() //关闭后 无法输入
	var a int
	fmt.Println("请输入a:")
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}
