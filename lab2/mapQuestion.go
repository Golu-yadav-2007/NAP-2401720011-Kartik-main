package main

import "fmt"

func main() {
	fmt.Println("Initializing Tech Stack Map")
	techScores := map[string]int{}
	techScores["Go"] = 92
	techScores["Docker"] = 88
	techScores["Kubernetes"] = 75
	fmt.Println(techScores)

	fmt.Println("Map state prior to adding new entry:")
	fmt.Println(techScores)
	techScores["Linux"] = 98
	fmt.Println("Map state after adding new entry:")
	fmt.Println(techScores)

	fmt.Println("Map state prior to deletion:")
	fmt.Println(techScores)
	delete(techScores, "Docker")
	fmt.Println("Map state after deletion:")
	fmt.Println(techScores)

	fmt.Println("Iterating over the map elements:")
	for key, val := range techScores {
		fmt.Println(key, val)
	}

	fmt.Println("Checking for a missing key")
	score, exists := techScores["AWS"]
	fmt.Println(exists, exists, score)
}
