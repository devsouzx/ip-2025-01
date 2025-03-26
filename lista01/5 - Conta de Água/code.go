package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	valorMetoCubicoComercial float64 = 0.25
	valorMetoCubicoIndustrial float64 = 0.04
)

func main() {
	var conta, consumo int
	var tipoConsumidor string
	var valorDaConta float64

	fmt.Scan(&conta, &consumo, &tipoConsumidor)

	if tipoConsumidor == "R" {
		valorDaConta = 5.0 + (0.05 * float64(consumo))
	} else if tipoConsumidor == "C" {
		valorDaConta = consumoComercial(consumo)
	} else if tipoConsumidor == "I" {
		valorDaConta = consumoIndustrial(consumo)
	} else {
		fmt.Println(errors.New("Tipo de Consumidor invalido"))
		os.Exit(1)
	}

	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valorDaConta)
}

func consumoComercial(consumo int) float64 {
	if consumo <= 80 {
		return 500.0
	}
	return 500.0 + (float64(consumo - 80) * valorMetoCubicoComercial)
}

func consumoIndustrial(consumo int) float64 {
	if consumo <= 100 {
		return 800.0
	}
	return 800.0 + (float64(consumo - 100) * valorMetoCubicoIndustrial)
}