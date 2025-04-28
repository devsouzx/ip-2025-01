package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	valores := make([]float64, n)
	for i := range n {
		fmt.Scan(&valores[i])
	}

	xmin, xmax := valores[0], valores[0]
	for _, v := range valores {
		if v < xmin {
			xmin = v
		}
		if v > xmax {
			xmax = v
		}
	}

	for i, v := range valores {
		var xnorm float64
		if xmax == xmin {
			xnorm = 0.0
		} else {
			xnorm = (v - xmin) / (xmax - xmin)
		}
		fmt.Printf("%.2f", xnorm)
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
