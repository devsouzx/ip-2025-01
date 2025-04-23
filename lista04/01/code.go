package main

import "fmt"

func main() {
	vetor := []int{23, 56, 78, 76, 98, 34, 45, 12, 23, 35}

	for i := range vetor {
		if vetor[i] > 50 {
			fmt.Printf("o numero %d na posicao %d é maior que 50\n", vetor[i], i)
		}
	}
}