package main

import (
  "errors"
  "fmt"
  "math/rand"
  "strings"
  "time"
)

// add is a function that takes two integers and returns their sum.
func add(a, b int) int {
  return a + b
}

// divide is a function that performs division and returns the result.
// It returns an error if the divisor is zero.
func divide(a, b float64) (float64, error) {
  if b == 0 {
    return 0, errors.New("cannot divide by zero")
  }
  return a / b, nil
}

func main() {
  // Variables and Data Types
  var a int = 10
  var b string = "Hello, Go!"
  var c float64 = 3.14
  var d bool = true

  fmt.Println("Variables:", a, b, c, d)

  // Control Structures (if, for, switch)
  x := 10
  if x > 5 {
    fmt.Println("x is greater than 5")
  }

  for i := 0; i < 5; i++ {
    fmt.Println("Loop iteration:", i)
  }

  y := 2
  switch y {
  case 1:
    fmt.Println("One")
  case 2:
    fmt.Println("Two")
  default:
    fmt.Println("Other")
  }

  // Functions
  result := add(3, 5)
  fmt.Println("Result of adding:", result)

  // Packages and Imports
  rand.Seed(time.Now().UnixNano()) // Seed the random number generator
  fmt.Println("Random number:", rand.Intn(100))

  // Error Handling
  resultDivide, err := divide(10, 0)
  if err != nil {
    fmt.Println("Error dividing:", err)
  } else {
    fmt.Println("Result of division:", resultDivide)
  }

  // String Manipulation
  str := "Hello, World!"
  fmt.Println("Original string:", str)
  fmt.Println("Uppercase:", strings.ToUpper(str))
  fmt.Println("Lowercase:", strings.ToLower(str))
  fmt.Println("Contains 'World'?", strings.Contains(str, "World"))
  fmt.Println("Split:", strings.Split(str, ","))
  fmt.Println("Substring:", str[0:5])

  // Arrays and Slices
  arr := [3]int{1, 2, 3}
  slice := []int{4, 5, 6}

  fmt.Println("Array:", arr)
  fmt.Println("Slice:", slice)

  // Maps
  m := make(map[string]int)
  m["a"] = 1
  m["b"] = 2
  m["c"] = 3

  fmt.Println("Map:", m)
  fmt.Println("Value for key 'b':", m["b"])

  // Structs
  type Person struct {
    Name string
    Age  int
  }

  p := Person{Name: "Alice", Age: 30}
  fmt.Println("Struct:", p)
}
