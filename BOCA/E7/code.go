package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var A = make([]int, N)
	var B = make([]int, N)

	for i := range N {
		fmt.Scan(&A[i])
	}

	for i := range N {
		fmt.Scan(&B[i])
	}

	var produtoEscalar int
	for i := range N {
		produtoEscalar += A[i] * B[i]
	}

	fmt.Println(produtoEscalar)
}
