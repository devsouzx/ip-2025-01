package main

import (
	"fmt"
)

func main() {
	var preco float64
	var opcao int

	fmt.Print("Digite o preço do produto: R$ ")
	fmt.Scan(&preco)

	fmt.Println("Escolha a condição de pagamento:")
	fmt.Println("1 - À vista (dinheiro ou cheque) - 10% de desconto")
	fmt.Println("2 - À vista (cartão) - 5% de desconto")
	fmt.Println("3 - Em 2 vezes (sem juros)")
	fmt.Println("4 - Em 3 vezes (com 10% de juros)")
	fmt.Print("Digite a opção: ")
	fmt.Scan(&opcao)

	switch opcao {
	case 1:
		total := preco * 0.90
		fmt.Printf("Total a pagar com 10%% de desconto: R$ %.2f\n", total)
	case 2:
		total := preco * 0.95
		fmt.Printf("Total a pagar com 5%% de desconto: R$ %.2f\n", total)
	case 3:
		parcela := preco / 2
		fmt.Printf("Pagamento em 2x de R$ %.2f (sem juros). Total: R$ %.2f\n", parcela, preco)
	case 4:
		total := preco * 1.10
		parcela := total / 3
		fmt.Printf("Pagamento em 3x de R$ %.2f (com 10%% de juros). Total: R$ %.2f\n", parcela, total)
	default:
		fmt.Println("Opçao de pagamento inválida")
	}
}
