package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Pointers")

	// var num *int
	// fmt.Println("Value of Pointer is : ", num)

	num := 27
	var ptr = &num

	fmt.Println("Memory address - ", ptr)
	fmt.Println("Value stored at that address - ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Modified Value - ", num)

}
