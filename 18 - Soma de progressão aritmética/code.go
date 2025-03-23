package main

import "fmt"

func main() {
	var a1, r, n, soma int
	fmt.Scan(&a1, &r, &n)
	for i := 0; i < n; i++ {
		soma += a1 + i*r
	}
	fmt.Println(soma)
}