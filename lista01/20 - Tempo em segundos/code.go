package main

import "fmt"

func main() {
	var horas, minutos, segundos int
	fmt.Scan(&horas, &minutos, &segundos)
	totalSegundos := (horas * 3600) + (minutos * 60) + segundos
	fmt.Printf("O TEMPO EM SEGUNDOS E = %d", totalSegundos)
}