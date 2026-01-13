package main

import "fmt"

func add(param1 int, param2 int) {

	sum := param1 + param2
	fmt.Println(sum)

}

func subtract(p1 int, p2 int) int {
	return p1 - p2
}

func main() {
	fmt.Println("Hello, That's my first function")

	param1 := 0
	param2 := 1
	add(param1, param2)

	result := subtract(10, 5)
	fmt.Println(result)
}
