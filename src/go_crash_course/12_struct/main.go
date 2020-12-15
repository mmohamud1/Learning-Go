package main

import ("fmt"
		"strconv")

// Define Person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string 
	age int
}

// Greeting method {value reciever}
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I live in " + p.city + " and I am " + strconv.Itoa(p.age) + " years old." 
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouceLastName string) {
	if p.gender == "f" {
		return
	} else {
		p.lastName = spouceLastName
	}

}

func main() {
	// init person using struct
	person1 := Person{firstName: "Mohamed", lastName: "Mohamud", city: "London", gender: "M", age: 24}
	// Aletrnative
	person2 := Person{"David", "Beckham", "Miami", "M", 45}

	fmt.Println(person1, person2)

	// Get single field
	fmt.Println(person1.firstName)
	fmt.Println(person2.age)


	person1.hasBirthday()
	person1.getMarried("Williams")

	fmt.Println(person1.greet())
}