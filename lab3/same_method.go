package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	rect := Rectangle{Width: 12.0, Height: 6.0}
	circle := Circle{Radius: 10.0}

	fmt.Println("Rectangle area:", rect.Area())
	fmt.Println("Circle area:", circle.Area())
}
