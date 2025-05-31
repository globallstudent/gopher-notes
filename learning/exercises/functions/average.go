package main

import "fmt"

func average(nums ...float64) float64 {
	sum := 0.0

	if len(nums) == 0 {
		return 0.0
	}

	for _, num := range nums {
		sum += num
	}

	return sum / float64(len(nums))
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	max_num := nums[0]
	for _, num := range nums {
		if num > max_num {
			max_num = num
		}
	}

	return max_num
}

func main() {
	fmt.Println("Average of nums 1, 2, 3, 4, 5: ", average(1, 2, 3, 4, 5))
	fmt.Println("Max of nums 1, 2, 3: ", max(1, 2, 3))
}
