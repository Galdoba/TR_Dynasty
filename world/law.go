package world

// import (
// 	"errors"
// 	"fmt"
// 	"strings"

// 	"github.com/Galdoba/TR_Dynasty/TrvCore"
// )

// //Codex - Описывает законы планеты
// type Codex struct {
// 	overall        int
// 	lawProfile     string
// 	branchesByName map[string]int
// }

// /*
// overall

// Weapons
// Trade
// Criminal Law
// Civil Law
// Personal Freedom

// Weapons
// Drugs
// Information
// Technology
// Travellers
// Psionics

// */

// func lawBranches() []string {
// 	return []string{
// 		"Trade",
// 		"Criminal Law",
// 		"Civil Law",
// 		"Personal Freedom",
// 		"Weapons",
// 		"Drugs",
// 		"Information",
// 		"Technology",
// 		"Travellers",
// 		"Psionics",
// 	}
// }

// //CodexFromULP -
// func CodexFromULP(ulp string) (*Codex, error) {
// 	if len(ulp) != 13 {
// 		return nil, errors.New("ulp Leight incorect: '" + ulp + "'")
// 	}
// 	ulpData := strings.Split(ulp, "")
// 	lawStruct := &Codex{}
// 	lawStruct.lawProfile = ulp
// 	lawStruct.branchesByName = make(map[string]int)
// 	branhces := lawBranches()
// 	brChecked := 0
// 	for i, val := range ulpData {
// 		switch i {
// 		case 1, 6:
// 			if val != "-" {
// 				return nil, errors.New("ulp Format unreadable: '" + ulp + "'")
// 			}
// 		case 0:
// 			lawStruct.overall = TrvCore.EhexToDigit(val)
// 		default:
// 			lawStruct.branchesByName[branhces[brChecked]] = TrvCore.EhexToDigit(val)
// 			brChecked++
// 		}
// 	}
// 	return lawStruct, nil
// }

// //GovermentProfile: 3-234131-2CC0 ///Grand Census Proto
// //Universal Cortporate Profile: Im-1415-15-PubS-Util7-Ret2-Svc1-0266-13 ///Suplement 15

// // func uwpCodexConverter(uwp string) (statGovr int, statLaws int) {
// // 	uwpData := strings.Split(uwp, "")
// // 	statLaws = TrvCore.EhexToDigit(uwpData[6])
// // 	statGovr = TrvCore.EhexToDigit(uwpData[5])
// // 	return statGovr, statLaws
// // }

// //NewCodex - Создает новый объект для описания законов на основе UWP Мира
// func NewCodex(world *World) (*Codex, error) {
// 	lawStruct := &Codex{}
// 	//statGovr, statLaws := uwpCodexConverter(worldUWP)
// 	if val, ok := world.data["ULP"]; ok {
// 		return CodexFromULP(val)
// 	}
// 	worldStats := world.Stats()
// 	statGovr := worldStats["Govr"]
// 	statLaws := worldStats["Laws"]
// 	lawStruct.overall = statLaws
// 	govr := statGovr
// 	lawStruct.branchesByName = make(map[string]int)
// 	//base := lawStruct.overall
// 	lawStruct.lawProfile = TrvCore.DigitToEhex(lawStruct.overall)
// 	for i, val := range lawBranches() {
// 		if i == 0 || i == 4 {
// 			lawStruct.lawProfile += "-"
// 		}
// 		lawStruct.branchesByName[val] = TrvCore.Roll2D() - 7 + govr
// 		if lawStruct.branchesByName[val] < 0 {
// 			lawStruct.branchesByName[val] = 0
// 		}
// 		lawStruct.lawProfile += TrvCore.DigitToEhex(lawStruct.branchesByName[val])
// 	}
// 	world.data["ULP"] = lawStruct.lawProfile
// 	return lawStruct, nil
// }

// func (world *World) UGP() string {
// 	if val, ok := world.data["UGP"]; ok {
// 		return val
// 	}
// 	culProf := NewCulturalProfile(world)
// 	if strings.Contains(culProf.societyProfile, "?") {
// 		fmt.Println(culProf)
// 		panic("stop")
// 	}
// 	return culProf.societyProfile

// }

// func compareStr(str string, check ...string) bool {
// 	for _, val := range check {
// 		if str == val {
// 			return true
// 		}
// 	}
// 	return false
// }
