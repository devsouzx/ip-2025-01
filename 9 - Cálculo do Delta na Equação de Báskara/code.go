package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	delta := float64((b * b) - (4 * a * c))
	fmt.Printf("O VALOR DE DELTA E = %.2f", delta)
}