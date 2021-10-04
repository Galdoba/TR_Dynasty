package cards

import (
	"fmt"
	"testing"
)

func TestCards(t *testing.T) {
	d := NewDeck()
	for _, card := range d.cards {
		fmt.Println(card.String())
	}

}
