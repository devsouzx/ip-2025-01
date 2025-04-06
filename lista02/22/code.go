package main

import (
	"fmt"
)

const (
	salarioMinimo float64 = 788.0
	valorHoraExtra float64 = 10.0
)

func main() {
	var matricula, horasExtras int

	fmt.Print("Digite a matrícula do funcionário: ")
	fmt.Scan(&matricula)

	fmt.Print("Digite a quantidade de horas-extras trabalhadas: ")
	fmt.Scan(&horasExtras)

	salarioHoraExtra := float64(horasExtras) * valorHoraExtra

	salarioBruto := 3*salarioMinimo + salarioHoraExtra

	var descontoINSS, descontoIR float64
	if salarioBruto > 1500 {
		descontoINSS = salarioBruto * 0.12
	}
	if salarioBruto > 2000 {
		descontoIR = salarioBruto * 0.20
	}

	salarioLiquido := salarioBruto - descontoINSS - descontoIR

	fmt.Printf("\nSalário Bruto: R$ %.2f\n", salarioBruto)
	fmt.Printf("Salário Líquido: R$ %.2f\n", salarioLiquido)
}