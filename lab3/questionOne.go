package main

import (
	"fmt"
)

type Employee struct {
	FullName string
	Years    int
	Position string
	Income   float64
}

func (e Employee) ReadData() Employee {
	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&e.FullName)

	fmt.Print("Enter Employee Age: ")
	fmt.Scan(&e.Years)

	fmt.Print("Enter Employee Job Role: ")
	fmt.Scan(&e.Position)

	fmt.Print("Enter Employee Salary: ")
	fmt.Scan(&e.Income)

	return e
}

func (e Employee) PrintData() {
	fmt.Printf("\n--- Employee Details ---\n")
	fmt.Printf("Employee Name: %v \n", e.FullName)
	fmt.Printf("Employee Age: %v \n", e.Years)
	fmt.Printf("Job Role: %v \n", e.Position)
	fmt.Printf("Salary: %.2f \n", e.Income)
}

func main() {
	var emp Employee
	emp = emp.ReadData()
	emp.PrintData()
}
