package traveller

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/T5/assets"
)

type TravellerT5 struct {
	notHuman       bool
	characteristic map[int]assets.Characteristic
	skills         map[int]assets.Skill
	knowledges     map[int]assets.Knowledge
	homeworld      string
	err            error
}

func NewTravellerT5() *TravellerT5 {
	trv := TravellerT5{}
	trv.characteristic = make(map[int]assets.Characteristic)
	trv.err = fmt.Errorf("generation of traveller not implemented")
	trv.GenerateCharactiristics()
	fmt.Println(trv)
	return &trv
}

type gptemplate struct {
	data []int //6 характеристик и 6 привязанных к ним дайсов
}

type geneticProfile interface {
	Profile() []int
}

func (gp *gptemplate) Profile() []int {
	return []int{
		assets.Strength, assets.Dexterity, assets.Endurance, assets.Intelligence, assets.Education, assets.Social,
		2, 2, 2, 2, 2, 2,
	}
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
	name      string
	uwp       string
	gp        string
	homeworld string
	skills    []string
	career    []string
	age       int
	birthdate string
	rewards   []string
	curDate   string
}

func (trv *TravellerT5) PrintCard() {

}
