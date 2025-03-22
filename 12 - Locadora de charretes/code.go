package main

import "fmt"

func main() {
	var horas int
	fmt.Scanln(&horas)
	valor := float64((horas / 3) * 10)
	if horas % 3 != 0 {
		valor += float64(horas % 3) * 5
	}
	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valor)
}