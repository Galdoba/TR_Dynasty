package entity

import (
	"errors"
	"fmt"
)

const (
	CharCodeEntity            = 0
	CharCodeType              = 1
	CharCodePosition          = 2
	CharCodeShortName         = 3
	CharCodeFullName          = 4
	CharCodeTrvC1STRENGTH     = "1=1=C1=STR=Strength"
	CharCodeTrvC2AGILITY      = "1=1=C2=AGI=Agility"
	CharCodeTrvC2DEXTERITY    = "1=1=C2=DEX=Dexterity"
	CharCodeTrvC2Grace        = "1=1=C2=GRA=Grace"
	CharCodeTrvC3ENDURANCE    = "1=1=C3=END=Endurance"
	CharCodeTrvC3STAMINA      = "1=1=C3=STA=Stamina"
	CharCodeTrvC3VIGOR        = "1=1=C3=STA=Vigor"
	CharCodeTrvC4INTELLIGENCE = "1=1=C4=INT=Intelligence"
	CharCodeTrvC5EDUCATION    = "1=1=C5=EDU=Education"
	CharCodeTrvC5TRAINING     = "1=1=C5=TRA=Training"
	CharCodeTrvC5INSTINCT     = "1=1=C5=INS=Instinct"
	CharCodeTrvC6SOCIAL       = "1=1=C6=SOC=Social"
	CharCodeTrvC6CHARISMA     = "1=1=C6=CHA=Charisma"
	CharCodeTrvC6CASTE        = "1=1=C6=CAS=Caste"
	CharCodeTrvC6NOBILITY     = "1=1=C6=NOB=Nobility"
	CharCodeTrvC6TERRITORY    = "1=1=C6=TER=Territory"
	CharCodeTrvCSSANITY       = "1=1=CS=SAN=Sanity"
)

//skill -
type characteristic struct {
	entity      string //Ненужно?
	cType       string //Ненужно?
	description string //Ненужно?
	value       int
}

func newCharacteristic(charCode string) characteristic {
	chr := characteristic{}
	chr.entity = GetFromCode(CharCodeEntity, charCode)
	chr.cType = GetFromCode(CharCodeType, charCode)
	chr.description = GetFromCode(CharCodeFullName, charCode)
	return chr
}

func (chr *characteristic) setValue(newVal int) {
	chr.value = newVal
}

type Characteristic interface {
	Parameter
}

//SkillMap - объект на экспорт именно с ним должны работать внешние библиотеки
//носитель для интерфейса Skill
type CharacteristicMap struct {
	chr map[string]characteristic
}

func NewCharacteristicMap() *CharacteristicMap {
	cm := CharacteristicMap{}
	cm.chr = make(map[string]characteristic)
	return &cm
}

/*Set(string, int)
GetValue(string) (int, error)
Train(string)
Remove(string)*/

//Set - Устанавливает значение х-ки равное val
func (cm *CharacteristicMap) Set(charCode string, val int) {
	if _, ok := cm.chr[charCode]; !ok { //Если такого х-ки нет - создаем всю группу со значением 0
		cm.chr[charCode] = newCharacteristic(charCode)
	}
	skl := newCharacteristic(charCode)
	skl.setValue(val)
	cm.chr[charCode] = skl
}

//GetValue - возвращает значение Х-ки
func (cm *CharacteristicMap) GetValue(code string) (int, error) {
	if val, ok := cm.chr[code]; ok {
		return val.value, nil
	}
	return 0, errors.New("No Value for '" + code + "'")
}

//Train - увеличивает значение х-ки на 1
func (cm *CharacteristicMap) Train(code string) {
	if _, ok := cm.chr[code]; !ok {
		cm.Set(code, 0)
	} //мы точно знаем что х-ка есть
	val, err := cm.GetValue(code)
	if err != nil {
		fmt.Println(err)
	}
	cm.Set(code, val+1)
}

//Remove - Удаляет запись о навыке - в теории эта функция вообще не должна использоваться
func (cm *CharacteristicMap) Remove(code string) {
	delete(cm.chr, code)
}

func CharacteristicsCodesList() []string {
	return []string{
		CharCodeTrvC1STRENGTH,
		CharCodeTrvC2AGILITY,
		CharCodeTrvC2DEXTERITY,
		CharCodeTrvC2Grace,
		CharCodeTrvC3ENDURANCE,
		CharCodeTrvC3STAMINA,
		CharCodeTrvC3VIGOR,
		CharCodeTrvC4INTELLIGENCE,
		CharCodeTrvC5EDUCATION,
		CharCodeTrvC5TRAINING,
		CharCodeTrvC5INSTINCT,
		CharCodeTrvC6SOCIAL,
		CharCodeTrvC6CHARISMA,
		CharCodeTrvC6CASTE,
		CharCodeTrvC6NOBILITY,
		CharCodeTrvC6TERRITORY,
		CharCodeTrvCSSANITY,
	}
}
