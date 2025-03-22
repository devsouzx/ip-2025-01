package main

import "fmt"

func main() {
	matriz := [2][2]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scanln(&matriz[i][j])
		}
	}
	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f", det(matriz))
}

func det(matriz [2][2]int) float64 {
	return float64((matriz[0][0] * matriz[1][1]) - (matriz[0][1] * matriz[1][0]))
}