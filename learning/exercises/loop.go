package main

import "fmt"

func main() {

	sum := 0

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("Adding %d to %d\n", sum, i)
			sum += i
		}
	}

	fmt.Println(sum)
}
