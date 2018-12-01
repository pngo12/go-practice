package main

import "fmt"

func main() {
	// Long Loops
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short loops
	for i:= 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	//FuzzBizz
	for i:= 1; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Println("fizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println("buzz")
		}
	}

}