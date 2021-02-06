package main

import (
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	//自动关闭连接
	defer conn.Close()

	//获取客户端网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println(addr, " connect success!")

	buf := make([]byte, 2*1024)

	for {
		//读取用户数据
		num, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err = ", err)
			return
		}
		fmt.Printf("[%s]:read buf = %s", addr, string(buf[:num]))

		//TCP知识: 按下RETURN键后产生的UNIX系统中的换行符,服务器发送给客户端的是:回车和换行字符( CR/LF )，即 "\r\n"，它们的作用是将光标回移到左边并移动到下一行
		//基于上面的解释:回送的字节会多2 (此方法只是临时解决)
		if string(buf[:num-2]) == "exit" {
			fmt.Println(addr, " close!")
			return
		}

		//把数据转换为大写再给用户发送
		conn.Write([]byte(strings.ToUpper(string(buf[:num]))))
	}
}

func main() {
	//1. 监听
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen err = ", err)
		return
	}

	//关闭
	defer listen.Close()

	//接收多个用户
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept err = ", err)
			return
		}

		//处理用户的请求(新建协程)
		go HandleConn(conn)
	}
}
