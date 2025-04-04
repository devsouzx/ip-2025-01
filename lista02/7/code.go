package main

import (
	"fmt"
	"sort"
)

func main() {
	var valores [3]int
	for i := 0; i < 3; i++ {
		fmt.Scan(&valores[i])
	}
	sort.Ints(valores[:])
	fmt.Println(valores)
}