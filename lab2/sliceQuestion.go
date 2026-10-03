package main

import (
	"fmt"
	"slices"
)

func SliceQuestion() {
	fmt.Println("Initializing Integer Slice")
	nums := []int{45, 60, 75}
	fmt.Println(nums)

	fmt.Println("State of slice prior to appending:")
	fmt.Println(nums)
	nums = append(nums, 90)
	fmt.Println("State of slice after appending:")
	fmt.Println(nums)

	fmt.Println("State of slice prior to deletion:")
	fmt.Println(nums)
	nums = slices.Delete(nums, 1, 2)
	fmt.Println("State of slice after deletion:")
	fmt.Println(nums)

	fmt.Println("State of slice prior to element modification:")
	fmt.Println(nums)
	nums[0] = 10
	fmt.Println("State of slice after modification:")
	fmt.Println(nums)
}

func main() {
	SliceQuestion()
}
