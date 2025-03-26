package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var numeros [3]int
	for i := 0; i < len(numeros); i++ {
		fmt.Scan(&numeros[i])

		if numeros[i] == 0 || numeros[i] > 9 {
			fmt.Println(errors.New("DIGITO INVALIDO"))
			os.Exit(1)
		}
	}
	concatenacao := (numeros[0] * 100) + (numeros[1] * 10) + numeros[2]
	quadrado := concatenacao * concatenacao
	fmt.Printf("%d, %d", concatenacao, quadrado)
}