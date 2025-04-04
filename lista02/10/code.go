package main

import (
	"fmt"
	"os"
)

func main() {
	var destino, tipoViagem int
	var preco float64

	fmt.Println("Escolha o destino:")
	fmt.Println("1 - Região Norte")
	fmt.Println("2 - Região Nordeste")
	fmt.Println("3 - Região Centro-Oeste")
	fmt.Println("4 - Região Sul")
	fmt.Print("Digite o numero do destino: ")
	fmt.Scanln(&destino)
	if destino < 1 || destino > 4 {
		fmt.Println("Destino invalido.")
		os.Exit(1)
	}

	fmt.Println("Escolha o tipo de viagem:")
	fmt.Println("1 - Ida e Volta")
	fmt.Println("2 - Só Ida")
	fmt.Print("Digite o número do tipo de viagem: ")
	fmt.Scanln(&tipoViagem)
	if tipoViagem != 1 && tipoViagem != 2 {
		fmt.Println("Tipo de viagem invalido")
		os.Exit(1)
	}

	switch destino {
	case 1:
		if tipoViagem == 1 {
			preco = 900.00
		} else {
			preco = 500.00
		}
	case 2:
		if tipoViagem == 1 {
			preco = 650.00
		} else {
			preco = 350.00
		}
	case 3:
		if tipoViagem == 1 {
			preco = 600.00
		} else {
			preco = 350.00
		}
	case 4:
		if tipoViagem == 1 {
			preco = 550.00
		} else {
			preco = 300.00
		}
	}

	fmt.Printf("O preço da passagem é: R$ %.2f", preco)
}