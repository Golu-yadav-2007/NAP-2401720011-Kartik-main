package main

import (
	"fmt"
	"time"
)

func squareWorker(num int, result chan<- int) {
	fmt.Println("Square Routine started")
	result <- num * num
	fmt.Println("Square Routine completed")
}

func cubeWorker(num int, result chan<- int) {
	fmt.Println("Cube Routine started")
	result <- num * num * num
	fmt.Println("Cube Routine completed")
}

func fibonacciWorker(num int, result chan<- int) {
	fmt.Println("Fibonacci Routine started")
	a, b := 0, 1
	for i := 0; i < num; i++ {
		a, b = b, a+b
	}
	result <- a
	fmt.Println("Fibonacci Routine completed")
}

func main() {
	result := make(chan int)

	// Online compiler ke liye direct value di hai taaki scanf par ruk na jaye
	input := 4 
	fmt.Println("Using Input:", input)

	go squareWorker(input, result)
	go cubeWorker(input, result)
	go fibonacciWorker(input, result)

	for i := 1; i <= 3; i++ {
		fmt.Println("Result:", <-result)
		time.Sleep(time.Millisecond * 500)
	}
}
