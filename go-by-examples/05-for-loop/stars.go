package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		star := ""
		for j := 0; j < i; j++ {
			star += "*"
		}
		fmt.Println(star)
	}

	for i := 5; i >= 0; i-- {
		star := ""
		for j := i; j > 0; j-- {
			star += "*"
		}
		fmt.Println(star)
	}

	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
