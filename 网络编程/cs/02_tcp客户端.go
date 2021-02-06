package main

import (
	"fmt"
	"net"
)

func main() {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial err = ", err)
		return
	}

	//关闭连接
	defer conn.Close()

	//发送数据
	num, err1 := conn.Write([]byte("are you ok?"))
	if err1 != nil {
		fmt.Println("Write err1 = ", err1)
		return
	}

	fmt.Println("send num = ", num)
}
