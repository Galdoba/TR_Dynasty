package npcmakerv2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/entity"
	"github.com/Galdoba/TR_Dynasty/name"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/wrld"
	"github.com/Galdoba/utils"
)

func Test() {
	fmt.Println("Start test")
	data, err := otu.GetDataOn("Staha")
	if err != nil {
		fmt.Println(err)
	}
	w, _ := wrld.FromOTUdata(data)
	utils.RandomSeed()
	trv := NewTraveller(w, utils.RandomFromList(ListCareers()))

	//trv.skills.Set(entity.SCTrvDriveWheel, 3)

	fmt.Println("--------------")

	fmt.Println(&trv)
	fmt.Println("--------------")
	fmt.Println("End test")
}

type Traveller struct {
	name            string
	originWorld     wrld.World
	characteristics entity.Characteristic
	skills          entity.Skill
	counters        map[string]int
	career          string
	rank            int
	term            int
}

func NewTraveller(originWorld wrld.World, career string) Traveller {
	trv := Traveller{}
	trv.name = name.RandomNew()
	trv.career = career
	trv.skills = entity.NewSkillMap()
	trv.characteristics = entity.NewCharacteristicMap()
	trv.originWorld = originWorld //TODO: сделать мир опциональным
	trv.counters = make(map[string]int)
	trv.term = 2 + TrvCore.FluxGood()
	trv.rank = 0
	for _, val := range RaceChars() {
		trv.characteristics.Set(val, dice.Roll("2d6").Sum())
	}
	trv.trainBackgroundSkills()
	trv.getCareerCodes()
	//trv.getBenefits()

	//fmt.Println("===")
	return trv
}

type Condition interface {
	Apply(string, string)
}

type condition struct {
	table, index string
}

func NewTraveller2() Traveller {
	trv := Traveller{}
	return trv
}

func (trv *Traveller) String() string {
	str := ""
	str += "   Name: " + trv.name
	str += "\n Origin: " + trv.originWorld.Name() + " [" + trv.originWorld.UWP() + "]"
	str += "\n    UPP: " + trv.PrintUPP()
	str += "\n    Age: " + strconv.Itoa(trv.term*4+18)
	str += "\n Career: " + entity.GetFromCode(4, trv.career)
	str += "\n   Rank: " + strconv.Itoa(trv.rank) + "/" + strconv.Itoa(trv.term)
	str += "\n Skills: " + trv.PrintSkills()
	return str
}

func (trv Traveller) PrintSkills() string {
	skills := ""
	for _, skillCode := range entity.SkillCodesList() {
		if entity.GetFromCode(entity.SCSpeciality, skillCode) != "0" {
			continue
		}
		propose := trv.skills.FPrintSkillGroupS(skillCode)
		if propose != "" {
			skills += propose + ", "
		}
	}
	skills = strings.TrimSuffix(skills, ", ")
	return skills
}

func (trv Traveller) PrintUPP() string {
	upp := ""
	for _, chrc := range RaceChars() {
		val, err := trv.characteristics.GetValue(chrc)
		if err != nil {
			trv.characteristics.Set(chrc, 2)
		}
		upp += TrvCore.DigitToEhex(val)
	}
	return upp
}

func (trv *Traveller) trainBackgroundSkills() {
	backgrounsTC := []string{}
	backgrounsTC = trv.originWorld.TradeCodes()
	for _, tc := range backgrounsTC {
		switch tc {
		case constant.TradeCodeAgricultural:
			trv.skills.Train(entity.SCTrvAnimalsHandling)
		case constant.TradeCodeAsteroid:
			trv.skills.Train(entity.SCTrvAthleticsDEX)
			trv.skills.Train(entity.SCTrvVaccsuit)
		case constant.TradeCodeCold:
			trv.skills.Train(entity.SCTrvAthleticsEND)
			trv.skills.Train(entity.SCTrvVaccsuit)
		case constant.TradeCodeSubsectorCapital:
			trv.skills.Train(entity.SCTrvAdmin)
		case constant.TradeCodeSectorCapital:
			trv.skills.Train(entity.SCTrvAdmin)
			trv.skills.Train(entity.SCTrvAdvocate)
		case constant.TradeCodeCapital:
			trv.skills.Train(entity.SCTrvLanguageAnglic)
		case constant.TradeCodeDangerous:
			trv.skills.Train(entity.SCTrvMeleeBlade)
		case constant.TradeCodeDesert:
			trv.skills.Train(entity.SCTrvSurvival)
		case constant.TradeCodeDeepSpace:
			trv.skills.Train(entity.SCTrvAthleticsDEX)
			trv.skills.Train(entity.SCTrvVaccsuit)
		case constant.TradeCodeFarming:
			trv.skills.Train(entity.SCTrvAnimalsHandling)
		case constant.TradeCodeFluidOceans:
			trv.skills.Train(entity.SCTrvAthleticsEND)
			trv.skills.Train(entity.SCTrvVaccsuit)
		case constant.TradeCodeFrozen:
			trv.skills.Train(entity.SCTrvSurvival)
			trv.skills.Train(entity.SCTrvVaccsuit)
		case constant.TradeCodeGarden:
			trv.skills.Train(entity.SCTrvBroker)
		case constant.TradeCodeHellworld:
			trv.skills.Train(entity.SCTrvSurvival)
			trv.skills.Train(entity.SCTrvVaccsuit)
		case constant.TradeCodeHighPopulation:
			trv.skills.Train(entity.SCTrvStreetwise)
		case constant.TradeCodeHot:
			trv.skills.Train(entity.SCTrvSurvival)
			trv.skills.Train(entity.SCTrvVaccsuit)
		case constant.TradeCodeIceCapped:
			trv.skills.Train(entity.SCTrvVaccsuit)
		case constant.TradeCodeIndustrial:
			trv.skills.Train(entity.SCTrvMechanic)
		case constant.TradeCodeLowPopulation:
			trv.skills.Train(entity.SCTrvFlyerGrav)
		case constant.TradeCodeNonAgricultural:
			trv.skills.Train(entity.SCTrvInvestigate)
		case constant.TradeCodeNonIndustrial:
			trv.skills.Train(entity.SCTrvDriveWheel)
		case constant.TradeCodePreAgricultural:
			trv.skills.Train(entity.SCTrvBroker)
		case constant.TradeCodePreIndustrial:
			trv.skills.Train(entity.SCTrvJackofalltrades)
		case constant.TradeCodePoor:
			trv.skills.Train(entity.SCTrvSteward)
		case constant.TradeCodePreRich:
			trv.skills.Train(entity.SCTrvProfessionAny)
		case constant.TradeCodeRich:
			trv.skills.Train(entity.SCTrvArtVisualMedia)
		case constant.TradeCodeTropic:
			trv.skills.Train(entity.SCTrvSurvival)
		case constant.TradeCodeTundra:
			trv.skills.Train(entity.SCTrvSurvival)
		case constant.TradeCodeVacuum:
			trv.skills.Train(entity.SCTrvVaccsuit)
		case constant.TradeCodeWaterWorld:
			trv.skills.Train(entity.SCTrvSeafarerPersonal)
		case constant.TradeCodeLowTech:
			trv.skills.Train(entity.SCTrvSurvival)
		case constant.TradeCodeHighTech:
			trv.skills.Train(entity.SCTrvElectronicsComputers)
		default:
		}
	}
}

func RaceChars() []string {
	return []string{
		entity.CharCodeTrvC1STRENGTH,
		entity.CharCodeTrvC2DEXTERITY,
		entity.CharCodeTrvC3ENDURANCE,
		entity.CharCodeTrvC4INTELLIGENCE,
		entity.CharCodeTrvC5EDUCATION,
		entity.CharCodeTrvC6SOCIAL,
	}
}

/*
НПС это сущность
Интерфейсы:
-делать проверки навыка (SkillTester)
-делать проверки характеристики (AttributeTester)
-рассказывать о своих навыках (SkillGiver)
-рассказывать о своих характеристиках (AtribbuteGiver)
-рассказывать о себе (Describer)

*/
