package main

import "fmt"

func main() {
	var base, expoente int

	fmt.Print("Digite a base: ")
	fmt.Scan(&base)
	fmt.Print("Digite o expoente: ")
	fmt.Scan(&expoente)

	resultado := 1
	for range expoente {
		resultado *= base
	}

	fmt.Printf("%d elevado a %d é igual a %d", base, expoente, resultado)
}