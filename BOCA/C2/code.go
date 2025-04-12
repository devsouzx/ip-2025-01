package main

import "fmt"

func main() {
	var l, r, soma, numerosPares, media int
	fmt.Scan(&l, &r)
	for i := l; i <= r; i++ {
		if i % 2 == 0 {
			soma += i
			numerosPares += 1
		}
	}
	if numerosPares != 0 {
		media = soma / numerosPares
	}
	fmt.Println(soma)
	fmt.Println(media)
}
