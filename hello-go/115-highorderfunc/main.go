package main

import "fmt"

func main() {
	// circle()
	highOrderFuncCircle()
}

func highOrderFuncCircle() {
	var query int
	var radius float64

	fmt.Println("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)

	fmt.Println("Enter \n 1. Area \n 2. Perimeter \n 3. Diameter")
	fmt.Scanf("%d", &query)

	printResult(radius, getFunction(query))
}

func printResult(radius float64, calcFunction func(r float64)) {
	calcFunction(radius)
	fmt.Println("Thank You!")
}

func getFunction(query int) func(r float64) {
	query_to_func := map[int]func(r float64){
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return query_to_func[query]
}

// Calculating for Circle without HOF
func circle() {
	var query int
	var radius float64

	fmt.Println("Enter the radius of the circle: ")
	fmt.Scanf("%f", &radius)

	fmt.Println("Enter \n 1. Area \n 2. Perimeter \n 3. Diameter")
	fmt.Scanf("%d", &query)

	if query == 1 {
		calcArea(radius)
	} else if query == 2 {
		calcPerimeter(radius)
	} else if query == 3 {
		calcDiameter(radius)
	} else {
		fmt.Println("Invalid Query")
	}
}

func calcArea(r float64) {
	area := 3.14 * r * r
	fmt.Println("Area is ", area)
}

func calcPerimeter(r float64) {
	perimeter := 2 * 3.14 * r
	fmt.Println("Perimeter is ", perimeter)
}

func calcDiameter(r float64) {
	diameter := 2 * r
	fmt.Println("Diameter is ", diameter)
}
