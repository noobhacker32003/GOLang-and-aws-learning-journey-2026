package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

// Method to calculate area
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method to calculate perimeter
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {
	rect := Rectangle{Width: 10.0, Height: 5.0}

	area := rect.Area()
	perimeter := rect.Perimeter()

	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)
}
