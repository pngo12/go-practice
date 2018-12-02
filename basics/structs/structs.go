package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// Shorter way 
// type Person struct {
// 	firstname, lastName, city, gender string
// 	age int
// }

// Greeting method (value receiver)
func (p Person) greet() string {
	// The P in person is hotdog, it is similiar to the "this" keyword
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M"{
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Create a person with a struct
	// person1 := Person{firstName: "Phillip", lastName: "Ngo", city: "Orange County", gender: "M", age: 25}
	// Shorter way below
	person1 := Person{"Phillip", "Ngo", "Orange County", "M", 26}
	// fmt.Println(person1)

	// Getting a single field from a struct
	// fmt.Println(person1.firstName)
	
	person1.hasBirthday()
	person1.getMarried("Johnson")
	fmt.Println(person1.greet())
}
