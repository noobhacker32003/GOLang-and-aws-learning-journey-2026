package main

import "fmt"

func main() {
	// 1. Variables
	fmt.Println("--- Variables ---")

	// Using var keyword
	var name string = "John"
	var age int = 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Type inference
	var country = "USA"
	fmt.Printf("Country: %s\n", country)

	// Short declaration (inside functions only)
	city := "New York"
	fmt.Printf("City: %s\n", city)
}
