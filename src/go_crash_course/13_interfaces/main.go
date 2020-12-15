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

// area of circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// area of rectangcle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	fmt.Println("Hello World")
}