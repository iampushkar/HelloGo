package main

import "fmt"

func main() {
	fmt.Println("Anonymous Functions")

	x := func(l int, b int) int {
		return l * b
	}(20, 30)
	fmt.Println("Result of x :", x)
	fmt.Printf("%T \n", x)
}
