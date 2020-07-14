package main

import (
	"strings"

	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"
)

const (
	difficultyVeryEasy       = 4
	difficultyEasy           = 2
	difficultyAverage        = 0
	difficultyHard           = -2
	difficultyVeryHard       = -4
	difficultyImpossible     = -6
	archConglomerate         = "Conglomerate"
	archMediaEmpire          = "Media Empire"
	archMerchantMarket       = "Merchant Market"
	archMilitaryCharter      = "Military Charter"
	archNobleLine            = "Noble Line"
	archReligiousFaith       = "Religious Faith"
	archSyndicate            = "Syndicate"
	charDEFAULT              = "charDefault"
	charNONE                 = "charNONE"
	charCleverness           = "Cleverness"
	charGreed                = "Greed"
	charLoyalty              = "Loyalty"
	charMilitarism           = "Militarism"
	charPopularity           = "Popularity"
	charScheming             = "Scheming"
	charTenacity             = "Tenacity"
	charTradition            = "Tradition"
	traitCulture             = "Culture"
	traitFiscalDefence       = "FiscalDefence"
	traitFleet               = "Fleet"
	traitTechnology          = "Technology"
	traitTerritorialDefence  = "Territorial Defence"
	apttNONE                 = "NONE"
	apttAcquisition          = "Acquisition"
	apttBureaucracy          = "Bureaucracy"
	apttConquest             = "Conquest"
	apttEconomics            = "Economics"
	apttEntertain            = "Entertain"
	apttExpression           = "Expression"
	apttHostility            = "Hostility"
	apttIllicit              = "Illicit"
	apttIntel                = "Intel"
	apttMaintenance          = "Maintenance"
	apttPolitics             = "Politics"
	apttPosturing            = "Posturing"
	apttPropaganda           = "Propaganda"
	apttPublicRelations      = "Public Relations"
	apttRecruitment          = "Recruitment"
	apttResearch             = "Research"
	apttSabotage             = "Sabotage"
	apttSecurity             = "Security"
	apttTactical             = "Tactical"
	apttTutelage             = "Tutelage"
	valueMorale              = "Morale"
	valuePopulace            = "Populace"
	valueWealth              = "Wealth"
	mAssetBoardOfDirectors   = "Board of Directors"
	mAssetCommandStaff       = "Command Staff"
	mAssetHeroicLeaders      = "Heroic Leaders"
	mAssetMatriarchPatriarch = "Matriarch/Patriarch"
	mAssetOverlord           = "Overlord"
	mAssetTheocrat           = "Theocrat"
)

//2DD
//2D
//2D+3
//D66
func trvRoll(code string) int {
	codeParts := strings.Split(code, "D")
	if len(codeParts) != 2 {
		//special case: Roll Destructive
	}
	diceQ := codeParts[0]
	resultMod := convert.StoI(codeParts[1])
	result := utils.RollDice(diceQ+"d6", resultMod)
	return result
}

type npc struct {
	str int
	dex int
	end int
	age int
	bil int
}

func seeAge() {
	npc := &npc{}
	npc.str = trvRoll("2D")
	npc.dex = trvRoll("2D")
	npc.end = trvRoll("2D")
	npc.age = 18
	npc.bil = 0
	for npc.alive() {
		npc.age = npc.age + 4
		
	}
}

func (npc *npc) alive() bool {
	if npc.str > 0 && npc.dex > 0 && npc.end > 0 {
		return true
	}
	return false
}
