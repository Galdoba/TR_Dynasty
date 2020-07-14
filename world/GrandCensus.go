package world

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/utils"
)

const (
	tlHIGH             = "High Common Tech Level"
	tlLOW              = "Low Common Tech Level"
	tlENERGY           = "Energy Tech Level"
	tlCOMPandROBOT     = "Computers/Robotics Tech Level"
	tlCOMMUNICATION    = "Communication Tech Level"
	tlMEDICAL          = "Medical Tech Level"
	tlENVIROMENT       = "Enviroment Tech Level"
	tlTRANSPORTland    = "Land Transport Tech Level"
	tlTRANSPORTwater   = "Water Transport Tech Level"
	tlTRANSPORTair     = "Air Transport Tech Level"
	tlTRANSPORTspace   = "Space Transport Tech Level"
	tlMILITARYpersonal = "Personal Military Tech Level"
	tlMILITARYheavy    = "Heavy Military Tech Level"
)

//CulturalProfile -
type CulturalProfile struct {
	socialOutlook     map[string]string
	authorityBranches map[string]string
	societyProfile    string
	lawUniformity     string
	techProfile       string
}

func coltOutLIST() []string {
	return []string{
		"Progressiveness Attitude",
		"Progressiveness Action",
		"Aggressiveness Attitude",
		"Aggressiveness Action",
		"Extensiveness Global",
		"Extensiveness Interstellar",
	}
}

func branchesList() []string {
	return []string{
		"Legislative",
		"Legislative Organization",
		"Executive",
		"Executive Organization",
		"Judicial",
		"Judicial Organization",
	}
}

func branchesShortList() []string {
	return []string{
		"Legislative",
		"Executive",
		"Judicial",
	}
}

//Info -
func (cp *CulturalProfile) Info() {
	fmt.Println("Social outlook:")
	cultList := coltOutLIST()
	for i := range cultList {
		fmt.Println(cultList[i], ":", cp.socialOutlook[cultList[i]])
	}
	fmt.Println("Tech Profile:", cp.techProfile)
	fmt.Println("Authority outlook:")
	fmt.Println("Branches division: ", cp.authorityBranches["division"])
	branches := branchesList()
	for i := range branches {
		fmt.Println(branches[i], ":", cp.authorityBranches[branches[i]])
	}

	fmt.Println("Law outlook:")
	fmt.Println("Law uniformity:", cp.lawUniformity)

}

//TechnologyProfile -
func (cp *CulturalProfile) TechnologyProfile() string {
	return cp.techProfile
}

func adjustDM(condition bool, increment int) int {
	if condition {
		return increment
	}
	return 0
}

func ehexToInt(s string) int {
	return TrvCore.EhexToDigit(s)
}

func (cp *CulturalProfile) fillSocialOutlook(world *World) {
	//progrAtt
	cp.socialOutlook[constant.PrGovr] = TrvCore.DigitToEhex(world.stat[constant.PrGovr])
	dm1 := 0
	dm1 += adjustDM(world.Stat(constant.PrPops) >= 6, 1)
	dm1 += adjustDM(world.stat[constant.PrPops] >= 9, 2)
	dm1 += adjustDM(world.stat[constant.PrLaws] >= ehexToInt("A"), 1)
	switch utils.RollDice("2d6", dm1) {
	case 2, 3:
		cp.socialOutlook["Progressiveness Attitude"] = "Radical"
	case 4, 5, 6, 7:
		cp.socialOutlook["Progressiveness Attitude"] = "Progressive"
	case 8, 9, 10:
		cp.socialOutlook["Progressiveness Attitude"] = "Conservative"
	default:
		cp.socialOutlook["Progressiveness Attitude"] = "Reactionary"
	}
	//progrAction
	dm2 := 0
	dm2 += adjustDM(cp.socialOutlook["Progressiveness Attitude"] == "Conservative", 3)
	dm2 += adjustDM(cp.socialOutlook["Progressiveness Attitude"] == "Reactionary", 3)
	switch utils.RollDice("2d6", dm2) {
	case 2, 3, 4, 5:
		cp.socialOutlook["Progressiveness Action"] = "Enterprising"
	case 6, 7, 8, 9:
		cp.socialOutlook["Progressiveness Action"] = "Advancing"
	case 10, 11, 12:
		cp.socialOutlook["Progressiveness Action"] = "Indifirent"
	default:
		cp.socialOutlook["Progressiveness Action"] = "Stagnant"
	}
	//progrAtt
	dm3 := 0
	dm3 += adjustDM(world.stat[constant.PrLaws] >= ehexToInt("A"), 1)
	switch utils.RollDice("2d6", dm3) {
	case 2, 3:
		cp.socialOutlook["Aggressiveness Attitude"] = "Expansionistic"
	case 4, 5, 6:
		cp.socialOutlook["Aggressiveness Attitude"] = "Competitive"
	case 7, 8, 9, 10:
		cp.socialOutlook["Aggressiveness Attitude"] = "Unagressive"
	default:
		cp.socialOutlook["Aggressiveness Attitude"] = "Passive"
	}
	//progrAction
	dm4 := 0
	dm4 += adjustDM(cp.socialOutlook["Aggressiveness Attitude"] == "Expansionistic", -2)
	dm4 += adjustDM(cp.socialOutlook["Aggressiveness Attitude"] == "Competitive", -1)
	dm4 += adjustDM(cp.socialOutlook["Aggressiveness Attitude"] == "Passive", 2)
	dm4 += adjustDM(world.stat[constant.PrLaws] >= ehexToInt("A"), 1)
	switch utils.RollDice("2d6", dm4) {
	case -1, 0, 1, 2, 3, 4:
		cp.socialOutlook["Aggressiveness Action"] = "Militant"
	case 5, 6, 7, 8:
		cp.socialOutlook["Aggressiveness Action"] = "Neutral"
	case 9, 10, 11:
		cp.socialOutlook["Aggressiveness Action"] = "Peaceable"
	default:
		cp.socialOutlook["Aggressiveness Action"] = "Conciliatory"
	}
	//progrAtt
	dm5 := 0
	dm5 += adjustDM(world.stat[constant.PrGovr] <= ehexToInt("2"), 1)
	dm5 += adjustDM(world.stat[constant.PrGovr] == ehexToInt("7"), 4)
	dm5 += adjustDM(world.stat[constant.PrGovr] == ehexToInt("F"), -1)
	dm5 += adjustDM(world.stat[constant.PrLaws] <= ehexToInt("4"), 1)
	dm5 += adjustDM(world.stat[constant.PrLaws] >= ehexToInt("A"), -1)
	switch utils.RollDice("2d6", dm5) {
	case 2, 3:
		cp.socialOutlook["Extensiveness Global"] = "Monolithic"
	case 4, 5, 6, 7:
		cp.socialOutlook["Extensiveness Global"] = "Harmonious"
	case 8, 9, 10, 11:
		cp.socialOutlook["Extensiveness Global"] = "Discordant"
	default:
		cp.socialOutlook["Extensiveness Global"] = "Fragmented"
	}
	//progrAction
	dm6 := 0

	dm6 += adjustDM(world.StarPort() == "A", -2)
	dm6 += adjustDM(world.StarPort() == "B", -1)
	dm6 += adjustDM(world.StarPort() == "D", 1)
	dm6 += adjustDM(world.StarPort() == "E", 2)
	dm6 += adjustDM(world.StarPort() == "X", 3)
	dm6 += adjustDM(cp.socialOutlook["Progressiveness Attitude"] == "Conservative", 2)
	dm6 += adjustDM(cp.socialOutlook["Progressiveness Attitude"] == "Reactionary", 4)
	dm6 += adjustDM(world.stat[constant.PrLaws] >= ehexToInt("A"), 1)
	switch utils.RollDice("2d6", dm6) {
	case -1, 0, 1, 2, 3:
		cp.socialOutlook["Extensiveness Interstellar"] = "Xenophilic"
	case 4, 5, 6, 7:
		cp.socialOutlook["Extensiveness Interstellar"] = "Friendly"
	case 8, 9, 10, 11:
		cp.socialOutlook["Extensiveness Interstellar"] = "Aloof"
	default:
		cp.socialOutlook["Extensiveness Interstellar"] = "Xenophobic"
	}

}

func determineAuthorityProfile(cp *CulturalProfile) string {
	profile := cp.socialOutlook[constant.PrGovr] + "-"
	switch cp.socialOutlook["Progressiveness Attitude"] {
	case "Radical":
		profile += "1"
	case "Progressive":
		profile += "2"
	case "Conservative":
		profile += "3"
	case "Reactionary":
		profile += "4"
	}
	switch cp.socialOutlook["Progressiveness Action"] {
	case "Enterprising":
		profile += "1"
	case "Advancing":
		profile += "2"
	case "Indifirent":
		profile += "3"
	case "Stagnant":
		profile += "4"
	}
	switch cp.socialOutlook["Aggressiveness Attitude"] {
	case "Expansionistic":
		profile += "1"
	case "Competitive":
		profile += "2"
	case "Unagressive":
		profile += "3"
	case "Passive":
		profile += "4"
	}
	switch cp.socialOutlook["Aggressiveness Action"] {
	case "Militant":
		profile += "1"
	case "Neutral":
		profile += "2"
	case "Peaceable":
		profile += "3"
	case "Conciliatory":
		profile += "4"
	}
	switch cp.socialOutlook["Extensiveness Global"] {
	case "Monolithic":
		profile += "1"
	case "Harmonious":
		profile += "2"
	case "Discordant":
		profile += "3"
	case "Fragmented":
		profile += "4"
	}
	switch cp.socialOutlook["Extensiveness Interstellar"] {
	case "Xenophilic":
		profile += "1"
	case "Friendly":
		profile += "2"
	case "Aloof":
		profile += "3"
	case "Xenophobic":
		profile += "4"
	}
	profile += "-"
	switch cp.authorityBranches["division"] {
	case "3-way division":
		profile += "2"
	case "2-way division":
		profile += "1"
	case "No division":
		profile += "0"
	}
	for _, v := range branchesShortList() {
		pVal := -1
		switch cp.authorityBranches[v+" Organization"] {
		case "Ruler":
			pVal = 0
		case "Elite Council":
			pVal = 1
		case "Several Councils":
			pVal = 2
		case "Demos":
			pVal = 3
		case "No authority":
			pVal = 4
		}
		if cp.authorityBranches[v] == "Other" {
			pVal = pVal + 10
		}
		profile += TrvCore.DigitToEhex(pVal)
	}
	return profile
}

func divisionOfAuthority() string {
	r := utils.RollDice("d6")
	div := ""
	switch r {
	case 1, 2:
		div = "3-way division"
	case 3, 4:
		div = "2-way division"
	case 5, 6:
		div = "No division"
	}
	return div
}

func representGovrAuth2Way() string {
	r := utils.RollDice("d6")
	rep := ""
	switch r {
	case 1, 2:
		rep = "Executive & Judicial"
	case 3, 4, 5:
		rep = "Executive & Legislative"
	case 6:
		rep = "Legislative & Judicial"
	}
	return rep
}

func representGovrAuth3Way() string {
	r := utils.RollDice("d6")
	rep := ""
	switch r {
	case 1, 2:
		rep = "Executive"
	case 3, 4:
		rep = "Legislative"
	case 5, 6:
		rep = "Judicial"
	}
	return rep
}

func rerollIfResult(table func() string, notWant ...string) string {
	i := 0
	for {
		suggest := table()
		met := false
		for i := range notWant {
			if suggest == notWant[i] {
				met = true
			}
		}
		if !met {
			return suggest
		}
		i++
		if i > 10000 {
			return suggest
		}
	}

}

func listTL() []string {
	return []string{
		tlHIGH,
		tlLOW,
		tlENERGY,
		tlCOMPandROBOT,
		tlCOMMUNICATION,
		tlMEDICAL,
		tlENVIROMENT,
		tlTRANSPORTland,
		tlTRANSPORTwater,
		tlTRANSPORTair,
		tlTRANSPORTspace,
		tlMILITARYpersonal,
		tlMILITARYheavy,
	}
}

func techLevelModifier() int {
	r := utils.RollDice("2d6")
	switch r {
	default:
		return 0
	case 2:
		return utils.RollDice("d6") * -1
	case 3:
		return -2
	case 4:
		return -1
	case 10:
		return 1
	case 11:
		return 2
	case 12:
		return utils.RollDice("d6")
	}
}

// func NewTechnologyProfile(world *World) (*TechnologyProfile, error) {
// 	profile := &TechnologyProfile{}
// 	constant.tLevel = make(map[string]int)
// 	tlCode := world.stat[constant.PrTL]
// 	for _, tlArea := range listTL() {
// 		switch tlArea {
// 		default:
// 			return nil, errors.New("Unknown Area: '" + tlArea + "'")
// 		case tlHIGH:
// 			constant.tLevel[tlHIGH] = tlCode
// 		case tlLOW:
// 			upLim := constant.tLevel[tlHIGH]
// 			loLim := constant.tLevel[tlHIGH] / 2
// 			dm := techLevelModifier()
// 			base := constant.tLevel[tlHIGH]
// 			constant.tLevel[tlLOW] = utils.BoundInt(base+dm, loLim, upLim)
// 		}
// 	}

// 	return profile, nil
// }

//NewCulturalProfile -
func NewCulturalProfile(world *World) *CulturalProfile {
	cp := &CulturalProfile{}
	cp.socialOutlook = make(map[string]string)
	cp.fillSocialOutlook(world)
	cp.authorityBranches = mapAuthority(world.stat[constant.PrGovr])
	cp.calculateLawUniformity(world)
	cp.determineTechnologyProfile(world)
	//rerollIfResult(test, "99", "00")
	cp.societyProfile = determineAuthorityProfile(cp)
	world.data["UGP"] = cp.societyProfile

	//fmt.Println("cpProfile TEST", cp.societyProfile)
	return cp
}

func (cp *CulturalProfile) calculateLawUniformity(world *World) {
	dm := adjustDM(cp.socialOutlook["Extensiveness Global"] == "Monolithic", 2)
	dm += adjustDM(world.stat[constant.PrLaws] >= TrvCore.EhexToDigit("A"), -1)
	r := utils.RollDice("2d6", dm)
	switch r {
	case 2, 3, 4, 5:
		cp.lawUniformity = "Personal"
	case 6, 7:
		cp.lawUniformity = "Territorial"
	default:
		cp.lawUniformity = "Undivided"
	}
}

//UTP -
func (world *World) UTP() string {
	if val, ok := world.data["UTP"]; ok {
		return val
	}
	cp := NewCulturalProfile(world)
	return cp.TechnologyProfile()
}

func (cp *CulturalProfile) determineTechnologyProfile(world *World) {
	if val, ok := world.data["UTP"]; ok {
		cp.techProfile = val
		return
	}
	tlMap := make(map[string]int)
	//limits
	tlMap[tlHIGH] = world.stat[constant.PrTL]

	tlMap[tlLOW+"_up"] = tlMap[tlHIGH]
	tlMap[tlLOW+"_lo"] = tlMap[tlHIGH] / 2

	tlMap[tlENERGY+"_up"] = tlMap[tlHIGH] + tlMap[tlHIGH]/5
	tlMap[tlENERGY+"_lo"] = tlMap[tlLOW+"_lo"]

	tlMap[tlCOMPandROBOT+"_up"] = tlMap[tlENERGY+"_up"]
	tlMap[tlCOMPandROBOT+"_lo"] = tlMap[tlCOMPandROBOT+"_up"] - 3

	tlMap[tlCOMMUNICATION+"_up"] = tlMap[tlENERGY+"_up"]
	tlMap[tlCOMMUNICATION+"_lo"] = tlMap[tlCOMMUNICATION+"_up"] - 3

	tlMap[tlMEDICAL+"_up"] = tlMap[tlENERGY+"_up"]
	tlMap[tlMEDICAL+"_lo"] = 0

	tlMap[tlENVIROMENT+"_up"] = tlMap[tlENERGY+"_up"]
	tlMap[tlENVIROMENT+"_lo"] = tlMap[tlENVIROMENT+"_up"] - 5

	tlMap[tlTRANSPORTland+"_up"] = tlMap[tlENERGY+"_up"]
	tlMap[tlTRANSPORTland+"_lo"] = tlMap[tlTRANSPORTland+"_up"] - 5

	tlMap[tlTRANSPORTspace+"_up"] = tlMap[tlENERGY+"_up"]
	tlMap[tlTRANSPORTspace+"_lo"] = tlMap[tlTRANSPORTspace+"_up"] - 3

	tlMap[tlMILITARYpersonal+"_up"] = tlMap[tlENERGY+"_up"]
	tlMap[tlMILITARYpersonal+"_lo"] = 0

	tlMap[tlMILITARYheavy+"_up"] = tlMap[tlENERGY+"_up"]
	tlMap[tlMILITARYheavy+"_lo"] = 0

	for key := range tlMap {
		if tlMap[key] < 0 {
			tlMap[key] = 0
		}
	}

	//actual values
	dm := techLevelModifier()
	dm += adjustDM(world.stat[constant.PrPops] <= 5, 1)
	dm += adjustDM(world.stat[constant.PrPops] >= 9, -1)
	dm += adjustDM(cp.socialOutlook["Extensiveness Global"] == "Monolithic", 1)
	dm += adjustDM(cp.socialOutlook["Extensiveness Global"] == "Discordant", -1)
	dm += adjustDM(cp.socialOutlook["Extensiveness Global"] == "Fragmented", -2)
	tlMap[tlLOW] = utils.BoundInt(tlMap[tlHIGH]+dm, tlMap[tlLOW+"_lo"], tlMap[tlLOW+"_up"])

	dm = techLevelModifier()
	tlMap[tlENERGY] = utils.BoundInt(tlMap[tlHIGH]+dm, tlMap[tlENERGY+"_lo"], tlMap[tlENERGY+"_up"])

	dm = techLevelModifier()
	dm += adjustDM(world.stat[constant.PrPops] <= 5, 1)
	dm += adjustDM(world.stat[constant.PrPops] >= 9, -1)
	tlMap[tlCOMPandROBOT] = utils.BoundInt(tlMap[tlHIGH]+dm, tlMap[tlCOMPandROBOT+"_lo"], tlMap[tlCOMPandROBOT+"_up"])

	dm = techLevelModifier()
	base := utils.Min(tlMap[tlENERGY], tlMap[tlCOMPandROBOT])
	tlMap[tlCOMMUNICATION] = utils.BoundInt(base+dm, tlMap[tlCOMMUNICATION+"_lo"], tlMap[tlCOMMUNICATION+"_up"])

	dm = techLevelModifier()
	dm += adjustDM(cp.socialOutlook["Extensiveness Interstellar"] == "Xenophilic", 1)
	tlMap[tlMEDICAL] = utils.BoundInt(tlMap[tlCOMPandROBOT]+dm, tlMap[tlMEDICAL+"_lo"], tlMap[tlMEDICAL+"_up"])

	dm = techLevelModifier()
	dm += adjustDM(utils.ListContains(world.tradeCodes, "Wa"), 1)
	dm += adjustDM(utils.ListContains(world.tradeCodes, "De"), 1)
	dm += adjustDM(matchValue(world.stat[constant.PrHydr], 0, 10), 1)
	dm += adjustDM(!matchValue(world.stat[constant.PrAtmo], 5, 6, 8), 1)
	tlMap[tlENVIROMENT] = utils.BoundInt(tlMap[tlHIGH]+dm, tlMap[tlENVIROMENT+"_lo"], tlMap[tlENVIROMENT+"_up"])

	dm = techLevelModifier()
	dm += adjustDM(matchValue(world.stat[constant.PrHydr], 10), -1)
	tlMap[tlTRANSPORTland] = utils.BoundInt(tlMap[tlENERGY]+dm, tlMap[tlTRANSPORTland+"_lo"], tlMap[tlTRANSPORTland+"_up"])

	dm = 0
	dm += adjustDM(matchValue(world.stat[constant.PrHydr], 0), -1)
	if tlMap[tlTRANSPORTland] >= 10 {
		tlMap[tlTRANSPORTwater] = tlMap[tlTRANSPORTland] + dm
	} else {
		dm = dm + techLevelModifier()
		tlMap[tlTRANSPORTwater] = utils.BoundInt(tlMap[tlTRANSPORTland]+dm, 0, tlMap[tlTRANSPORTland])
	}

	dm = 0
	if tlMap[tlTRANSPORTland] >= 10 {
		tlMap[tlTRANSPORTair] = tlMap[tlTRANSPORTland] + dm
	} else {
		dm = dm + techLevelModifier()
		tlMap[tlTRANSPORTair] = utils.BoundInt(tlMap[tlENERGY]+dm, 0, 9)
		if world.stat[constant.PrAtmo] == 0 || tlMap[tlTRANSPORTair] <= 2 {
			tlMap[tlTRANSPORTair] = 0
		}
	}

	dm = 0
	dm += adjustDM(world.StarPort() == "A", 1)
	dm += adjustDM(world.StarPort() == "B", 1)
	dm += adjustDM(cp.socialOutlook["Extensiveness Interstellar"] == "Xenophilic", 1)
	dm += adjustDM(cp.socialOutlook["Extensiveness Interstellar"] == "Friendly", 1)
	dm += adjustDM(cp.socialOutlook["Extensiveness Interstellar"] == "Aloof", -1)
	dm += adjustDM(cp.socialOutlook["Extensiveness Interstellar"] == "Xenophobic", -1)
	dm += adjustDM(true, techLevelModifier())
	base = utils.Min(tlMap[tlENERGY], tlMap[tlCOMPandROBOT])
	tlMap[tlTRANSPORTspace] = utils.BoundInt(base+dm, tlMap[tlTRANSPORTspace+"_lo"], tlMap[tlTRANSPORTspace+"_up"])
	if tlMap[tlTRANSPORTspace] <= 4 {
		tlMap[tlTRANSPORTspace] = 0
	}
	if world.StarPort() == "X" {
		tlMap[tlTRANSPORTspace] = utils.Min(tlMap[tlTRANSPORTspace+"_lo"], 8)
	}

	dm = techLevelModifier()
	dm += adjustDM(cp.socialOutlook["Aggressiveness Attitude"] == "Expansionistic", 1)
	dm += adjustDM(cp.socialOutlook["Aggressiveness Attitude"] == "Passive", -2)
	dm += adjustDM(cp.socialOutlook["Aggressiveness Action"] == "Militant", 1)
	dm += adjustDM(cp.socialOutlook["Aggressiveness Action"] == "Conciliatory", -1)
	tlMap[tlMILITARYpersonal] = utils.BoundInt(tlMap[tlENERGY]+dm, tlMap[tlMILITARYpersonal+"_lo"], tlMap[tlMILITARYpersonal+"_up"])

	dm = techLevelModifier()
	dm += adjustDM(cp.socialOutlook["Aggressiveness Attitude"] == "Expansionistic", 1)
	dm += adjustDM(cp.socialOutlook["Aggressiveness Attitude"] == "Passive", -2)
	dm += adjustDM(cp.socialOutlook["Aggressiveness Action"] == "Militant", 1)
	dm += adjustDM(cp.socialOutlook["Aggressiveness Action"] == "Conciliatory", -1)
	tlMap[tlMILITARYheavy] = utils.BoundInt(tlMap[tlTRANSPORTland]+dm, tlMap[tlMILITARYheavy+"_lo"], tlMap[tlMILITARYheavy+"_up"])

	list := listTL()
	for i := range list {

		cp.techProfile = cp.techProfile + TrvCore.DigitToEhex(tlMap[list[i]])
		if i == 1 {
			cp.techProfile = cp.techProfile + "-"
		}
		if i == 6 {
			cp.techProfile = cp.techProfile + "-"
		}
		if i == 10 {
			cp.techProfile = cp.techProfile + "-"
		}
		if i == 12 {
			cp.techProfile = cp.techProfile + "-"
		}
	}
	cp.techProfile = cp.techProfile + TrvCore.DigitToEhex(world.stat[constant.PrTL])
	world.data["UTP"] = cp.techProfile
}

func authorityBranches() []string {
	return []string{
		"Executive",
		"Legislative",
		"Judicial",
	}
}

func mapAuthority(govrCode int) map[string]string {
	authMap := make(map[string]string)
	authMap["division"] = divisionOfAuthority()
	prime := ""
	switch authMap["division"] {
	case "3-way division":
		prime = representGovrAuth3Way()
	case "2-way division":
		prime = representGovrAuth2Way()
	case "No division":
		prime = "Executive & Legislative & Judicial"
	}
	if strings.Contains(prime, "Executive") {
		authMap["Executive"] = "Representative"
		authMap["Executive Organization"] = repGovrOrganization(govrCode)
	} else {
		authMap["Executive"] = "Other"
		authMap["Executive Organization"] = organizationType()
	}
	if strings.Contains(prime, "Legislative") {
		authMap["Legislative"] = "Representative"
		authMap["Legislative Organization"] = repGovrOrganization(govrCode)
	} else {
		authMap["Legislative"] = "Other"
		authMap["Legislative Organization"] = organizationType()
	}
	if strings.Contains(prime, "Judicial") {
		authMap["Judicial"] = "Representative"
		authMap["Judicial Organization"] = repGovrOrganization(govrCode)
	} else {
		authMap["Judicial"] = "Other"
		authMap["Judicial Organization"] = organizationType()
	}
	return authMap
}

func organizationType() string {
	r := utils.RollDice("2d6")
	org := ""
	switch r {
	default:
		org = "Demos"
	case 3, 4, 5:
		org = "Elite Council"
	case 6, 7:
		org = "Ruler"
	case 8, 9, 10, 11:
		org = "Several Councils"
	}
	return org
}

func repGovrOrganization(govrCode int) string {
	switch govrCode {
	default:
		return "No authority"
	case 1:
		return organizationType()
	case 2:
		return "Demos"
	case 3:
		return utils.RandomFromList([]string{"Elite Council", "Elite Council", "Several Councils"})
	case 4:
		return organizationType()
	case 5:
		return organizationType()
	case 6:
		return organizationType()
	case 7:
		return organizationType() //TODO: это надо делать 1d6+1 раз для каждой нации
	case 8:
		return "Several Councils"
	case 9:
		return "Several Councils"
	case 10:
		return utils.RandomFromList([]string{"Ruler", "Ruler", "Ruler", "Ruler", "Ruler", "Elite Council"})
	case 11:
		return utils.RandomFromList([]string{"Ruler", "Ruler", "Ruler", "Ruler", "Ruler", "Elite Council"})
	case 12:
		return utils.RandomFromList([]string{"Elite Council", "Elite Council", "Several Councils"})
	case 13:
		return rerollIfResult(organizationType, "Demos")
	case 14:
		return rerollIfResult(organizationType, "Demos")
	case 15:
		return utils.RandomFromList([]string{"Elite Council", "Elite Council", "Several Councils"})
	}
}
