package main

import "fmt"

type Deck struct {
	cards []Card
}

type Card struct {
	kind  Suit
	value CardValue
}

type Suit string

const (
	Club    Suit = "club"
	Diamond Suit = "diamond"
	Heart   Suit = "heart"
	Spade   Suit = "spade"
)

type CardValue string

const (
	Two   CardValue = "two"
	Three CardValue = "three"
	Four  CardValue = "four"
	Five  CardValue = "five"
	Six   CardValue = "six"
	Seven CardValue = "seven"
	Eight CardValue = "eight"
	Nine  CardValue = "nine"
	Ten   CardValue = "ten"
	Jack  CardValue = "jack"
	Queen CardValue = "queen"
	King  CardValue = "king"
	Ace   CardValue = "ace"
)

var AllKinds []Suit = []Suit{Club, Diamond, Heart, Spade}
var AllValues []CardValue = []CardValue{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

const AmountOfEach int = 4

func main() {
	fmt.Println("WELCOME TO SOKER")
	var deck = Deck{}
	deck.cards = make([]Card, len(AllValues)*AmountOfEach)
	deckIndex := 0
	for cardValue := range AllValues {
		for i := 0; i < AmountOfEach; i++ {
			deck.cards[deckIndex] = Card{
				kind:  AllKinds[i],
				value: AllValues[cardValue],
			}
			deckIndex++
		}

	}
	fmt.Print(deck.cards)
}
