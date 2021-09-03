package traveller

import (
	"github.com/Galdoba/TR_Dynasty/T5/assets"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

func (trv *TravellerT5) GenerateHomeworld() {
	switch trv.randomHomeWorld {
	case true:
		roll := dice.RollD66()
		trv.homeworld = setHW(roll)
	case false:
		homeworldData, _ := otu.GetDataOn("Regina")
		homeworld, _ := wrld.FromOTUdata(homeworldData)
		trv.homeworld = &homeworld
	}
	trv.getBackgroundSkills()
}

func setHW(code string) *wrld.World {
	worldName := callWorld(code)
	homeworldData, err := otu.GetDataOnTest(worldName)
	if err != nil {
		panic(code + " invalid")
	}
	homeworld, err2 := wrld.FromOTUdata(homeworldData)
	if err2 != nil {
		panic(code + " invalid 2")
	}
	return &homeworld
}

func callWorld(code string) string {
	hwNames := make(map[string]string)
	hwNames["11"] = "Alell"
	hwNames["12"] = "Boughene"
	hwNames["13"] = "Capital"
	hwNames["14"] = "Dorannia" //c
	hwNames["15"] = "Efate"
	hwNames["16"] = "Feri"
	hwNames["21"] = "Magash"
	hwNames["22"] = "Hefry"
	hwNames["23"] = "Jenghe"
	hwNames["24"] = "Terra"
	hwNames["25"] = "Lakou"
	hwNames["26"] = "Sperle"
	hwNames["31"] = "Knorbes"
	hwNames["32"] = "Preslin"
	hwNames["33"] = "Yori"
	hwNames["34"] = "Regina"
	hwNames["35"] = "Regina"
	hwNames["36"] = "Regina"
	hwNames["41"] = "Ruie"
	hwNames["42"] = "Tremous Dex"
	hwNames["43"] = "Uakye"
	hwNames["44"] = "Vland"
	hwNames["45"] = "Wroclaw"
	hwNames["46"] = "Menorb" //c
	hwNames["51"] = "Yorbund"
	hwNames["52"] = "Traltha"
	hwNames["53"] = "Dentus"
	hwNames["54"] = "Vanzeti"
	hwNames["55"] = "Syr Darya"
	hwNames["56"] = "Noricum"
	hwNames["61"] = "Rhylanor"
	hwNames["62"] = "Raschev"
	hwNames["63"] = "Ara Pacis"
	hwNames["64"] = "Roup"
	hwNames["65"] = "Pax Rulin"
	hwNames["66"] = "Drinax"
	return hwNames[code]
}

func (trv *TravellerT5) getBackgroundSkills() {
	for _, tc := range trv.homeworld.TradeClassificationsSl() {
		skill := 0
		switch tc {
		default:
			continue
		case "Ag":
			skill = assets.SKILL_Animals
		case "As":
			skill = assets.SKILL_ZeroG
		case "Co":
			skill = assets.SKILL_HostileEnviron
		case "Cp":
			skill = assets.SKILL_Admin
		case "Cs":
			skill = assets.SKILL_Bureaucrat
		case "Cx":
			skill = assets.SKILL_Language
		case "Da":
			skill = assets.SOLDER_Fighter
		case "De":
			skill = assets.SKILL_Survival
		case "Ds":
			skill = assets.SKILL_VaccSuit
			skill = assets.SKILL_ZeroG
		case "Fa":
			skill = assets.SKILL_Animals
		case "Fl":
			skill = assets.SKILL_HostileEnviron
		case "Fr":
			skill = assets.SKILL_HostileEnviron
		case "Ga":
			skill = assets.SKILL_Trader
		case "He":
			skill = assets.SKILL_HostileEnviron
		case "Hi":
			skill = assets.SKILL_Streetwise
		case "Ho":
			skill = assets.SKILL_HostileEnviron
		case "Ic":
			skill = assets.SKILL_VaccSuit
		case "In":
			trade := dice.New().RollFromListInt([]int{assets.TRADE_Biologics, assets.TRADE_Craftsman, assets.TRADE_Electronics, assets.TRADE_Fluidics, assets.TRADE_Gravitics, assets.TRADE_Magnetics, assets.TRADE_Mechanic, assets.TRADE_Photonics, assets.TRADE_Polymers, assets.TRADE_Programmer})
			skill = trade
		case "Lo":
			skill = assets.SKILL_Flyer
		case "Mi":
			skill = assets.SKILL_Survey
		case "Na":
			skill = assets.SKILL_Survey
		case "Ni":
			skill = assets.SKILL_Driver
		case "Oc":
			skill = assets.SKILL_HighG
		case "Pa":
			skill = assets.SKILL_Trader
		case "Pi":
			skill = assets.SKILL_JOT
		case "Po":
			skill = assets.SHIP_Steward
		case "Pr":
			skill = assets.TRADE_Craftsman
		case "Ri":
			art := dice.New().RollFromListInt([]int{assets.ART_Actor, assets.ART_Artist, assets.ART_Author, assets.ART_Chef, assets.ART_Dancer, assets.ART_Musician})
			skill = art
		case "Tr":
			skill = assets.SKILL_Survival
		case "Tu":
			skill = assets.SKILL_Survival
		case "Tz":
			skill = assets.SKILL_Driver
		case "Va":
			skill = assets.SKILL_VaccSuit
		case "Wa":
			skill = assets.SKILL_Seafarer
		}
		trv.skills[skill] = assets.NewSkill(skill)
		trv.skills[skill].Train()
	}
}
