package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	if isPremiumUser {
		finalCost = bulkMessageCost * (1 - discountRate)
	} else {
		finalCost = bulkMessageCost
	}

	if finalCost <= accountBalance {
		fmt.Println(purchaseSuccessMessage)
		accountBalance -= finalCost
	} else {
		fmt.Println(insufficientFundMessage)
	}

	fmt.Println("Account balance:", accountBalance)
}
