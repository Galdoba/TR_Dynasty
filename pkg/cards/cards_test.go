package cards

import (
	"fmt"
	"testing"
)

func TestCards(t *testing.T) {
	d := NewDeck()
	fmt.Println(d)
	d.Shuffle()
	fmt.Println(d)
	card := d.DrawCard()
	fmt.Println(card)
	fmt.Println(d)

}
