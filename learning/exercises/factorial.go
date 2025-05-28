package main

import "fmt"

func factorial(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}

	return fact
}

func main() {
	fmt.Println(factorial(4))
}
