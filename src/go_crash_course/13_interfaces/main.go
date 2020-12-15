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

// area of shape
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f\n", circle)
	fmt.Printf("Rectangle Area: %f\n", rectangle)

}