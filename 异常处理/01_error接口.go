package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := fmt.Errorf("%s", "this is err1")
	fmt.Println("err1 = ", err1)

	err2 := errors.New("this is err2")
	fmt.Println("err2 = ", err2)
}
