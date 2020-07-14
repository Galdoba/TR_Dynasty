package main

import "github.com/Galdoba/utils"

func successOf2d6(dm int, tn int) float64 {
	var validRolls int
	for d1 := 1; d1 < 7; d1++ {
		for d2 := 1; d2 < 7; d2++ {
			if d1+d2+dm >= tn {
				validRolls++
			}
		}
	}
	flChance := utils.RoundFloat64(float64(validRolls)/36.0, 3)
	return flChance
}

func (dyn *Dynasty) aptitideCheck(apt string, char string, gmMod int) (effect int) {
	aptMod := dyn.pickVal(apt)
	if aptMod < 0 {
		aptMod = -2
	}
	if apt == apttNONE {
		aptMod = 0
	}
	charMod := charDM(dyn.pickVal(char))
	r := utils.RollDice("2d6", aptMod, charMod, gmMod)
	effect = r - 8
	if effect > 6 {
		effect = 6
	}
	if effect < -6 {
		effect = -6
	}
	if effect > 0 {
		dyn.aptTotalEffect[apt] = dyn.aptTotalEffect[apt] + effect
	}
	return effect
}

func (dyn *Dynasty) probeAptitideCheck(apt string, char string, gmMod int) float64 {
	aptMod := dyn.pickVal(apt)
	if aptMod < 0 {
		aptMod = -2
	}
	if apt == apttNONE {
		aptMod = 0
	}
	if char == charDEFAULT {
		panic("да это ветка нужна")
		//char = dyn.bestDefaultCharOf(apt)
	}
	charMod := charDM(dyn.pickVal(char))
	return successOf2d6(aptMod+charMod+gmMod, 9) //Дает вероятность позитивного эффекта
}

func (dyn *Dynasty) failureCheck(atr string, tn int) bool {
	if isCharacteristic(atr) {
		charMod := charDM(dyn.pickVal(atr))
		r := utils.RollDice("2d6", charMod, 0)
		effect := r - tn
		if effect > -1 {
			return true
		}
		return false
	}
	charMod := dyn.pickVal(atr)
	if charMod < 0 {
		charMod = -2
	}
	r := utils.RollDice("2d6", charMod, 0)
	effect := r - tn
	if effect > -1 {
		return true
	}
	return false
}

func (dyn *Dynasty) probeFailureCheck(atr string, tn int) float64 {
	atrMod := 0
	if isCharacteristic(atr) {
		atrMod = charDM(dyn.pickVal(atr))
	} else {
		atrMod = dyn.pickVal(atr)
		if atrMod < 0 {
			atrMod = -2
		}
	}
	return successOf2d6(atrMod, tn) //Дает вероятность позитивного изхода
}

// func (dyn *Dynasty) bestDefaultCharOf(apt string) string {
// 	possibleOption := posssibleOptionsOfAptitude(apt)
// 	posibilityesMap := make(map[string]int)
// 	for i := range possibleOption {
// 		posibilityesMap[possibleOption[i]] = charDM(dyn.characteristics[possibleOption[i]])
// 	}
// 	bestOption := "none Chosen"
// 	maxDM := -999
// 	for key, val := range posibilityesMap {
// 		if val > maxDM {
// 			bestOption = key
// 		}
// 	}
// 	return bestOption
// }

// func posssibleOptionsOfAptitude(apt string) []string {
// 	var possibleOption []string
// 	switch apt {
// 	case apttBureaucracy:
// 		possibleOption = append(possibleOption, charGreed, charPopularity, charLoyalty)
// 	case apttConquest:
// 		possibleOption = append(possibleOption, charGreed, charMilitarism, charLoyalty)
// 	case apttEconomics:
// 		possibleOption = append(possibleOption, charGreed, charCleverness, charScheming)
// 	case apttEntertain:
// 		possibleOption = append(possibleOption, charPopularity, charLoyalty)
// 	case apttExpression:
// 		possibleOption = append(possibleOption, charScheming, charPopularity)
// 	case apttHostility:
// 		possibleOption = append(possibleOption, charMilitarism, charScheming)
// 	case apttIllicit:
// 		possibleOption = append(possibleOption, charLoyalty, charScheming)
// 	case apttIntel:
// 		possibleOption = append(possibleOption, charCleverness, charScheming)
// 	case apttMaintenance:
// 		possibleOption = append(possibleOption, charLoyalty, charTenacity, charGreed)
// 	case apttPolitics:
// 		possibleOption = append(possibleOption, charCleverness, charPopularity, charScheming)
// 	case apttPosturing:
// 		possibleOption = append(possibleOption, charLoyalty, charTradition)
// 	case apttPropaganda:
// 		possibleOption = append(possibleOption, charScheming, charPopularity)
// 	case apttPublicRelations:
// 		possibleOption = append(possibleOption, charCleverness, charPopularity, charTradition)
// 	case apttRecruitment:
// 		possibleOption = append(possibleOption, charPopularity, charTradition, charGreed)
// 	case apttResearch:
// 		possibleOption = append(possibleOption, charScheming, charCleverness)
// 	case apttSabotage:
// 		possibleOption = append(possibleOption, charScheming, charMilitarism)
// 	case apttSecurity:
// 		possibleOption = append(possibleOption, charTenacity, charMilitarism, charLoyalty)
// 	case apttTactical:
// 		possibleOption = append(possibleOption, charPopularity, charMilitarism)
// 	case apttTutelage:
// 		possibleOption = append(possibleOption, charTradition, charCleverness, charTenacity)
// 	}
// 	return possibleOption
// }
