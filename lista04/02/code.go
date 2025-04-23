package main

import "fmt"

func main() {
	vetor1 := make([]int, 10)
	vetor2 := make([]int, 5)
	var vetorResult1, vetorResult2 []int

	fmt.Println("Digite os 10 numeros do primeiro vetor: ")
	for i := range(vetor1) {
		fmt.Scan(&vetor1[i])
	}

	fmt.Println("Digite os 5 numeros do primeiro vetor: ")
	for i := range(vetor2) {
		fmt.Scan(&vetor2[i])
	}

	for i := range(vetor1) {
		if vetor1[i] % 2 == 0{
			result := vetor1[i]
			for i := range(vetor2) {
				result += vetor2[i]
			}
			vetorResult1 = append(vetorResult1, result)
		} else {
			result := vetor1[i]
			for i := range(vetor2) {
				result += vetor2[i]
			}
			vetorResult2 = append(vetorResult2, result)
		}
	}

	fmt.Println(vetorResult1)
	fmt.Println(vetorResult2)
}