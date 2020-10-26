package deck

import (
    _ "fmt"
)

type Suite string

const (
	Spade   Suite = "Spade"
	Heart   Suite = "Heart"
	Diamond Suite = "Diamond"
	Club    Suite = "Club"
)

type Rank string

const (
	A Rank = "A"
	One Rank = "1"
	Two Rank = "2"
	Three Rank = "3"
	Four Rank = "4"
	Five Rank = "5"
	Six Rank = "6"
	Seven Rank = "7"
	Eight Rank = "8"
	Nine Rank = "9"
	Ten Rank = "10"
	Jack Rank = "Jack"
	Queen Rank = "Queen"
	King Rank = "King"
)

type Card struct {
	Suite Suite
	Rank  Rank
}

type Cards []Card

func New(countOfDeck int) []Card {
	var cards []Card
	for i := countOfDeck; i > 0; i-- {
		deck := generateDeck()
		cards = append(cards, deck...)
	}
	return cards
}

func generateDeck() []Card {
	var deck []Card
	suite := []Suite{Spade, Heart, Diamond , Club}
	rank := []Rank{A, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	for _, s := range suite {
		for _, r := range rank {
			card := Card{Suite: s, Rank: r}
			deck = append(deck, card)
		}
	}
	return deck
}

