package main

import "fmt"

func main() {
	ids := []int{33,76,49,48,20,4}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum =", sum)

	// Range with map 
	players := map[string]string{"Ronaldo":"Juventus", "Messi":"Barcelona", "Neymar":"PSG", "Salah":"Liverpool"} 

	for k, v := range players {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range players {
		fmt.Println("Name: " + k)
	}
}
