package npcmakerv2

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/TR_Dynasty/entity"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

func Test() {
	fmt.Println("Start test")
	data, err := otu.GetDataOn("Paal")
	if err != nil {
		fmt.Println(err)
	}
	w, _ := wrld.FromOTUdata(data)

	trv := newTraveller(w)
	trv.skills.Train(entity.SCTrvEngineer)
	trv.skills.Train(entity.SCTrvEngineerLifesupport)
	trv.skills.Train(entity.SCTrvEngineerLifesupport)
	trv.skills.Train(entity.SCTrvEngineerLifesupport)
	trv.skills.Train(entity.SCTrvEngineer)
	trv.skills.Train(entity.SCTrvAdmin)
	trv.skills.Train(entity.SCTrvAdmin)
	trv.skills.Train(entity.SCTrvDriveWheel)
	fmt.Println(showSkill(trv.skills))
	//fmt.Println(trv.skills)
	fmt.Println(entity.GetFromCode(3, entity.CharCodeTrvC3ENDURANCE))
	fmt.Println("End test")
}

type traveller struct {
	originWorld wrld.World
	//characteristics entity.Characteristic
	skills entity.Skill
}

func newTraveller(originWorld wrld.World) traveller {
	trv := traveller{}
	//trv.attributes = entity.NewCharactiristicMap()
	trv.skills = entity.NewSkillMap()
	trv.originWorld = originWorld //TODO: сделать мир опциональным
	trv.TrainBackgroundSkills()
	return trv
}

func showSkill(skmap entity.Skill) string {
	skills := ""
	//ТУДУ: попробывать удалить все нулевые скилы если специализация не равна 0
	for _, skillCode := range entity.SkillCodesList() {
		propose := skmap.Get(skillCode).String()
		if propose != "" {
			skills += skmap.Get(skillCode).String() + ", "
		}
	}
	skills = strings.TrimSuffix(skills, ", ")
	return skills
}

func (trv *traveller) TrainBackgroundSkills() {
	// data, err := otu.GetDataOn(trv.originWorld)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	backgrounsTC := []string{}
	//if trv.originWorld != nil {
	backgrounsTC = trv.originWorld.TradeCodes()
	//}
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

/*
НПС это сущность
Интерфейсы:
-делать проверки навыка (SkillTester)
-делать проверки характеристики (AttributeTester)
-рассказывать о своих навыках (SkillGiver)
-рассказывать о своих характеристиках (AtribbuteGiver)
-рассказывать о себе (Describer)

*/
