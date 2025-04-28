package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var vetor = make([]int, N)

	for i := range N {
		fmt.Scan(&vetor[i])
	}

	palindromo := true
	for i := range N/2 {
		if vetor[i] != vetor[N-1-i] {
			palindromo = false
			break
		}
	}

	if palindromo {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}
}
