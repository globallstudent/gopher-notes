package main

import (
	"fmt"
	"strings"
)

func reverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func reverseStringBetterWa(s string) (reversed string) {
	var builder strings.Builder
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}
	return builder.String()
}

func main() {
	fmt.Println(reverseString("Hello"))
	fmt.Println(reverseStringBetterWa("Hello"))
}
