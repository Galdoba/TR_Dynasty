package cards

import (
	"math/rand"
	"time"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

const (
	STR_SUIT_CLUB    = "\u2663"
	STR_SUIT_HEART   = "\u2665"
	STR_SUIT_SPADE   = "\u2660"
	STR_SUIT_DIAMOND = "\u2666"
	//French
	JOKER        = iota //Джокер
	SUIT_HEART          //Черви
	SUIT_SPADE          //Пики
	SUIT_DIAMOND        //Буби
	SUIT_CLUB           //Трефы
	RANK_ACE            //Туз
	RANK_KING           //Король
	RANK_QUEEN          //Дама
	RANK_JACK           //Валет
	RANK_10
	RANK_9
	RANK_8
	RANK_7
	RANK_6
	RANK_5
	RANK_4
	RANK_3
	RANK_2
)

type Card struct {
	suit int
	rank int
}

type Deck struct {
	cards []Card
}

func NewDeck() *Deck {
	d := Deck{}
	//создаем карты
	for _, suit := range []int{SUIT_HEART, SUIT_SPADE, SUIT_DIAMOND, SUIT_CLUB} {
		for _, rank := range []int{RANK_ACE, RANK_KING, RANK_QUEEN, RANK_JACK, RANK_10, RANK_9, RANK_8, RANK_7, RANK_6, RANK_5, RANK_4, RANK_3, RANK_2} {
			d.cards = append(d.cards, NewCard(suit, rank))
		}
	}
	d.cards = append(d.cards, NewCard(JOKER, JOKER))
	d.cards = append(d.cards, NewCard(JOKER, JOKER))
	//тусуем колоду
	return &d
}

func (d *Deck) AddCard(c Card) {
	d.cards = append(d.cards, c)
}

func NewCard(suit, rank int) Card {
	return Card{suit, rank}
}

func (c *Card) String() string {
	str := ""
	switch c.rank {
	case JOKER:
		str += "Jkr"
	case RANK_ACE:
		str += "A"
	case RANK_KING:
		str += "K"
	case RANK_QUEEN:
		str += "Q"
	case RANK_JACK:
		str += "J"
	case RANK_10:
		str += "10"
	case RANK_9:
		str += "9"
	case RANK_8:
		str += "8"
	case RANK_7:
		str += "7"
	case RANK_6:
		str += "6"
	case RANK_5:
		str += "5"
	case RANK_4:
		str += "4"
	case RANK_3:
		str += "3"
	case RANK_2:
		str += "2"
	}
	switch c.suit {
	case SUIT_CLUB:
		str += "\u2663"
	case SUIT_HEART:
		str += "\u2665"
	case SUIT_SPADE:
		str += "\u2660"
	case SUIT_DIAMOND:
		str += "\u2666"

	}
	return str
}

func (d *Deck) Shuffle() {
	r := dice.New().RollNext("1d50").Sum()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < r; i++ {
		rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
	}
}

func (d *Deck) DrawCard() *Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return &card
}
