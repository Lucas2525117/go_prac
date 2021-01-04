package main

import "fmt"

func swap(p, q *int) {
	*p, *q = *q, *p
	fmt.Printf("*p = %d, *q = %d\n", *p, *q)
}

func main() {
	a, b := 10, 20
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}
