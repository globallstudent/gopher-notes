package main

import "fmt"

func main() {
	// initialize variables here
	var (
		smsSendingLimit int
		costPerSMS      float64
		hasPermission   bool
		username        string
	)

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
