package main

import "fmt"

func main() {
	var valor, lucro float64
	fmt.Scanln(&valor)
	if valor < 10.0 {
		lucro = valor * 0.7
	} else if valor >= 10.0 && valor < 30.0 {
		lucro = valor * 0.5
	} else if valor >= 30.0 && valor < 50.0 {
		lucro = valor * 0.4
	} else {
		lucro = valor * 0.3
	}
	fmt.Printf("O lucro é de %.2f", lucro)
}