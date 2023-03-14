package main

import "fmt"

func tes() {
	fruits := []string{"Apple", "Mango", "Durian", "Apple", "Durian", "Banana", "Mango"}

	fruitGroup, fruitAmount := handleFruitGroup(fruits)

	fmt.Println("fruitGroup:", fruitGroup)
	fmt.Println("fruitAmount:", fruitAmount)
}

func handleFruitGroup(fruits []string) (groupResult map[string]int, length int) {
	var result map[string]int = map[string]int{}

	for _, eachFruit := range fruits {
		result[eachFruit]++
	}

	groupResult = result

	length = len(fruits)

	return
}
