/*
1. 分文件编程(多个源文件),必须放在src目录
2. 设置GOPATH环境变量
3. 同一个目录,包名必须一样
4. go env查看go相关的环境路径
5. 同一个目录,调用别的文件的函数,直接调用即可，无需包名引用
*/
package main

func main() {
	test()
}
