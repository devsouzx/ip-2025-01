package main

import (
	"fmt"
	"math"
)

func main() {
	var X float64
	fmt.Scan(&X)
	S := X
	for i := 1; i < 20; i++ {
		if i % 2 == 0 {
			S += math.Pow(X, float64(i)) / fatorial(i)
		} else {
			S -= math.Pow(X, float64(i)) / fatorial(i)
		}
	}

	fmt.Println(S)
}

func fatorial(num int) float64 {
	if num == 0 || num == 1 {
		return 1.0
	}
	fatorial := 1.0
	for i := 2; i <= num; i++ {
		fatorial *= float64(i)
	}
	return fatorial
}
