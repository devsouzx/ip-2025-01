package main

import (
	"fmt"
	"strings"
)

const (
	precoArCondicionado = 1750.00
	precoPinturaMetalica = 800.00
	precoVidroEletrico = 1200.00
	precoDirecaoHidraulica = 2000.00
)

func main() {
	var arCondicionado, pinturaMetalica, vidroEletrico, direcaoHidraulica string
	var precoInicial float64

	fmt.Print("Digite o preço de fábrica do carro: ")
	fmt.Scan(&precoInicial)

	precoFinal := precoInicial

	fmt.Printf("Deseja Ar condicionado (R$ %.2f)? (s/n): ", precoArCondicionado)
	fmt.Scan(&arCondicionado)
	if strings.ToUpper(arCondicionado) == "S" {
		precoFinal += precoArCondicionado
	}

	fmt.Printf("Deseja Pintura metálica (R$ %.2f)? (s/n): ", precoPinturaMetalica)
	fmt.Scan(&pinturaMetalica)
	if strings.ToUpper(pinturaMetalica) == "S" {
		precoFinal += precoPinturaMetalica
	}
	
	fmt.Printf("Deseja Vidro elétrico (R$ %.2f)? (s/n): ", precoVidroEletrico)
	fmt.Scan(&vidroEletrico)
	if strings.ToUpper(vidroEletrico) == "S" {
		precoFinal += precoVidroEletrico
	}
	
	fmt.Printf("Deseja Direção hidráulica (R$ %.2f)? (s/n): ", precoDirecaoHidraulica)
	fmt.Scan(&direcaoHidraulica)
	if strings.ToUpper(direcaoHidraulica) == "S" {
		precoFinal += precoDirecaoHidraulica
	}

	fmt.Printf("Preço final do carro: R$ %.2f", precoFinal)
}