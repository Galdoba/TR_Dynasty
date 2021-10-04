package motivation

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/cards"
	"github.com/Galdoba/TR_Dynasty/pkg/cards/motivation/virus"
)

func TestMtvProcess(t *testing.T) {
	motList := virus.BeliefStructureMap()
	var cardList []cards.Card
	cardList = append(cardList, cards.NewCard(cards.SUIT_DIAMOND, cards.RANK_6))
	cardList = append(cardList, cards.NewCard(cards.SUIT_CLUB, cards.RANK_ACE))
	cardList = append(cardList, cards.NewCard(cards.SUIT_CLUB, cards.RANK_QUEEN))
	for _, val := range cardList {
		fmt.Println(Describe(val, motList))
	}
}
