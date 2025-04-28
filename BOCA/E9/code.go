package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var vetor = make([]int, N)
	var resultado = make([]int, N)

	for i := range N {
		fmt.Scan(&vetor[i])
	}

	resultado[0] = vetor[0]
	for i := 1; i < N; i++ {
		resultado[i] = resultado[i-1] + vetor[i]
	}

	for i := range N {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(resultado[i])
	}
	fmt.Println()
}
