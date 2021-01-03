/*
src: 放源代码
bin: 放可执行程序
pkg: 放平台相关的库

如果有多个文件或多个包
1) 配置GOPATH变量 src同级目录的绝对路径
2) 自动生成bin或pkg,使用go install命令
   除了要配置GOPATH环境变量,还要配置GOBIN环境变量
*/

package main

import (
	//_ "calc" //_操作是引入该包，不直接使用包里面的函数，而是调用该包里面的init函数
	"calc"
	"fmt"
)

func init() {
	fmt.Println("this is main init")
}

func main() {
	a := calc.Add(1, 2)
	fmt.Println("a = ", a)
}
