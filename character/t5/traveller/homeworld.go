package traveller

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/T5/assets"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

func (trv *TravellerT5) GenerateHomeworld() {
	switch trv.randomHomeWorld {
	case true:
		roll := dice.RollD66()
		fmt.Println(roll)
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
	homeworldData, err := otu.GetDataOn(worldName)
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
	hwNames["24"] = "Earth"
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
		switch tc {
		case "Ab":

		case "Ag":
			trv.skills[assets.SKILL_Animals] = *assets.NewSkill(assets.SKILL_Animals)
		case "An":

		case "As":
			//trv.GetSkill//Asteroid Zero-G
			trv.skills[assets.SKILL_ZeroG] = *assets.NewSkill(assets.SKILL_ZeroG)
		case "Ba":

		case "Co":
			//trv.GetSkill//Cold Hostile Env
			trv.skills[assets.SKILL_HostileEnviron] = *assets.NewSkill(assets.SKILL_HostileEnviron)
		case "Cp":
			//trv.GetSkill//Subsector Capital Admin
			trv.skills[assets.SKILL_Admin] = *assets.NewSkill(assets.SKILL_Admin)
		case "Cs":
			//trv.GetSkill//Sector Capital Bureaucrat
			trv.skills[assets.SKILL_Bureaucrat] = *assets.NewSkill(assets.SKILL_Bureaucrat)
		case "Cx":
			//trv.GetSkill//Capital Language
			trv.skills[assets.SKILL_Language] = *assets.NewSkill(assets.SKILL_Language)
		case "Da":
			//trv.GetSkill//Dangerous Fighter
			trv.skills[assets.SOLDER_Fighter] = *assets.NewSkill(assets.SOLDER_Fighter)
		case "De":
			//trv.GetSkill//Desert Survival
			trv.skills[assets.SKILL_Survival] = *assets.NewSkill(assets.SKILL_Survival)
		case "Di":
			//trv.GetSkill//Die Back (no skill)
		case "Ds":
			//trv.GetSkill//Deep Space Vacc Suit +Zero-G
			trv.skills[assets.SKILL_VaccSuit] = *assets.NewSkill(assets.SKILL_VaccSuit)
			trv.skills[assets.SKILL_ZeroG] = *assets.NewSkill(assets.SKILL_ZeroG)
		case "Fa":
			//trv.GetSkill//Farming Animals
			trv.skills[assets.SKILL_Animals] = *assets.NewSkill(assets.SKILL_Animals)
		case "Fl":
			//trv.GetSkill//Fluid Hostile Env
			trv.skills[assets.SKILL_HostileEnviron] = *assets.NewSkill(assets.SKILL_HostileEnviron)
		case "Fo":
			//trv.GetSkill//Forbidden (no skill)
		case "Fr":
			//trv.GetSkill//Frozen Hostile Env
			trv.skills[assets.SKILL_HostileEnviron] = *assets.NewSkill(assets.SKILL_HostileEnviron)
		case "Ga":
			//trv.GetSkill//Garden World Trader
			trv.skills[assets.SKILL_Trader] = *assets.NewSkill(assets.SKILL_Trader)
		case "He":
			//trv.GetSkill//Hellworld Hostile Env
			trv.skills[assets.SKILL_HostileEnviron] = *assets.NewSkill(assets.SKILL_HostileEnviron)
		case "Hi":
			//trv.GetSkill//High Population Streetwise
			trv.skills[assets.SKILL_Streetwise] = *assets.NewSkill(assets.SKILL_Streetwise)
		case "Ho":
			//trv.GetSkill//Hot Hostile Env
			trv.skills[assets.SKILL_HostileEnviron] = *assets.NewSkill(assets.SKILL_HostileEnviron)
		case "Ic":
			//trv.GetSkill//Ice-Capped Vacc Suit
			trv.skills[assets.SKILL_VaccSuit] = *assets.NewSkill(assets.SKILL_VaccSuit)
		case "In":
			//trv.GetSkill//Industrial One Trade
		case "Lk":
			//trv.GetSkill//Locked (no skill)
		case "Lo":
			//trv.GetSkill//Low Population Flyer
		case "Mi":
			//trv.GetSkill//Mining Survey
		case "Mr":
			//trv.GetSkill//Military Rule (no skill)
		case "Na":
			//trv.GetSkill//Non-agricultural Survey
		case "Ni":
			//trv.GetSkill//Non-industrial Driver
		case "Oc":
			//trv.GetSkill//Ocean World Hi-G
		case "Pa":
			//trv.GetSkill//Pre-Agricultural Trader
		case "Ph":
			//trv.GetSkill//Pre-High Population (no skill)
		case "Pi":
			//trv.GetSkill//Pre-Industrial JOT
		case "Po":
			//trv.GetSkill//Poor Steward
		case "Pr":
			//trv.GetSkill//Pre-Rich Craftsman
		case "Px":
			//trv.GetSkill//Prison Exile Camp (no skill)
		case "Pz":
			//trv.GetSkill//Puzzling (no skill)
		case "Re":
			//trv.GetSkill//Reserve (no skill)
		case "Ri":
			//trv.GetSkill//Rich One Art
		case "Tr":
			//trv.GetSkill//Tropic Survival
		case "Tu":
			//trv.GetSkill//Tundra Survival
		case "Tz":
			//trv.GetSkill//Twilight Zone Driver
		case "Va":
			//trv.GetSkill//Vacuum Vacc Suit
		case "Wa":
			//trv.GetSkill//Water World Seafarer
		}
	}
}
