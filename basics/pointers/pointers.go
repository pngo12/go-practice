package main

import "fmt"

func main(){
	a := 5
	b := &a

	fmt.Println(a,b)
	fmt.Printf("%T\n", a)

	// Use * to read value from adress
	fmt.Println(*b)

	// Change the value with pointer
	*b = 10
	fmt.Println(a)
}