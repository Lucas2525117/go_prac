package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listen.Close()

	conn, err1 := listen.Accept()
	if err != nil {
		fmt.Println("listen.Accept err1 = ", err1)
		return
	}

	defer conn.Close()

	//Read时需创建切片存储数据
	buf := make([]byte, 4*1024)
	num, err2 := conn.Read(buf)
	if num == 0 {
		fmt.Println("conn.Read err2 = ", err2)
		return
	}

	fmt.Printf("%v", string(buf[:num]))
}
