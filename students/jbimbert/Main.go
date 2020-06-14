package main

import (
	"Gophercizes/deck/students/jbimbert/deck"
	"fmt"
)

// custom user comparison function to sort the deck
func compare1(d []deck.Card) func(i, j int) bool {
	return func(i, j int) bool {
		ci, cj := d[i], d[j]
		return int(ci.Suit)*13+int(ci.Rank) > int(cj.Suit)*13+int(cj.Rank)
	}
}

func main() {
	c1 := deck.Card{Rank: deck.V2, Suit: deck.None}
	c2 := deck.Card{Rank: deck.V3, Suit: deck.None}
	d := deck.NewDeck(deck.WithoutCards(c1, c2), deck.WithJockers(4), deck.WithCustomSort(compare1), deck.WithDecks(3))
	fmt.Println(d)
}
