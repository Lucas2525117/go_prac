package main

import (
	"fmt"
	"net/http"
)

//w:给客户端回复数据
//req:读取客户端回复的数据
func HandleConnect(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req.Method = ", req.Method)
	fmt.Println("req.URL = ", req.URL)
	fmt.Println("req.Header = ", req.Header)
	fmt.Println("req.Body = ", req.Body)

	//给客户端回复数据
	w.Write([]byte("Hello Go"))
}

func main() {
	//注册处理函数
	http.HandleFunc("/", HandleConnect)

	//监听绑定
	http.ListenAndServe(":8000", nil)
}
