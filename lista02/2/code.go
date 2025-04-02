package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	if number > 0 {
		fmt.Printf("%d é positivo", number)
	} else if number < 0 {
		fmt.Printf("%d é negativo", number)
	} else {
		fmt.Printf("%d é nulo", number)
	}
}