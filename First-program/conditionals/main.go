package main

import (
	"fmt"
	"time"
)

func main() {
	// 2. Conditional Statements (If/Else)
	fmt.Println("--- If/Else ---")
	age := 20
	if age >= 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// If with a short statement
	if n := 10; n > 5 {
		fmt.Println("n is greater than 5")
	}

	// 3. Switch Statements
	fmt.Println("\n--- Switch ---")
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Friday":
		fmt.Println("Weekend is coming")
	default:
		fmt.Println("Regular day")
	}

	// Switch without expression (acts like if-else chain)
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}
