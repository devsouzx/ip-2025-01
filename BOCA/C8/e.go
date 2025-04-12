package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scanf("%f", &x)

	for i := 0.0; i <= x+1e-9; i += 0.1 {
		radian := i
		sin := maclaurinSin(radian)
		fmt.Printf("%.1f %.4f\n", radian, sin)
	}
}

func maclaurinSin(x float64) float64 {
	term1 := x
	term2 := math.Pow(x, 3) / float64(factorial(3))
	term3 := math.Pow(x, 5) / float64(factorial(5))
	term4 := math.Pow(x, 7) / float64(factorial(7))
	return term1 - term2 + term3 - term4
}

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}
