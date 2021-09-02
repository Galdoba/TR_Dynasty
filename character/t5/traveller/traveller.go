package traveller

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/T5/assets"
	"github.com/Galdoba/TR_Dynasty/T5/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/core/calendar"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

type TravellerT5 struct {
	notHuman        bool
	randomHomeWorld bool
	characteristic  map[int]assets.Characteristic
	skills          map[int]assets.Skill
	knowledges      map[int]assets.Knowledge
	homeworld       *wrld.World
	birthdate       *calendar.ImperialDate
	currentDate     *calendar.ImperialDate
	err             error
}

func NewTravellerT5() *TravellerT5 {
	trv := TravellerT5{}
	trv.currentDate = calendar.NewImperialDate(calendar.GameStartDay)
	trv.randomHomeWorld = true
	trv.characteristic = make(map[int]assets.Characteristic)
	trv.skills = make(map[int]assets.Skill)
	trv.knowledges = make(map[int]assets.Knowledge)
	//trv.err = fmt.Errorf("generation of traveller not implemented")
	trv.GenerateCharactiristics()
	trv.GenerateHomeworld()

	fmt.Println(trv)
	return &trv
}

func (trv *TravellerT5) GenerateCharactiristics() {
	switch trv.notHuman {
	case true:
		trv.err = fmt.Errorf("generation of non-Human not implemented")
	case false:
		charList := []int{
			assets.Strength,
			assets.Dexterity,
			assets.Endurance,
			assets.Intelligence,
			assets.Education,
			assets.Social,
			assets.Psionics,
			assets.Sanity,
		}
		for _, val := range charList {
			trv.characteristic[val] = *assets.NewCharacteristic(val, 2)
			if trv.characteristic[val].Err != nil {
				trv.err = trv.characteristic[val].Err
			}
		}
	}
}

func (trv *TravellerT5) UPP() string {
	upp := ""
	for _, code := range assets.ListCharacteristics() {
		if val, ok := trv.characteristic[code]; ok {
			upp += ehex.New().Set(val.Value()).String()
		}
	}
	return upp
}

func (trv *TravellerT5) GeneticProfile() string {
	gp := ""
	for _, code := range assets.ListCharacteristics() {
		if val, ok := trv.characteristic[code]; ok {
			if val.PositionCode() == assets.C5 || val.PositionCode() == assets.C6 {
				gp += "X"
				continue
			}
			gp += ehex.New().Set(val.Genetics()).String()
		}
	}
	return gp
}

type CardPrinter interface {
	PrintCard()
}

type charcterCard struct {
	/*
		Eneri Dinsha ABA69A. Genetic 4553XX
		Homeworld: Regina (1910 Spinward Marches)
		Trader-1, Animals-1, Bureaucrat-1
		Psychology-4, Robotics-2, Astrogator-2, Pilot-4. Strategy-1
		Fleet Tactics-1, Computer-1, Counsellor,-1 Diplomat-1.
		Imperial Navy Lieutenant O3.
		Age 31. Born 069-1074 (note it has been pushed back again)
		MCUF-1. XS-1. WB-1.
		Current Date: 001-1105
	*/
	data      *TravellerT5
	name      string
	upp       string
	gp        string
	homeworld string
	skills    []string
	career    []string
	age       int
	birthdate string
	rewards   []string
	curDate   string
}

func newCard(trv *TravellerT5) *charcterCard {
	cc := charcterCard{}
	cc.data = trv
	cc.name = "Eneri Dirshar"
	cc.upp = trv.UPP()
	cc.gp = trv.GeneticProfile()
	cc.homeworld = trv.homeworld.Name() + " (" + trv.homeworld.Hex() + " " + trv.homeworld.Sector() + ")"
	cc.curDate = trv.currentDate.String()
	cc.skills = trv.homeworld.TradeCodes()
	for _, val := range trv.skills {
		cc.skills = append(cc.skills, val.String())
	}
	return &cc
}

func (cc charcterCard) PrintCard() {
	fmt.Printf("%v %v. Genetic %v\n", cc.name, cc.upp, cc.gp)
	fmt.Printf("Homeworld: %v\n", cc.homeworld)
	fmt.Printf("SKILS %v\n", cc.skills)
	fmt.Printf("CAREER\n")
	fmt.Printf("Age %v. Born %v\n", cc.age, cc.birthdate)
	fmt.Printf("REWARDS\n")
	fmt.Printf("Current Date: %v\n", cc.curDate)
}
