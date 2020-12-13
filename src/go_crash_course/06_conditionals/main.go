package main

import "fmt"

func main() {
	x := 10
	y := 15

	// If else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else  {
		fmt.Printf("%d is greater than %d\n", x, y)
	}

	// Else if
	color := "blue"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue or red")
	}

	// Switch 
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or red")	
	}

}