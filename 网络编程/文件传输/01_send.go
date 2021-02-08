package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	//只读方式打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err = ", err)
		return
	}

	defer file.Close()

	//读文件内容
	buf := make([]byte, 4*1024)
	for {
		num, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("file.Read err = ", err)
			}

			return
		}

		//发送内容
		conn.Write(buf[:num])
	}
}

func main() {
	//提示输入文件
	fmt.Println("请输入需要传输的文件: ")
	var path string
	fmt.Scan(&path)

	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err = ", err)
		return
	}

	//主动连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:8000")
	if err1 != nil {
		fmt.Println("net.Dial err1 = ", err1)
		return
	}

	defer conn.Close()

	//给接受方发送文件名
	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("conn.Write err2 = ", err2)
		return
	}

	//接收对方的回复，如果ok,说明对方准备好了，可以发文件
	buf := make([]byte, 1024)
	num, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn.Write Read = ", err3)
		return
	}

	if "ok" == string(buf[:num]) {
		//发送文件内容
		SendFile(path, conn)
	}
}
