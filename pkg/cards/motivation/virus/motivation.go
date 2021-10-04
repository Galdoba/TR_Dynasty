package virus

import (
	"github.com/Galdoba/TR_Dynasty/pkg/cards"
)

func MotivationDeck() *cards.Deck {
	d := cards.Deck{}
	//создаем карты
	for _, suit := range []int{cards.SUIT_HEART, cards.SUIT_SPADE, cards.SUIT_DIAMOND, cards.SUIT_CLUB} {
		for _, rank := range []int{cards.RANK_ACE, cards.RANK_KING, cards.RANK_QUEEN, cards.RANK_JACK} {
			d.AddCard(cards.NewCard(suit, rank))
		}
	}
	d.AddCard(cards.NewCard(cards.JOKER, cards.JOKER))
	//тусуем колоду
	d.Shuffle()
	return &d
}

func BeliefStructureMap() map[cards.Card][]string {
	bsm := make(map[cards.Card][]string)
	bsm[cards.NewCard(cards.JOKER, cards.JOKER)] = []string{"Peacemaker", "TODO: Descr"}

	bsm[cards.NewCard(cards.SUIT_SPADE, cards.RANK_ACE)] = []string{"Doomslayer", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_SPADE, cards.RANK_KING)] = []string{"Reproducer", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_SPADE, cards.RANK_QUEEN)] = []string{"Destroyer", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_SPADE, cards.RANK_JACK)] = []string{"Suicide Inducer", "TODO: Descr"}

	bsm[cards.NewCard(cards.SUIT_CLUB, cards.RANK_ACE)] = []string{"Puppeteer", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_CLUB, cards.RANK_KING)] = []string{"Alliance Builder", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_CLUB, cards.RANK_QUEEN)] = []string{"Empire Builder", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_CLUB, cards.RANK_JACK)] = []string{"Reproducing Doomslayer", "TODO: Descr"}

	bsm[cards.NewCard(cards.SUIT_DIAMOND, cards.RANK_ACE)] = []string{"Hobbyist", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_DIAMOND, cards.RANK_KING)] = []string{"Naturalist", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_DIAMOND, cards.RANK_QUEEN)] = []string{"Explorer", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_DIAMOND, cards.RANK_JACK)] = []string{"Parent", "TODO: Descr"}

	bsm[cards.NewCard(cards.SUIT_HEART, cards.RANK_ACE)] = []string{"Prophet", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_HEART, cards.RANK_KING)] = []string{"Priest", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_HEART, cards.RANK_QUEEN)] = []string{"God", "TODO: Descr"}
	bsm[cards.NewCard(cards.SUIT_HEART, cards.RANK_JACK)] = []string{"Mother", "TODO: Descr"}

	return bsm
}
