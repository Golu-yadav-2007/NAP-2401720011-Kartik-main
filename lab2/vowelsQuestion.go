package main

import "fmt"

func VowelsQuestion(text string) int {
	vowelCounter := 0
	for _, ch := range text {
		switch ch {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowelCounter++
		}
	}
	return vowelCounter
}

func main() {
	ans := VowelsQuestion("Kartik")
	fmt.Println("Total Vowels:", ans)
}
