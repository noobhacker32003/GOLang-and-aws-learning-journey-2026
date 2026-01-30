package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// Function (standalone)
func CalculateArea(r Rectangle) float64 {
	return r.Width * r.Height
}

// Method (attached to type)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{
		Width:  5,
		Height: 10,
	}

	area1 := CalculateArea(rect) // Function call
	area2 := rect.Area()         // Method call

	fmt.Println("Area using function:", area1)
	fmt.Println("Area using method:", area2)
}
