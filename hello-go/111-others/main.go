package main

import "fmt"

func main() {
	fmt.Println("Hello Maps")
	codes := map[string]string{"en": "English", "fr": "French"}
	codes["hin"] = "Hindi"
	fmt.Println(codes)
	delete(codes, "fr")
	fmt.Println(codes)

	my_map := make(map[string]string)
	my_map["mp"] = "MP"
	my_map["ld"] = "LD"
	fmt.Println(my_map)
	val, found := my_map["mp"]
	fmt.Println(val, found)

	sum := addNum(32, 23)
	fmt.Println("Sum of 2 nums : ", sum)

	fmt.Println("-----------Operations-----------")
	sum, diff, mul, div := operation(123, 43)
	fmt.Println("Operation : ", sum, diff, mul, div)
}

func addNum(a int, b int) int {
	sum := a + b
	return sum
}

func operation(a int, b int) (sum int, diff int, mul int, div int) {
	sum = a + b
	diff = a - b
	mul = a * b
	div = a / b
	return
}
