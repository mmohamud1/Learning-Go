package main

import "fmt"

func main() {
	// GO DATA TYPES
	// string
	// bool
	// int
	// int int8 int16 in32 int64
	// uint uint8 uint16 uint32 uint 64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var keyword
	var name = "Mohamed"
	var age = 24

	// Using const 
	var isCool = true

	// show name and age in terminal
	fmt.Println(name, age, isCool)
	// get type of variable e.g string
	fmt.Printf("%T\n", name)
}