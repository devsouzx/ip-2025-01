package main

import "fmt"

func main() {
	var salario, salarioReajuste float64
	fmt.Scanln(&salario)
	if salario > 300.0 {
		salarioReajuste = salario * 1.3
	} else {
		salarioReajuste = salario * 1.5
	}
	fmt.Printf("SALARIO COM REAJUSTE = %.2f", salarioReajuste)
}