package main

import "fmt"

// Multiple return values in go

func Names() (string, string) {
	return "Foo", "Bar"
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)

	n3, _ := Names()
	fmt.Println(n3)
}
