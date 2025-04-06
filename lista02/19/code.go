package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var raio, altura float64
	fmt.Print("Escolha a figura (1-cone / 2-cilindro / 3-esfera): ")
	fmt.Scan(&opcao)

	switch opcao {
	case 1:
		fmt.Print("Digite o raio do cone: ")
		fmt.Scan(&raio)
		fmt.Print("Digite a altura do cone: ")
		fmt.Scan(&altura)

		volume := (math.Pi * math.Pow(raio, 2) * altura) / 3
		area := math.Pi * raio * math.Sqrt(math.Pow(raio, 2) + math.Pow(altura, 2))

		fmt.Printf("Volume do cone: %.2f\n", volume)
		fmt.Printf("Área do cone: %.2f\n", area)
	case 2:
		fmt.Print("Digite o raio do cilindro: ")
		fmt.Scan(&raio)
		fmt.Print("Digite a altura do cilindro: ")
		fmt.Scan(&altura)

		volume := math.Pi * math.Pow(raio, 2) * altura
		area := 2 * math.Pi * raio * altura

		fmt.Printf("Volume do cilindro: %.2f\n", volume)
		fmt.Printf("Área do cilindro: %.2f\n", area)
	case 3:
		var raio float64
		fmt.Print("Digite o raio da esfera: ")
		fmt.Scan(&raio)

		volume := (4.0 / 3.0) * math.Pi * math.Pow(raio, 3)
		area := 4 * math.Pi * math.Pow(raio, 2)

		fmt.Printf("Volume da esfera: %.2f\n", volume)
		fmt.Printf("Área da esfera: %.2f\n", area)
	default:
		fmt.Println("Opção inválida.")
	}
}