package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Dial err = ", err)
		return
	}

	defer conn.Close()

	//可以添加自己的http请求报文
	request := ""
	conn.Write([]byte(request))

	buf := make([]byte, 4*1024)
	num, err1 := conn.Read(buf)
	if num == 0 {
		fmt.Println("err1 = ", err1)
		return
	}

	fmt.Printf("%v", string(buf[:num]))
}
