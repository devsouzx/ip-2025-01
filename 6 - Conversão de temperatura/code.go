package main

import "fmt"

func main() {
	var numeroDeCasos int
	fmt.Scanln(&numeroDeCasos)
	temperaturasFahrenheit := make([]int, numeroDeCasos)
	resultados := make([]int, numeroDeCasos)
	for i := 0; i < numeroDeCasos; i++ {
		fmt.Scanln(&temperaturasFahrenheit[i])
	}
	for i := 0; i < len(resultados); i++ {
		temperaturaEmCelsius := float64((5 * (temperaturasFahrenheit[i] - 32)) / 9)
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", float64(temperaturasFahrenheit[i]), temperaturaEmCelsius)
	}
}