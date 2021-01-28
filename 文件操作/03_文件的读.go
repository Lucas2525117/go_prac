package main

import (
	"bufio"
	"fmt"
	"io"
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

//文件读
func ReadFile(path string) {
	//打开文件 "Open"函数
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer file.Close()

	//创建切片
	buf := make([]byte, 1024*2)

	//num为读取的字节数目
	num, err1 := file.Read(buf)
	if err1 != nil && err1 != io.EOF { //文件出错 没有到结尾
		fmt.Println("err1 = ", err1)
		return
	}

	//buf[:num]为指定读取的大小
	fmt.Println("buf = ", string(buf[:num]))
}

//每次读取一行
func ReadFileLine(path string) {
	//打开文件 "Open"函数
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer file.Close()

	//新建一个缓冲区，把内容先放在缓冲区
	r := bufio.NewReader(file)

	for {
		//遇到'\n'结束读取，但是'\n'也读取进入
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("err = ", err)
		}

		fmt.Printf("buf = #%s#\n", string(buf))
	}
}

func main() {
	path := "D:/test.txt"

	//WriteFile(path)
	//ReadFile(path)
	ReadFileLine(path)
}
