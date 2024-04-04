package main

import "fmt"

func main() {
	fmt.Println("Welcome to learning Slices")

	var fruitList = []string{"apple", "banana", "grapes", "peach", "oranges"}
	fmt.Println("Fruitlist is : ", fruitList)
	fmt.Printf("Type of fruitlist is %T \n", fruitList)

	for i := 0; i < len(fruitList); i++ {
		fmt.Println(fruitList[i])
	}

	slice_1 := fruitList[1:5]
	fmt.Println("slice 1 - ", slice_1)

	slice := make([]int, 5, 7)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice_11 := append(slice, 32, 65, 9, 12)
	fmt.Println("slice 11 - ", slice_11)
	fmt.Println(len(slice_11))
	fmt.Println(cap(slice_11))

	fmt.Println("Using copy function")
	src_slice := []int{1, 2, 3, 4, 5}
	dest_slice := make([]int, 3)
	fmt.Println("src - ", src_slice)
	fmt.Println("dest - ", dest_slice)
	num := copy(dest_slice, src_slice)
	fmt.Println("Number of copied elements - ", num)

	for index, value := range src_slice {
		fmt.Println(index, "=>", value)
	}

	for _, value := range src_slice {
		fmt.Println("Just value -> ", value)
	}
}
