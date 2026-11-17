package main

import "testing"

func TestAddItem(t *testing.T) {
	items = nil

	err := AddItem(Item{1, "Laptop", 5, 50000})

	if err != nil || len(items) != 1 {
		t.Fatal("Add Item Failed")
	}
}

func TestUpdateItem(t *testing.T) {
	items = []Item{{1, "Laptop", 5, 50000}}

	err := UpdateItem(1, 10, 60000)

	if err != nil || items[0].Quantity != 10 {
		t.Fatal("Update Failed")
	}
}

func TestRemoveItem(t *testing.T) {
	items = []Item{{1, "Laptop", 5, 50000}}

	err := RemoveItem(1)

	if err != nil || len(items) != 0 {
		t.Fatal("Remove Failed")
	}
}
