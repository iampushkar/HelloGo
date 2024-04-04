package main

import "fmt"

func main() {
	fmt.Println("Hello Recursion")
	n := 5
	var result int = factorial(n)
	fmt.Println("Factorial of ", n, "is :", result)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
