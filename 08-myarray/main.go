package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Grapes"
	fruitList[3] = "BlueBerry"

	fmt.Println("FruitList is : ", fruitList)
	fmt.Println("Length of FruitList is : ", len(fruitList))

	var veggies = [5]string{"Onion", "Beans", "Paneer"}
	fmt.Println("veggies is : ", veggies)
	fmt.Println("Length of veggies is : ", len(veggies))
}
