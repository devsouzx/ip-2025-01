package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Digite a idade: ")
	fmt.Scan(&idade)
	if idade < 16 {
		fmt.Println("Não-eleitor")
	} else if idade >= 18 && idade <= 65 {
		fmt.Println("Eleitor obrigatório")
	} else {
		fmt.Println("Eleitor facultativo")
	}
}