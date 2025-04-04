package main

import "fmt"

func main() {
	var x float64
	fmt.Scanln(&x)
	fmt.Printf("%.2f", f(x))
}

func f(x float64) float64 {
	return 8 / (2 - x)
}