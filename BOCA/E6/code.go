package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	var vetor = make([]int, N)

	for i := range N {
		fmt.Scan(&vetor[i])
	}

	sort.Ints(vetor)

	for i := range N {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(vetor[i])
	}
	fmt.Println()
}
