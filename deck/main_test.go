package deck

import (
	"fmt"
	"testing"
)

func TestNewCards(t *testing.T) {
	c := New(Deck(2), Jocker(3), Sorted, Shuffle)
	fmt.Printf("%+v", c.Cards)
	if c != nil {
		t.Errorf("%+v", c)
	}
}

func TestConstLookup(t *testing.T) {
	fmt.Printf("%s", constLookup[A])
}
