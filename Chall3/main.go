package main

import "fmt"

func main() {
	var greeting string = "selamat malam"
	for _, letter := range greeting {
		fmt.Printf("%c \n", letter)
	}
	mapGreeting := map[string]int{
		"s": 1,
		"e": 1,
		"l": 2,
		"a": 4,
		"m": 3,
		"t": 1,
		" ": 1,
	}
	fmt.Println(mapGreeting)
}
