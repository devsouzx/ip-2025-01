package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if triangular(n) {
		fmt.Printf("%d eh triangular\n", n)
	} else {
		fmt.Printf("%d nao eh triangular\n", n)
	}
}

func triangular(n int) bool {
	for i := 1; i*(i+1)*(i+2) <= n; i++ {
		if i*(i+1)*(i+2) == n {
			return true
		}
	}
	return false
}