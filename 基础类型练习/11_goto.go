package main

import (
	"fmt"
)

func main() {
	fmt.Println("test1")

	goto End //可以用在任何地方

	fmt.Println("test2")

End:
	fmt.Println("test end")
}
