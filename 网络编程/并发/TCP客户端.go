package main

import (
	"fmt"
	"net"
	"os"
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

	//从键盘输入内容 给服务器发送内容
	go func() {
		//切片缓存
		buf_key := make([]byte, 1024)

		for {
			num, err := os.Stdin.Read(buf_key)
			if err != nil {
				fmt.Println("os.Stdin.Read err = ", err)
				return
			}

			//发送数据
			conn.Write(buf_key[:num])
			fmt.Println("client send buf = ", string(buf_key[:num]))
		}
	}()

	//接收服务器回复的数据
	buf := make([]byte, 1024)
	for {
		num, err := conn.Read(buf)
		if err != nil { //当读取失败时，主协程直接退出
			fmt.Println("Read err = ", err)
			return
		}

		fmt.Println("client recv buf = ", string(buf[:num]))
	}
}
