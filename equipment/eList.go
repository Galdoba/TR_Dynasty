package equipment

//Weapon	TL	Range	Damage	Kg	Cost	Magazine	Power Pack Cost	Traits
//Armour Type	Protection	TL	Rad	Kg	Cost	Required Skill
//Item	TL	Kg	Cost

//ANY	DESCR	TL	KG	COST	SECTION	CATEGORY	BOOK

type itemData struct {
	name        string
	description string
	tl          int
	weight      float64
	section     float64
	category    int
	book        string
}

/*
hull 600t
150t - armor
2t - drive
tl 15


15M
24M
60M
2M
5.3M
150M
0.75M
2M
10M
0.2M
/2
/2
/2
*/
