package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		classificacao string
		idade int
	)
	fmt.Print("Digite a idade: ")
	fmt.Scan(&idade)
	if idade >= 0 && idade <= 2 {
		classificacao = "Recém-nascido"
	} else if idade >= 3 && idade <= 11 {
		classificacao = "Criança"
	} else if idade >= 12 && idade <= 19 {
		classificacao = "Adolescente"
	} else if idade >= 20 && idade <= 55 {
		classificacao = "Adulto"
	} else if idade > 55 {
		classificacao = "Idoso"
	} else {
		fmt.Println("Idade invalida")
		os.Exit(1)
	}
	fmt.Println(classificacao)
}