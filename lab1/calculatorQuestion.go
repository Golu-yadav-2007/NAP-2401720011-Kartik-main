package lab1

import "fmt"

func CalculatorQuestion() {

	for {
		var choice int
		fmt.Println("=== CALCULATOR MENU ===")
		fmt.Println("1. Integer Calculations")
		fmt.Println("2. Float Calculations")
		fmt.Println("3. Exit Program")
		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		if choice == 1 {
			var a int
			fmt.Print("Enter first number (int): ")
			fmt.Scan(&a)

			var b int
			fmt.Print("Enter second number (int): ")
			fmt.Scan(&b)

			var addition int = a + b
			fmt.Printf("Sum: %d \n", addition)

			var subtraction int = a - b
			fmt.Printf("Difference: %d \n", subtraction)

			if b == 0 {
				fmt.Println("Error: Cannot divide by zero!")
			} else {
				var division int = a / b
				fmt.Printf("Quotient: %d \n", division)
			}
		} else if choice == 2 {
			var x float32
			fmt.Print("Enter first number (float): ")
			fmt.Scan(&x)

			var y float32
			fmt.Print("Enter second number (float): ")
			fmt.Scan(&y)

			var add float32 = x + y
			fmt.Printf("Sum: %.2f \n", add)

			var sub float32 = x - y
			fmt.Printf("Difference: %.2f \n", sub)

			if y == 0 {
				fmt.Println("Error: Cannot divide by zero!")
			} else {
				var div float32 = x / y
				fmt.Printf("Quotient: %.2f \n", div)
			}
			var multi float32 = x * y
			fmt.Printf("Product: %.2f \n", multi)
		} else if choice == 3 {
			fmt.Println("Exiting calculator...")
			break
		} else {
			fmt.Println("Invalid choice, please try again.")
		}
	}

}
