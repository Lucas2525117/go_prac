/*
select: 监听channel上的数据流动
*/
package main

import "fmt"

//ch只写 q只读
func fibonacci(ch chan<- int, q <-chan bool) {
	x, y := 1, 1

	for {
		select {
		case ch <- x: //往ch写数据
			x, y = y, x+y
		case data := <-q: //当q中可以读取到数据时,返回读取到的数据data并打印
			fmt.Println("data = ", data)
			return
		}
	}
}

func main() {
	ch := make(chan int)
	q := make(chan bool)

	//消费者 从ch读取内容
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch //从ch读数据
			fmt.Println("num = ", num)
		}

		//执行完成后 将true写入q
		q <- true
	}()

	//生产者 产生数字 写入ch
	fibonacci(ch, q)
}
