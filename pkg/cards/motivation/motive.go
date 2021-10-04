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

func Describe(card cards.Card, descrMap map[cards.Card][]string) []string {
	if len(descrMap[card]) != 2 {
		return []string{card.String(), "Card is not in the Deck"}
	}
	return descrMap[card]
}

func exampleMotivation() map[cards.Card][]string {
	em := make(map[cards.Card][]string)
	em[cards.NewCard(cards.JOKER, cards.JOKER)] = []string{"jkr", "Joker Motive"}
	em[cards.NewCard(cards.SUIT_CLUB, cards.RANK_ACE)] = []string{"A" + cards.STR_SUIT_CLUB, "Club Ace Motive"}
	em[cards.NewCard(cards.SUIT_DIAMOND, cards.RANK_6)] = []string{"6" + cards.STR_SUIT_DIAMOND, "Dimond 6 Motive"}
	return em
}

func StandardMap() map[cards.Card][]string {
	sm := make(map[cards.Card][]string)
	sm[cards.NewCard(cards.JOKER, cards.JOKER)] = []string{"jkr", "Joker Motive"}
	sm[cards.NewCard(cards.SUIT_SPADE, cards.RANK_JACK)] = []string{"Pompous", "NPC is conceited and arrogant in their dealings with others. They consider themselves to be clearly superior to everyone around them, and they make no secret of that conviction."}
	sm[cards.NewCard(cards.SUIT_SPADE, cards.RANK_QUEEN)] = []string{"Ruthless", "NPC will let nothing stand in the way of achieving any goal and feels no consern for the need of others. Such NPSc can feign affection, devotion, sincerity, or anything else that serves their purpose, but actualy they feel nothing."}
	sm[cards.NewCard(cards.SUIT_SPADE, cards.RANK_QUEEN)] = []string{"Deceitful", "NPC has not respect for honesty. Depending upon the referee's decidion, deceitful NPC's may be pathological liars, or they they may use thuth deceitfully, giving just enough information to guarantee that their victims are misled. The actual direction of their deceitfullnes will generally depend upon their secondary motivation. Often, such characters are unable to believe that other people are not lying. They expect to be lied to, and expect thw worst from others."}
	sm[cards.NewCard(cards.SUIT_SPADE, cards.RANK_ACE)] = []string{"Charismatic", "NPC is a charismatic leader to whom others are naturaly drawn. This usualy implies a high CHR attribute and perhaps skill is in Leadership. Some of these NPC are honorable and just; others are cruel and manipulative. The referee can decide based on situation or secondary motivation."}
	sm[cards.NewCard(cards.SUIT_HEART, cards.RANK_JACK)] = []string{"Wise", "NPC is unusually wise, either as a result of years of expirience, or simply because of astute observation. Such NPC almost exhibit good judgment and, if asked, offer sound advice."}
	sm[cards.NewCard(cards.SUIT_HEART, cards.RANK_JACK)] = []string{"Loving", "NPC loves some other person devotedly, perhaps a spouse, parent, child, or close friend. Such NPCs would willingly sacrifice themselves for the one they love. Alternatevly, the NPC mat be loving towards absolutely everyone. The choice is up to the referee."}
	sm[cards.NewCard(cards.SUIT_HEART, cards.RANK_JACK)] = []string{"Honorable", "NPCs are scrupulosly hornest in their dealings with everyone."}
	sm[cards.NewCard(cards.SUIT_HEART, cards.RANK_JACK)] = []string{"Pompous", "NPC is conceited and arrogant in their dealings with others. They consider themselves to be clearly superior to everyone around them, and they make no secret of that conviction."}
	sm[cards.NewCard(cards.SUIT_HEART, cards.RANK_6)] = []string{"6" + cards.STR_SUIT_DIAMOND, "Dimond 6 Motive"}
	return sm
}
