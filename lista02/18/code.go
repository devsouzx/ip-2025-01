package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var precoNormal float64
	fmt.Print("Digite o preço normal do DVD alugado: ")
	fmt.Scanln(&precoNormal)

	var lancamento string
	fmt.Print("Digite se o DVD é lançamento (s/n): ")
	fmt.Scan(&lancamento)

	weekDay := time.Now().Weekday()

	desconto := 0.0
	if weekDay == time.Monday || weekDay == time.Tuesday || weekDay == time.Thursday {
		desconto = 0.4
	}

	acresimo := 0.0
	if strings.ToLower(lancamento) == "s" {
		acresimo = 0.15
	}
	
	precoDesconto := precoNormal * (1 - desconto)
	precoFinal := precoDesconto * (1 + acresimo)
	fmt.Printf("Preço final da locação: R$ %.2f\n", precoFinal)
}