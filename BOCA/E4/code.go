package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	var vetor = make([]float64, N)

	for i := range N {
		fmt.Scan(&vetor[i])
	}

	var somatorio float64 = 0

	for i := range N/2 {
		somatorio += math.Pow(vetor[i]-vetor[N-1-i], 3)
	}

	fmt.Printf("%.2f\n", somatorio)
}
