package main

import (
	"fmt"
	"net"
	"strings"
	//"time"
)

type Client struct {
	channel chan string //用于发送数据的管道
	Name    string      //用户名
	Addr    string      //网络地址
}

var mapOnline map[string]Client

var message = make(chan string)

func MsgTrans() {
	//给map分配空间
	mapOnline = make(map[string]Client)

	for {
		msg := <-message //无消息之前会阻塞

		//遍历map
		for _, client := range mapOnline {
			client.channel <- msg
		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.channel {
		//给当前客户端发送信息
		conn.Write([]byte(msg + "\n"))
	}
}

func CreateMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg

	return
}

func HandleConn(conn net.Conn) {
	//获取客户端网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个结构体
	cli := Client{make(chan string), cliAddr, cliAddr}

	//把结构体添加到map
	mapOnline[cliAddr] = cli

	//给当前客户端发送信息
	go WriteMsgToClient(cli, conn)

	//广播某用户在线
	message <- CreateMsg(cli, "login")

	//用户是否退出
	bQuit := make(chan bool)
	//对方是否连接
	bConnect := make(chan bool)

	//接收用户发送的数据
	go func() {
		buf := make([]byte, 2*1024)

		for {
			num, err := conn.Read(buf)
			if num == 0 {
				bQuit <- true

				fmt.Println("conn.Read err = ", err)
				return
			}

			msg := string(buf[:num-1])
			if len(msg) == 3 && msg == "who" {
				//遍历map 给当前用户发送所有在线成员
				conn.Write([]byte("user list:\n"))
				for _, data := range mapOnline {
					msg := data.Addr + ":" + data.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				//重命名
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				mapOnline[cliAddr] = cli //重新添加进map

				conn.Write([]byte("rename ok!\n"))
			} else {
				//转发此内容
				message <- CreateMsg(cli, msg)
			}

			bConnect <- true
		}
	}()

	//for循环防止程序退出
	for {
		//通过select检测channel的变化
		select {
		case <-bQuit:
			//当前用户从map中移除
			delete(mapOnline, cliAddr)
			message <- CreateMsg(cli, "login out") //广播用户下线
			return
		case <-bConnect:
			continue
			// case <-time.After(5 * time.Second): //60s超时 -----备注:超时存在问题 暂时屏蔽掉
			// 	delete(mapOnline, cliAddr)
			// 	message <- CreateMsg(cli, "time out, leave out") //广播用户下线
			// 	return
		}
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

	//转发消息
	go MsgTrans()

	//主协程 阻塞等待用户连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err = ", err)
			continue
		}

		defer conn.Close()

		//处理用户连接
		go HandleConn(conn)
	}
}
