package main

import "fmt"

func main() {
	var numero int
	fmt.Scanln(&numero)
	if numero < 5 || numero > 2000 {
		fmt.Println("Numero inválido")
	} else {
		for i := 1; i <= numero; i++ {
			if i % 2 == 0 {
				fmt.Printf("%dˆ2 = %d\n", i, i*i)
			}
		}
	}
}