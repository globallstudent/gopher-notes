package main

import "fmt"

func printArray(a [3]int) {
	fmt.Println("Inside function: ", a)
	a[0] = 99
}

func main() {
	printArray([3]int{1, 2, 3})
}
