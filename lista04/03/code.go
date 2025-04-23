package main

import "fmt"

func main() {
	vetor := make([]int, 10)
	var numerosPares, numerosImpares []int
	var somaPares, quantidadeImpares int

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite o %d° numero: ", i+1)
		fmt.Scan(&vetor[i])
		if vetor[i] % 2 == 0{
			numerosPares = append(numerosPares, vetor[i])
			somaPares += vetor[i]
		} else {
			numerosImpares = append(numerosImpares, vetor[i])
			quantidadeImpares += 1
		}
	}

	fmt.Println("números pares digitados:", numerosPares)
	fmt.Println("soma dos números pares digitados:", somaPares)
	fmt.Println("números ímpares digitados:", numerosImpares)
	fmt.Println("quantidade de números ímpares digitados:", quantidadeImpares)
}