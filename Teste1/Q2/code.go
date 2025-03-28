package main

import "fmt"

func main() {
	var numero int
	fmt.Scanln(&numero)
	if numero >= 20 && numero <= 90 {
		fmt.Printf("%d está entre 20 e 90", numero)
	} else {
		fmt.Printf("%d não está entre 20 e 90", numero)
	}
}