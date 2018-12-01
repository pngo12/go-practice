package main

import "fmt"

func main(){
	// Define a map
	// emails := make(map[string]string)

	// Assign key values

	// emails["Joe"] = "Joemail@joe.com"
	// emails["Sharon"] = "sharon@mail.com"
	// emails["Bill"] = "bill@gmail.com"

	// // Delete from a map
	// delete(emails, "Joe")

	emails := map[string]string{"Bob":"bob@bob.com", "joe":"joemail.com"}

	fmt.Println(emails)
}