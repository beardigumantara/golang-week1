package main

import "fmt"

func main() {

	var a string = "a"

	var b int = 10

	var c bool = true

	_, _, _ = a, b, c

	fmt.Printf(" %s \n %T \n %d \n %T \n %s \n %T \n", a, a, b, b, "c", c)

}

//Buatlah kode yang dapat menghasilkan suatu result di terminal seperti pada comment dibawah ini berdasarkan dari variable yang telah disediakan diatas :)
/**
	a
	string
	10
	int
	c
	bool
**/
