package main

import "fmt"

func main() {
	// var fruitArr [2]string

	// Assign values to fruitArr

	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare an array and assign it immediately
	fruitArr := []string{"Apple, Orange", "Banana"}

	fmt.Println(fruitArr)
	//Print the length
	fmt.Println(len(fruitArr))
}