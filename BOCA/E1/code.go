package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	alturas := make([]float64, n)
	var soma float64

	for i := range n {
		fmt.Scan(&alturas[i])
		soma += alturas[i]
	}

	media := soma / float64(n)

	for _, altura := range alturas {
		if altura > media {
			fmt.Printf("%.2f\n", altura)
		}
	}
}
