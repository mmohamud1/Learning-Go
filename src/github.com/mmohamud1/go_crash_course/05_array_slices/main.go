package main

import "fmt"

func main() {
	// Arrays declare
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	teamArr := [2] string{"Liverpool", "Arsenal"}

	fmt.Println(fruitArr)
	fmt.Println(fruitArr[1])
	fmt.Println(teamArr)

}