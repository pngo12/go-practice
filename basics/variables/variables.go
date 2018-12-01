package main

import "fmt"

func main() {
	// var name string = "Phillip Ngo"
	// var age int = 100
	// Same as JS es6 Const
	const isCool = true

	// Shorthand variables
	// No need to declare type
	// It is inferred
	//  name := "Phillip"
	 size := 5.0
	 
	 // Combine variable declarations
	 // Similar to JS destructuring ish
	 name, email := "Phillip", "pngo12@gmail.com"


	fmt.Println(name + " " + email)
	fmt.Printf("%T\n", size)
}