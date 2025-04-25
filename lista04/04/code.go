package main

import "fmt"

func main() {
	indicador := make(map[int]int, 10)
	vetor := make([]int, 10)

	for i := range vetor {
		fmt.Scan(&vetor[i])
		indicador[vetor[i]]++
	}

	for i, value := range indicador {
		fmt.Printf("O numero %d se repete %d vezes\n", value, i)
	}
}