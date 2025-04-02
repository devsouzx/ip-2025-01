package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	if number % 2 == 0 {
		fmt.Printf("%d é par", number)
	} else {
		fmt.Printf("%d é impar", number)
	}
}