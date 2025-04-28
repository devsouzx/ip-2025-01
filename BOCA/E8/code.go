package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var vetor = make([]int, N)
	distinct := make(map[int]bool)

	for i := range N {
		fmt.Scan(&vetor[i])
		distinct[vetor[i]] = true
	}

	fmt.Println(len(distinct))
}
