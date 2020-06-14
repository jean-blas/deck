package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Jocker //Used for Jockers
	None
)

func (s Suit) String() string {
	return [...]string{"spade", "diamond", "club", "heart", "Jocker", "None"}[s]
}

type Rank uint8

const (
	_       = iota
	VA Rank = iota
	V2
	V3
	V4
	V5
	V6
	V7
	V8
	V9
	V10
	VJ
	VQ
	VK
	VJocker
)

func (s Rank) String() string {
	return [...]string{"", "Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Jocker"}[s]
}

type Card struct {
	Rank
	Suit
}

func (c Card) String() string {
	return fmt.Sprintf("(%s of %s)", c.Rank, c.Suit)
}

// FunctionalOption options used to modify the []Card
type FunctionalOption func(c *[]Card)

// WithShuffle shuffle the deck randomly
func WithShuffle() FunctionalOption {
	return func(o *[]Card) {
		c := *o
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(c), func(i, j int) { c[i], c[j] = c[j], c[i] })
	}
}

// WithCustomSort sort the deck with a custom user function
func WithCustomSort(compare func(o []Card) func(i, j int) bool) FunctionalOption {
	return func(o *[]Card) {
		sort.Slice(*o, compare(*o))
	}
}

// WithSort sort the deck in the standard way
func WithSort() FunctionalOption {
	return func(o *[]Card) {
		c := *o
		sort.Slice(*o, func(i, j int) bool {
			if c[i].Suit == c[j].Suit {
				return c[i].Rank < c[j].Rank
			}
			return c[i].Suit < c[j].Suit
		})
	}
}

// WithJockers add n Jockers to the deck
func WithJockers(n int) FunctionalOption {
	return func(o *[]Card) {
		for i := 0; i < n; i++ {
			*o = append(*o, Card{Rank: VJocker, Suit: Jocker})
		}
	}
}

// WithoutCards Suppress some cards from the deck
// to suppress all cards with same Rank, put Suit=Jocker
// example, suppress all V2 => WithoutCards(Card{V2, Jocker})
func WithoutCards(cardsToSuppress ...Card) FunctionalOption {
	return func(o *[]Card) {
		res := make([]Card, 0)
		for _, c := range *o {
			found := false
			for _, bad := range cardsToSuppress {
				if c.Rank == bad.Rank && (bad.Suit == None || c.Suit == bad.Suit) {
					found = true
					break
				}
			}
			if !found {
				res = append(res, c)
			}
		}
		*o = res
	}
}

// WithDecks duplicates the deck n times
func WithDecks(n int) FunctionalOption {
	return func(o *[]Card) {
		if n <= 1 {
			return
		}
		d := make([]Card, 0)
		for i := 0; i < n; i++ {
			d = append(d, *o...)
		}
		*o = d
	}
}

// NewDeck Generate a deck of cards
// May use some options to modify the deck
func NewDeck(opts ...FunctionalOption) []Card {
	var cards []Card
	for s := Spade; s <= Heart; s++ {
		for r := VA; r <= VK; r++ {
			cards = append(cards, Card{Rank: r, Suit: s})
		}
	}
	for _, o := range opts {
		o(&cards)
	}
	return cards
}
