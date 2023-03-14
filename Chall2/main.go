package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}

	var j = 0

	for {
		if j == 5 {
			var tajikWord string = "САШАРВО"

			for index, value := range tajikWord {
				fmt.Printf("character %#U starts at byte position %d \n", value, index)
			}
		} else {
			fmt.Println("Nilai j = ", j)
		}
		if j == 10 {
			break
		}

		j++
	}
}
