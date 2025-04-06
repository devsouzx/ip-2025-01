package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var data string
	meses := []string{
		"janeiro", "fevereiro", "março", "abril", "maio", "junho",
		"julho", "agosto", "setembro", "outubro", "novembro", "dezembro",
	}
	fmt.Print("Digite a data (dd/MM/yyyy): ")
	fmt.Scan(&data)
	datas := strings.Split(data, "/")
	dia, err1 := strconv.Atoi(datas[0])
	mes, err2 := strconv.Atoi(datas[1])
	ano, err3 := strconv.Atoi(datas[2])

	if err1 != nil || err2 != nil || err3 != nil || mes < 1 || mes > 12 {
		fmt.Println("Data inválida.")
		return
	}

	fmt.Printf("%d de %s de %d", dia, meses[mes-1], ano)
}