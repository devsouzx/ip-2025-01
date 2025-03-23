package main

import "fmt"

func main() {
	var numero int
	fmt.Scanln(&numero)
	soma := 0.0
	for i := 1; i <= numero; i++ {
		soma += 1.0 / float64(i)
	}
	fmt.Printf("%.6f\n", soma)
}