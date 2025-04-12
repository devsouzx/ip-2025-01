package main

import "fmt"

func main() {
	var nums []int
	var num int

	for {
		fmt.Scan(&num)
		if num <= 0 {
			break
		}
		nums = append(nums, num)
	}

	for _, n := range nums {
		if quadradoPerfeito(n) {
			fmt.Printf("%d eh quadrado perfeito\n", n)
		} else {
			fmt.Printf("%d nao eh quadrado perfeito\n", n)
		}
	}
}

func quadradoPerfeito(n int) bool {
	for i := 1; i*i <= n; i++ {
		if i*i == n {
			return true
		}
	}
	return false
}