package main

import (
	"fmt"
)

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}
func calcSquares(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func main() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

// import (
// 	"fmt"
// )

// func producer(chnl chan int) {
// 	for i := 0; i < 10; i++ {
// 		chnl <- i
// 	}
// 	close(chnl)
// }

// func main() {
// 	ch := make(chan int)
// 	go producer(ch)
// 	for v := range ch {
// 		fmt.Println("Received ", v)
// 	}
// }

// import (
// 	"fmt"
// )

// func producer(chnl chan int) {
// 	for i := 0; i < 10; i++ {
// 		chnl <- i
// 	}
// 	close(chnl)
// }

// func main() {
// 	ch := make(chan int)
// 	go producer(ch)
// 	for {
// 		v, ok := <-ch
// 		if ok == false {
// 			break
// 		}
// 		fmt.Println("Received ", v, ok)
// 	}
// }

// import "fmt"

// func sendData(sendch chan<- int) {
// 	sendch <- 10
// }

// func main() {
// 	sendch := make(chan<- int) //创建一个只发送信道

// 	go sendData(sendch)
// 	fmt.Println(<-sendch)
// }

// func main() {
// 	ch := make(chan int)
// 	ch <- 5
// }

// func hello(done chan bool) {
// 	fmt.Println("Hello world goroutine")

// 	done <- true
// }

// func main02() {
// 	done := make(chan bool)

// 	go hello(done)
// 	<-done
// 	//time.Sleep(1 * time.Second)
// 	fmt.Println("main function")
// }

// func main01() {
// 	var a chan int
// 	if a == nil {
// 		fmt.Println("channel a is nil, going to define it")
// 		a = make(chan int)
// 		fmt.Printf("Type of a is %T", a)
// 	}
// }
