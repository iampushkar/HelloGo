package main

import "fmt"

var jwtToken string = "66762828"

const LoginToken string = "hvjgvchvbjgfukbuv"

func main() {
	var userName string = "Manish"
	fmt.Println(userName)
	fmt.Printf("Variable is of type : %T \n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var myNum uint32 = 245
	fmt.Println(myNum)
	fmt.Printf("Variable is of type : %T \n", myNum)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type : %T \n", anotherString)

	// implicit type
	var website = "manishpushkar.me"
	fmt.Println(website)

	// no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(jwtToken)
	fmt.Println(LoginToken)
}
