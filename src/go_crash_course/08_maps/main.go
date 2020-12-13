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


	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// Declare map and key values short way
	players := map[string]string{"Ronaldo":"Juventus", "Messi":"Barcelona", "Neymar":"PSG", "Salah":"Liverpool"} 
	fmt.Println(players)
	fmt.Println(len(players))
}