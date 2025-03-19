package main

import "fmt"

func main() {
	var salario int
	var kwConsumo float64

	fmt.Scanln(&salario)
	fmt.Scanln(&kwConsumo)

	custoDoConsumo := (0.7 * float64(salario)) * (kwConsumo / 100)
	custoPorKw := custoDoConsumo / kwConsumo
	custoComDesconto := custoDoConsumo * 0.9

	fmt.Printf("Custo por kW: R$ %.2f\n", custoPorKw)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoDoConsumo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", custoComDesconto)
}