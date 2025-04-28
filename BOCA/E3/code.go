package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	vetor := make([]int, n)
	for i := range n {
		fmt.Scan(&vetor[i])
	}

	novo := make([]int, n)

	if n == 1 {
		novo[0] = 0
	} else {
		for i := range n {
			if i == 0 {
				novo[i] = vetor[i+1]
			} else if i == n-1 {
				novo[i] = vetor[i-1]
			} else {
				novo[i] = vetor[i-1] + vetor[i+1]
			}
		}
	}

	for i, v := range novo {
		fmt.Print(v)
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
