package main

import "fmt"

func main() {
	fmt.Println("Writing Variadic Functions")
	fmt.Println(sumNumbers()) 
	fmt.Println(sumNumbers(1, 2))
	fmt.Println(sumNumbers(1, 2, 3, 4))
	fmt.Println(sumNumbers(1, 23, 4, 5, 6, 7, 7))
}

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
