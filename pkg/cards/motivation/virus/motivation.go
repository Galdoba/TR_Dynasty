package virus

import (
	"github.com/Galdoba/TR_Dynasty/pkg/cards"
	"github.com/Galdoba/TR_Dynasty/pkg/cards/motivation"
)

type VirusMotivation struct {
	m motivation.Motive
}

func (vm *VirusMotivation) Primary() string {
	//virusMotiveMap := virusMotivationMap()
	return "virusMotiveMap[vm.m.Primary()]"
}

func BeliefStructure(card cards.Card) string {
	bs := "Undetermined"
	switch card {
	case cards.NewCard(cards.JOKER, cards.JOKER):
		bs = "Peacemaker"
	}
	return bs
}
