package main

import "fmt"

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

func main() {
	var items []Item

	items = append(items, Item{1, "Laptop", 5, 50000})
	items = append(items, Item{2, "Mouse", 10, 500})

	fmt.Println("Inventory Items:")

	for _, item := range items {
		fmt.Println(item)
	}

	fmt.Println("Total Items:", len(items))
}
