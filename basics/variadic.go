package main

import "fmt"

func sum(nums ...int) int {
	sum_ := 0
	for _, val := range(nums) {
		sum_ += val
	}
	fmt.Println("Sum of", nums, ":", sum_)
	return sum_
}

func main() {
	sum(124, 124, 124, 52, 123)
	sum(124)
	sum(124, 45, 76, 23)
}
