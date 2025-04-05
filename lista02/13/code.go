package main

import (
	"fmt"
)

func main() {
	var numero int
	fmt.Scanln(&numero)
	if numero >= 100 && numero <= 999 {
		dezena := (numero / 10) % 10
		fmt.Printf("Algarismo da casa das dezenas: %d", dezena)
	} else {
		fmt.Println("Numero invalido")
	}
}