package main

import "fmt"

func getCreator(os string) string {
	var creator string
	switch os {
	case "linux":
		creator = "Linus Torvalds"
	case "windows":
		creator = "Bill Gates"
	case "mac":
		creator = "A Steve"
	default:
		creator = "Unknown"
	}
	return creator
}

func main() {
	fmt.Println(getCreator("linux"))
}
