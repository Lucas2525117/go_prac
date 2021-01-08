/*
重点: 切片是2倍扩容
*/
package main

import (
	"fmt"
)

func main() {
	//如果超过原来容量，通常以2倍容量扩容
	s := make([]int, 0, 1)
	oldCap := cap(s)
	for i := 0; i < 20; i++ {
		s = append(s, i)
		if newCap := cap(s); oldCap < newCap {
			fmt.Printf("cap:%d ===> %d\n", oldCap, newCap)

			oldCap = newCap
		}
	}
}
