package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func RecvFile(filename string, conn net.Conn) {
	//新建文件，file相当于文件句柄(可读可写)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("os.Create err = ", err)
		return
	}

	buf := make([]byte, 4*1024)
	for {
		num, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件接收完毕")
				return
			} else {
				fmt.Println("conn.Read err = ", err)
				return
			}
		}

		//写入内容
		file.Write(buf[:num])
	}
}

func main() {
	//监听
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listen.Close()

	//阻塞用户等待连接
	conn, err1 := listen.Accept()
	if err1 != nil {
		fmt.Println("listen.Accept err1 = ", err1)
		return
	}

	defer conn.Close()

	//读取对方发送的文件名
	buf := make([]byte, 1024)
	num, err2 := conn.Read(buf)
	if err2 != nil {
		fmt.Println("conn.Read err2 = ", err2)
		return
	}

	filename := string(buf[:num])

	//成功接收到文件名 并返回ok
	conn.Write([]byte("ok"))

	//接收文件内容
	RecvFile(filename, conn)
}
