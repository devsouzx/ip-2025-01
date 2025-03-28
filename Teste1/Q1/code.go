package main

import "fmt"

func main() {
	var numero int
	fmt.Scanln(&numero)
	if numero > 0 {
		fmt.Printf("%d é positivo", numero)
	} else if numero < 0 {
		fmt.Printf("%d é negativo", numero)
	} else {
		fmt.Printf("0 é nulo")
	}
}