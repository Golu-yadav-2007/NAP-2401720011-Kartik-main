package main

import "fmt"

func PowerQuestion(base int, exponent int) int {
	if base == 0 {
		return 0
	}
	if exponent == 0 {
		return 1
	}
	
	computedVal := 1
	stepCount := 1

	for stepCount <= exponent {
		computedVal *= base
		stepCount++
	}

	return computedVal
}

func main() {
	result := PowerQuestion(2, 3)
	fmt.Println("Result:", result)
}
