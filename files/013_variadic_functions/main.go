package main

import (
	"fmt"
)

func sum(nums ...int) int {
	fmt.Println(nums)
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	totalSum := sum(1, 2)
	fmt.Println("total", totalSum)
	totalSum = sum(1, 2, 3)
	fmt.Println("total", totalSum)

	nums := []int{1, 2, 3, 4, 5}
	totalSum = sum(nums...)
	fmt.Println("total", totalSum)
}
