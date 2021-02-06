package main

import (
	"fmt"
	"net"
)

func main() {
	//1. 监听
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen err = ", err)
		return
	}

	//关闭
	defer listen.Close()

	//2. 阻塞等待用户连接
	conn, err1 := listen.Accept()
	if err1 != nil {
		fmt.Println("Accept err1 = ", err1)
		return
	}

	//3. 接收用户的请求
	buf := make([]byte, 1024)
	num, err2 := conn.Read(buf)
	if err1 != nil {
		fmt.Println("Read err2 = ", err2)
		return
	}

	fmt.Println("buf = ", string(buf[:num]))

	defer conn.Close()
}
