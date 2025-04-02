package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)
	if number % 5 == 0 {
		fmt.Printf("%d é divisivel por 5", number)
	} else {
		fmt.Printf("%d não é divisivel por 5", number)
	}
}