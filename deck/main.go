package deck

import (
	_ "fmt"
	"math/rand"
	"sort"
)

type Suite string

const (
	Spade        Suite = "Spade"
	Heart        Suite = "Heart"
	Diamond      Suite = "Diamond"
	Club         Suite = "Club"
	SUITE_Jocker Suite = "Jocker"
)

type Rank int

const (
	A Rank = iota + 1
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suite Suite
	Rank  Rank
}

var RankToString = map[Rank]string{
	A:     "A",
	One:   "1",
	Two:   "2",
	Three: "3",
	Four:  "4",
	Five:  "5",
	Six:   "6",
	Seven: "7",
	Eight: "8",
	Nine:  "9",
	Ten:   "10",
	Jack:  "J",
	Queen: "Q",
	King:  "K",
}

func (c *Card) String() string {
	return fmt.Sprintf("[%s --> %s],", c.Suite, RankToString[c.Rank])
}

type Cards struct {
	Cards       []Card
	CountOfDeck int
}

func New(options ...func(*Cards) error) *Cards {
	// default deck of size 1
	c := Cards{Cards: deck(), CountOfDeck: 1}

	for _, option := range options {
		option(&c)
	}
	return &c
}

func Deck(count int) func(*Cards) error {
	return func(c *Cards) error {
		return c.addDeck(count)
	}
}

func Jocker(count int) func(*Cards) error {
	return func(c *Cards) error {
		return c.addJocker(count)
	}
}

func Sorted(c *Cards) error {
	sort.Sort(c)
	return nil
}

func (c *Cards) Len() int {
	return len(c.Cards)
}

func (c *Cards) Swap(i, j int) {
	c.Cards[i], c.Cards[j] = c.Cards[j], c.Cards[i]
}

func (c *Cards) Less(i, j int) bool {
	c1 := c.Cards[i]
	c2 := c.Cards[j]
	f := func(suite Suite) int {
		var value int
		switch suite {
		case Spade:
			value = 1
			break
		case Heart:
			value = 2
			break
		case Diamond:
			value = 3
			break
		case Club:
			value = 4
			break
		case SUITE_Jocker:
			value = 5
			break
		}
		return value
	}
	if f(c1.Suite) < f(c2.Suite) {
		return true
	}
	if c1.Rank < c2.Rank {
		return true
	}
	return false
}

func Shuffle(c *Cards) error {
	c.Shuffle()
	return nil
}

func (c *Cards) Shuffle() error {
	rand.Shuffle(len(c.Cards), c.Swap)
	return nil
}

func deck() []Card {
	var deck []Card
	suite := []Suite{Spade, Heart, Diamond, Club}
	rank := []Rank{A, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}
	for _, s := range suite {
		for _, r := range rank {
			card := Card{Suite: s, Rank: r}
			deck = append(deck, card)
		}
	}
	return deck
}

func (c *Cards) addDeck(count int) error {
	currentDeckSize := c.CountOfDeck
	for i := currentDeckSize; i <= count; i++ {
		newDeck := deck()
		c.Cards = append(c.Cards, newDeck...)
		c.CountOfDeck++
	}
	return nil

}

func (c *Cards) addJocker(count int) error {
	for i := 0; i < count; i++ {
		newCard := Card{Suite: SUITE_Jocker, Rank: Rank(-1)}
		c.Cards = append(c.Cards, newCard)
	}
	return nil
}
