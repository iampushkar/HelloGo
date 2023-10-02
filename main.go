package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World in GO!")
  var name string = "Manish Pushkar"
  fmt.Println("My name is " + name)

  cards := []string {"Ace of Diamonds", newCard()}
  cards = append(cards, "Six of Spades")

  fmt.Println(cards)

  for i, card := range cards {
    fmt.Println(i, card)
  }
}

func newCard() string {
  return "Five of Diamonds"
}
