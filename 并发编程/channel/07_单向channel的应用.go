package main

import "fmt"

//此通道只能写不能读
func Producer(wCh chan<- int) {
	for i := 0; i < 10; i++ {
		wCh <- i * i
	}
	close(wCh)
}

//只能读 不能写
func Consumer(rCh <-chan int) {
	for num := range rCh { //rCh没有数据会阻塞
		fmt.Println("num = ", num)
	}
}

func main() {
	//创建一个双向通道
	ch := make(chan int)

	//生产者
	go Producer(ch) //channel传参是引用传递

	//消费者
	Consumer(ch)
}
