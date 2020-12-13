package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Assign key values
	emails["Bob"] = "bob@gmail.com"
	emails["dave"] = "dave@gmail.com"
	emails["susan"] = "susan@gmail.com"
	emails["jane"] = "jane@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails["Bob"])
}