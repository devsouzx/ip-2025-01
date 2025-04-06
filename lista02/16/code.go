package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Print("Digite o valor de A: ")
	fmt.Scan(&a)
	fmt.Print("Digite o valor de B: ")
	fmt.Scan(&b)
	fmt.Print("Digite o valor de C: ")
	fmt.Scan(&c)

	if a == 0 {
		fmt.Println("Nao e equacao do segundo grau")
		return
	}

	delta := b*b - 4*a*c

	if delta > 0 {
		r1 := (-b + math.Sqrt(delta)) / (2 * a)
		r2 := (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Printf("RAÍZES DISTINTAS: %.2f e %.2f\n", r1, r2)
	} else if delta == 0 {
		r := -b / (2 * a)
		fmt.Printf("RAIZ ÚNICA: %.2f\n", r)
	} else {
		parteReal := -b / (2 * a)
		parteImaginaria := math.Sqrt(-delta) / (2 * a)
		fmt.Printf("RAÍZES IMAGINÁRIAS: %.2f + %.2fi e %.2f - %.2fi\n",
			parteReal, parteImaginaria, parteReal, parteImaginaria)
	}
}
