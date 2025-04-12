package main

import "fmt"

func main() {
	var B, E int
	resultado := 1
	fmt.Scan(&B, &E)
	if E == 0 {
		fmt.Println(1)
	} else {
		for i := 0; i < E; i++ {
			resultado *= B
		}
		fmt.Println(resultado)
	}
}
