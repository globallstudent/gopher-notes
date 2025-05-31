package main

import "fmt"

func Square(n int) int {
	return n * n
}

func main() {
	var n int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&n)
	fmt.Println(Square(n))
}
