package cards

const (
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
		for _, rank := range []int{RANK_10, RANK_9, RANK_8, RANK_7, RANK_6, RANK_5, RANK_4, RANK_3, RANK_2} {
			d.cards = append(d.cards, NewCard(suit, rank))
		}
	}
	d.cards = append(d.cards, NewCard(JOKER, JOKER))
	d.cards = append(d.cards, NewCard(JOKER, JOKER))
	//тусуем колоду
	return &d
}

func NewCard(suit, rank int) Card {
	return Card{suit, rank}
}
