package deck

import (
	"testing"
)

func TestNewCards(t *testing.T) {
	c := New(Deck(2), Jocker(3), Sorted, Shuffle)
	if c != nil {
		t.Errorf("%+v", c)
	}
}
