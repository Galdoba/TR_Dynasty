package motivation

import "github.com/Galdoba/TR_Dynasty/pkg/cards"

type Motive struct {
	motive      cards.Card
	meaning     string
	description string
}

func New() *Motive {
	m := Motive{}
	//d := cards.NewDeck()
	return &m
}

func CallMotiveDesc(card cards.Card, descrMap map[cards.Card]string) string {
	return descrMap[card]
}
