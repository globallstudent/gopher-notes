package main

import "fmt"

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	var n int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&n)
	fmt.Println(isEven(n))
}
