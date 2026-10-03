package main

import "fmt"

type Scholar struct {
	FullName string
	Years    int
	Score    float32
}

func UpdateReference(valPtr *int) {
	*valPtr = 500
}

func QuestionPointers() {
	val := 25
	var addressRef *int = &val

	fmt.Printf("Initial value of val: %v \n", val)
	fmt.Printf("Address of val (&val): %v \n", &val)
	fmt.Printf("Stored address in pointer: %v \n", addressRef)
	fmt.Printf("Dereferenced value (*ptr): %v \n", *addressRef)

	targetNum := 50
	fmt.Printf("Value prior to modification: %v \n", targetNum)
	UpdateReference(&targetNum)
	fmt.Printf("Value post modification: %v \n", targetNum)

	s1 := new(Scholar)
	fmt.Printf("Default Scholar struct state: %v \n", *s1)
	s1.FullName = "Kartik"
	s1.Years = 21
	s1.Score = 95.5
	fmt.Printf("Updated Scholar struct state: %v \n", *s1)
}

func main() {
	QuestionPointers()
}
