package main

import "fmt"

func main() {
	var numberA, numberB int
	fmt.Scanln(&numberA, &numberB)
	if numberA % numberB == 0 {
		fmt.Printf("%d é divisivel por %d", numberA, numberB)
	} else {
		fmt.Printf("%d não é divisivel por %d", numberA, numberB)
	}
}