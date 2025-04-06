package main

import "fmt"

const (
	valorMetroCubicoResidencial float64 = 0.05
	valorMetroCubicoComercial float64 = 0.25
	valorMetroCubicoIndustrial float64 = 0.04
)

func main() {
	var conta int
	var consumo, valor float64
	var tipo rune

	fmt.Scan(&conta, &consumo, &tipo)

	switch tipo {
	case 'R':
		valor = 5.00 + valorMetroCubicoResidencial*consumo
	case 'C':
		if consumo <= 80 {
			valor = 500.00
		} else {
			valor = 500.00 + valorMetroCubicoComercial*(consumo-80)
		}
	case 'I':
		if consumo <= 100 {
			valor = 800.00
		} else {
			valor = 800.00 + valorMetroCubicoIndustrial*(consumo-100)
		}
	default:
		fmt.Println("Tipo de consumidor inválido")
		return
	}

	fmt.Printf("CONTA = %d\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f\n", valor)
}