package main

import ("fmt"
		"math")

// Define interfaces
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func main() {
	fmt.Println("Hello World")
}