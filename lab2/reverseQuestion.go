package main

import "fmt"

func ReverseQuestion(inputStr string) {
	reversedStr := ""

	index := len(inputStr) - 1
	for index >= 0 {
		reversedStr += string(inputStr[index])
		index--
	}

	fmt.Println("Reversed String Output:", reversedStr)
}

func main() {
	ReverseQuestion("Kartik")
}
