package lab1

import "fmt"

func TypeConvertionQuestion() {

	fmt.Println("Welcome, Kartik!")

	var num int
	// var decimal float32
	// var text string
	// var flag bool
	fmt.Println("Default integer value:", num)

	var originalInt int = 45
	var convertedFloat float32 = float32(originalInt)
	fmt.Printf("Converted int to float: %.2f\n", convertedFloat)

	var asciiCode int = 66
	var convertedChar string = string(asciiCode)
	fmt.Printf("Converted int to string/char: %s\n", convertedChar)

	// sampleText := "golang"
	// fmt.Println("Length of text:", len(sampleText))
	// // len([]rune(sampleText))

}
