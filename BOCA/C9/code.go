package main

import (
	"fmt"
	"math"
)

func fatorial(n int) float64 {
	result := 1.0
	for i := 2; i <= n; i++ {
		result *= float64(i)
	}
	return result
}

func cosTaylor(x float64) float64 {
	result := 0.0
	for i := 0; i < 20; i++ {
		term := math.Pow(-1, float64(i)) * math.Pow(x, float64(2*i)) / fatorial(2*i)
		result += term
	}
	return result
}

func main() {
	var x float64
	fmt.Scan(&x)

	valorCalculado := cosTaylor(x)
	valorFuncao := math.Cos(x)
	diferenca := valorCalculado - valorFuncao

	fmt.Printf("%.4f %.4f %.4f\n", valorCalculado, valorFuncao, diferenca)
}
