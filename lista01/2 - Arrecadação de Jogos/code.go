package main

import "fmt"

const (
	valorPopular float64 = 1.0
	valorGeral float64 = 5.0
	valorArquibancada float64 = 10.0
	valorCadeira float64 = 20.0
)

func main() {
	var numeroDeCasos, totalIngressos int
	var percPopular, percGeral, percArquibancada, percCadeiras float64

	fmt.Scanln(&numeroDeCasos)

	arrecadacoes := make([]string, numeroDeCasos)

	for i := 0; i < numeroDeCasos; i++ {
		fmt.Scan(&totalIngressos, &percPopular, &percGeral, &percArquibancada, &percCadeiras)

		rendaDoJogo := (float64(totalIngressos) * (percPopular / 100) * valorPopular) + (float64(totalIngressos) * (percGeral / 100) * valorGeral) + (float64(totalIngressos) * (percArquibancada / 100) * valorArquibancada) + (float64(totalIngressos) * (percCadeiras / 100) * valorCadeira)

		arrecadacoes[i] = fmt.Sprintf("A RENDA DO JOGO N. %d E = %.2f", i+1, rendaDoJogo)
	}

	for i := 0; i < numeroDeCasos; i++ {
		fmt.Println(arrecadacoes[i])
	}
}