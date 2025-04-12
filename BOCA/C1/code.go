package main

import "fmt"

func main() {
	var n, soma int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i)
		soma += i
	}
	fmt.Printf("\n%d\n", soma)
}
