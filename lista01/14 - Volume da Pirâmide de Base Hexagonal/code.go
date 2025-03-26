package main

import (
	"fmt"
	"math"
)

func main() {
	var altura, aresta float64
	fmt.Scanln(&altura)
	fmt.Scanln(&aresta)
	areaBase := (3 * math.Pow(aresta, 2) * math.Sqrt(3)) / 2
	volumePiramide := (areaBase * altura) / 3
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS", volumePiramide)
}