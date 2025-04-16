package main

import (
	"fmt"
)

func main() {
	var pInicial, d, deltaP, pMin float64
	var qInicial, deltaQ int

	fmt.Scan(&pInicial, &qInicial, &d, &deltaP, &deltaQ, &pMin)

	preco := pInicial
	qtdVendidos := qInicial

	lucroMaximo := -1.0
	var melhorPreco float64
	var melhorQtd int

	fmt.Println("Preco Ingressos Vendidos Lucro")

	for preco >= pMin {
		lucro := preco*float64(qtdVendidos) - d
		fmt.Printf("%.2f %d %.2f\n", preco, qtdVendidos, lucro)

		if lucro > lucroMaximo {
			lucroMaximo = lucro
			melhorPreco = preco
			melhorQtd = qtdVendidos
		}

		preco -= deltaP
		qtdVendidos += deltaQ
	}

	fmt.Printf("Lucro maximo: %.2f\n", lucroMaximo)
	fmt.Printf("Na faixa de preco: %.2f com %d ingressos.\n", melhorPreco, melhorQtd)
}
