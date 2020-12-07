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
	var size float32 = 2.3
	
	// Using const 
	const isCool = true
	
	// shorthand
	nameTwo := "Mo"
	ageTwo := "32"
	sizeTwo := 1.3

	job, numOfYears := "develper", 1

	// show name and age in terminal
	fmt.Println(name, age, isCool, nameTwo, ageTwo, sizeTwo, job, numOfYears)
	// get type of variable e.g string
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", size)

}