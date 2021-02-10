package main

import (
	"fmt"
	"net/http"
)

func main() {
	//req, err := http.Get("http://www.baidu.com")
	req, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("Get err = ", err)
		return
	}

	defer req.Body.Close()

	fmt.Println("Status = ", req.Status)
	fmt.Println("StatusCode = ", req.StatusCode)
	fmt.Println("Header = ", req.Header)

	buf := make([]byte, 4*1024)
	var str string

	for {
		num, err := req.Body.Read(buf)
		if num == 0 {
			fmt.Println("Read err = ", err)
			break
		}

		str += string(buf[:num])
	}

	fmt.Println("str = ", str)
}
