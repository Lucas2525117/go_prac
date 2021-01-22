package main

import (
	"errors"
	"fmt"
)

func testDiv(a, b int) (r int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母为0")
	} else {
		r = a / b
	}

	return
}

func main() {
	c, err := testDiv(4, 0)
	if err != nil {
		fmt.Println("err =", err)
	} else {
		fmt.Println("c =", c)
	}
}
