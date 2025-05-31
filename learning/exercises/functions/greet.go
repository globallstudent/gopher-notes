package main

import (
	"fmt"
)

func Greet(name string) string {
	return "Hello, " + name
}

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%v", &name)
	fmt.Println(Greet(name))
}
