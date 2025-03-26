package main

import (
	"fmt"
	"math"
)

func main() {
	var raio, altura float64
	fmt.Scan(&raio)
	fmt.Scan(&altura)
	custo := calcularAreaCilindro(raio, altura) * 100.0
	fmt.Printf("O VALOR DO CUSTO E = %.2f", custo)
}

func calcularAreaCilindro(raio float64, altura float64) float64 {
	return (2 * math.Pi * math.Pow(raio, 2)) + (2 * math.Pi * raio * altura)
}