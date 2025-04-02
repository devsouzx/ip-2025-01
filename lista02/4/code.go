package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64
	fmt.Scanln(&number)
	if number < 0 {
		fmt.Printf("O quadrado de %.2f é %.2f", number, number * number)
	} else {
		fmt.Printf("A raiz quadrada de %.2f é %.2f", number, math.Sqrt(float64(number)))
	}
}