package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var vetor = make([]int, N)

	for i := range N {
		fmt.Scan(&vetor[i])
	}

	maxLength := 1
	currentLength := 1

	for i := 1; i < N; i++ {
		if vetor[i] > vetor[i-1] {
			currentLength++
		} else {
			if currentLength > maxLength {
				maxLength = currentLength
			}
			currentLength = 1
		}
	}

	if currentLength > maxLength {
		maxLength = currentLength
	}

	fmt.Println(maxLength)
}
