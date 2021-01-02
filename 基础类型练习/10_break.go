package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for {
		i++
		time.Sleep(time.Second)

		if i == 5 {
			break
			//continue
		}

		fmt.Println("i = ", i)
	}
}
