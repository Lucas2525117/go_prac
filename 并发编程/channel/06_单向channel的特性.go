package main

//import "fmt"

func main() {
	//创建一个channel
	ch := make(chan int)

	//双向channel可以隐式转换为单向channel
	var wCh chan<- int = ch //只写
	wCh <- 12
	//<-wCh

	ch <- 12
	var rCh <-chan int = ch //只读
	<-rCh
	//rCh <- 11

	//单向是无法转换为双向的
	//var ch2 chan int = rCh
}
