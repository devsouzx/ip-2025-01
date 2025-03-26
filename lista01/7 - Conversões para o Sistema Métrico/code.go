package main

import "fmt"

func main() {
	var temperaturaFahrenheit, chuvaPolegadas int
	fmt.Scanln(&temperaturaFahrenheit)
	fmt.Scanln(&chuvaPolegadas)
	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", fahrenheiParaCelsius(temperaturaFahrenheit))
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f", polegadasParaMilimetros(chuvaPolegadas))
}

func fahrenheiParaCelsius(temperaturaFahrenheit int) float64 {
	return ((5.0 * float64(temperaturaFahrenheit)) - 160.0) / 9.0
}

func polegadasParaMilimetros(temperaturaPolegadas int) float64 {
	return float64(temperaturaPolegadas) * 25.4
}