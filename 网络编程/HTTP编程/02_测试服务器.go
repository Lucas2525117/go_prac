package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go")
}

func main() {
	http.HandleFunc("/go", myHandler)

	//在指定地址进行监听，开启一个HTTP
	http.ListenAndServe("127.0.0.1:8000", nil)
}
