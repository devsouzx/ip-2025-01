package main

import "fmt"

func main() {
	var soma, quantidadeNumeros int
	for i := 50; i <= 70; i++ {
		if i % 2 == 0 {
			soma += i
			quantidadeNumeros += 1
		}
	}
	media := soma / quantidadeNumeros
	fmt.Println("Soma dos numeros pares entre 50 e 70:", soma)
	fmt.Println("media dos numeros pares entre 50 e 70:", media)
}