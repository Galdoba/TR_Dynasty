package traveller

import (
	"errors"

	"github.com/Galdoba/TR_Dynasty/pkg/core/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

type traveller struct {
	dice  *dice.Dicepool //дайспул для бросков кубиков
	name  string
	race  string //структ с данными по рассе
	chars map[string]ehex.DataRetriver
	//skills map[string]Skill
	education education
	careers   []career
}

type education struct {
}

type career struct {
	name             string              //название карьеры
	specs            []string            //имена специализаций
	rank             map[string]int      //полученные ранги в рамках специализаций
	rollTNCode       map[string]string   //тест необходимый для прохождения персонажем TODO: это должна быть структура rollTN со стрингов внутри
	developmentTable map[string][]string //таблицы обучения
	terms            int                 //проведенные в карьере термы
	aquiredBenefits  []string            //полученные бонусы
	eventList        []string            //произошедшие события в термах - возможно нужен отдельный структ
}

type Traveller interface {
	Name() string
}

func NewCharacter() (*traveller, error) {
	err := errors.New("Initial")
	chr := &traveller{}
	chr.dice = dice.New()
	chr.chars = make(map[string]ehex.DataRetriver)
	if err.Error() == "Initial" {
		return chr, nil
	}
	return chr, err
}

//SetChar -
func (trv *traveller) SetChar(char string, val int) {
	trv.chars[char] = ehex.New(val)
}

func (trv *traveller) rollCharacteristics() {
	for _, val := range []string{"STR", "DEX", "END", "INT", "EDU", "SOC"} {
		trv.SetChar(val, trv.dice.RollNext("2d6").Sum())
	}
}

//Name - return Traveller name
func (trv *traveller) Name() string {
	return trv.name
}
