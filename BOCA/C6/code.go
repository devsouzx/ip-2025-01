package main

import "fmt"

func main() {
	var n, aprovados, exames, reprovados int
	var nota1, nota2, media float64
	var mediasAlunos []float64
	fmt.Scan(&n)

	for range n {
		fmt.Scan(&nota1, &nota2)
		media = (nota1 + nota2) / 2
		mediasAlunos = append(mediasAlunos, media)
	}

	for i, media := range mediasAlunos {
		if media <= 3 {
			fmt.Printf("Aluno %d: Reprovado\n", i+1)
			reprovados += 1
		} else if media >= 7 {
			fmt.Printf("Aluno %d: Aprovado\n", i+1)
			aprovados += 1
		} else {
			fmt.Printf("Aluno %d: Exame\n", i+1)
			exames += 1
		}
	}

	fmt.Printf("Total Aprovados: %d\n", aprovados)
	fmt.Printf("Total Exame: %d\n", exames)
	fmt.Printf("Total Reprovados: %d\n", reprovados)
}