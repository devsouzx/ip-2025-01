package main

import (
	"fmt"
)

func main() {
	var id int
	var nota1, nota2, nota3, mediaExercicos float64

	fmt.Print("Digite o número de identificação do aluno: ")
	fmt.Scan(&id)

	fmt.Print("Digite suas 3 notas: ")
	fmt.Scan(&nota1, &nota2, &nota3)

	fmt.Print("Digite sua média dos exercícios: ")
	fmt.Scan(&mediaExercicos)

	mediaFinal := (nota1 + nota2*2 + nota3*3 + mediaExercicos) / 7

	var conceito string
	var status string

	switch {
	case mediaFinal >= 9.0 && mediaFinal <= 10.0:
		conceito = "A"
		status = "APROVADO"
	case mediaFinal >= 7.5:
		conceito = "B"
		status = "APROVADO"
	case mediaFinal >= 6.0:
		conceito = "C"
		status = "APROVADO"
	case mediaFinal >= 4.0:
		conceito = "D"
		status = "REPROVADO"
	default:
		conceito = "E"
		status = "REPROVADO"
	}

	fmt.Printf("\nNúmero do aluno: %d\n", id)
	fmt.Printf("Notas: %.2f, %.2f, %.2f\n", nota1, nota2, nota3)
	fmt.Printf("Média dos exercícios: %.2f\n", mediaExercicos)
	fmt.Printf("Média de aproveitamento: %.2f\n", mediaFinal)
	fmt.Printf("Conceito: %s\n", conceito)
	fmt.Printf("Situação: %s\n", status)
}
