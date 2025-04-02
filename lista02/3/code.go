package main

import "fmt"

func main() {
	var number1, number2 int
	fmt.Scanln(&number1, &number2)
	soma := number1 + number2
	if soma <= 20 {
		fmt.Println(soma - 5)
	} else {
		fmt.Println(soma + 8)
	}
}