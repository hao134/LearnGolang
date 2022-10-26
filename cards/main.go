package main

import "fmt"

var deckSize int

func main() {
  cards := []string{"Ace of Diamonds",newCard()}
  cards = append(cards, "Six of spades")

  fmt.Println(cards)
  for i, card := range cards {
	fmt.Println(i, card)
  }
}

func newCard() string {
	return "Five of Diamonds"
}

