package main

import "fmt"

func swap(a, b int) (int, int) {
	temp := a
	a = b
	b = temp
	return a, b
}

func main() {
	var a, b int

	fmt.Print("Enter a and b: ")
	fmt.Scan(&a, &b)
	fmt.Println(swap(a, b))
}
