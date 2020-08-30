package law

import (
	"strings"

	"github.com/Galdoba/TR_Dynasty/constant"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/world"
)

//Security - obj for describing state of security Forces of the World
type Security struct {
	profile           string
	planetaryPresence int
	orbitalPresence   int
	systemPresence    int
	stance            int
	securityCodes     []string
}

func pops(w *world.World) int {
	return TrvCore.EhexToDigit(w.PlanetaryData(constant.PrPops))
}

func govr(w *world.World) int {
	return TrvCore.EhexToDigit(w.PlanetaryData(constant.PrGovr))
}

func laws(w *world.World) int {
	return TrvCore.EhexToDigit(w.PlanetaryData(constant.PrLaws))
}

func size(w *world.World) int {
	return TrvCore.EhexToDigit(w.PlanetaryData(constant.PrSize))
}

func atmo(w *world.World) int {
	return TrvCore.EhexToDigit(w.PlanetaryData(constant.PrAtmo))
}

func tl(w *world.World) int {
	return TrvCore.EhexToDigit(w.PlanetaryData(constant.PrTL))
}

//NewSecurity - creates random obj to draw info from using World data
func NewSecurity(world *world.World) *Security {
	sp := &Security{}
	if pops(world) == 0 {
		return sp
	}
	if govr(world) == 0 && laws(world) == 0 {
		sp.profile = "S000-0"
		return sp
	}
	sp.planetaryPresence = calculatePlanetaryPresence(world)
	sp.orbitalPresence = calculateOrbitalPresence(world)
	sp.systemPresence = calculateSystemPresence(world, sp.orbitalPresence)
	sp.stance = calculateStanse(world)
	sp.securityCodes = assignSecurityCodes(world, sp.planetaryPresence)
	sp.formProfile(world)
	return sp
}

//NewSecurityFromProfile - creates fixed obj using data from profile
func NewSecurityFromProfile(profile string) *Security {
	if len(profile) < 5 {
		return nil
	}
	sp := &Security{}
	sp.profile = profile
	data := []byte(profile)
	if string(data[0]) != "S" || string(data[4]) != "-" {
		return nil
	}
	sp.planetaryPresence = TrvCore.EhexToDigit(string(data[1]))
	sp.orbitalPresence = TrvCore.EhexToDigit(string(data[2]))
	sp.systemPresence = TrvCore.EhexToDigit(string(data[3]))
	sp.stance = TrvCore.EhexToDigit(string(data[5]))
	codes := strings.Split(profile, " ")
	for i := range codes {
		if i == 0 {
			continue
		}
		sp.securityCodes = append(sp.securityCodes, codes[i])
	}
	return sp
}

//Profile - returns string with Security Profile and Security Codes
func (sp *Security) Profile() string {
	if sp.profile == "" {
		return "[No Security]"
	}
	return sp.profile
}

func calculatePlanetaryPresence(world *world.World) int {
	dm := 0
	if match(size(world), 0, 1) {
		dm += 2
	}
	if match(size(world), 2, 3) {
		dm++
	}
	if match(size(world), 9, 10) {
		dm--
	}
	if match(govr(world), 6, 13, 14, 15) {
		dm += 2
	}
	if match(govr(world), 1, 5, 11) {
		dm++
	}
	if match(govr(world), 7, 10) {
		dm--
	}
	if match(govr(world), 2, 12) {
		dm -= 2
	}
	if match(world.TradeCodes(), "Ht", "Ri") {
		dm++
	}
	if match(world.TradeCodes(), "Lo") {
		dm--
	}
	if match(world.TradeCodes(), "Hi") {
		dm -= 2
	}
	roll := TrvCore.Roll2D(dm) + laws(world) - 7
	if roll < 0 {
		roll = 0
	}
	return roll
}

func calculateOrbitalPresence(world *world.World) int {
	if match(world.StarPort(), "X") {
		return 0
	}
	dm := 0
	if match(world.StarPort(), "E") {
		dm -= 2
	}
	if match(world.StarPort(), "D") {
		dm--
	}
	if match(world.StarPort(), "B") {
		dm++
	}
	if match(world.StarPort(), "A") {
		dm += 2
	}
	if match(size(world), 10, 11, 12) {
		dm--
	}
	if match(size(world), 3, 4) {
		dm++
	}
	if match(size(world), 0, 1, 2) {
		dm += 2
	}
	if match(govr(world), 2, 7, 12) {
		dm -= 2
	}
	if match(govr(world), 10) {
		dm--
	}
	if match(govr(world), 1, 5, 11) {
		dm++
	}
	if match(govr(world), 6, 13, 14, 15) {
		dm += 2
	}
	if match(world.TradeCodes(), "Lo", "Lt") {
		dm -= 2
	}
	if match(world.TradeCodes(), "Po") {
		dm--
	}
	if match(world.TradeCodes(), "Ag", "In", "Ht") {
		dm++
	}
	if match(world.TradeCodes(), "Ri") {
		dm += 2
	}
	if match(world.Bases(), "N") {
		dm++
	}
	roll := TrvCore.Roll2D(dm) + laws(world) - 7
	if roll < 0 {
		roll = 0
	}
	return roll
}

func calculateSystemPresence(world *world.World, orbPrez int) int {
	if match(world.StarPort(), "X") {
		return 0
	}
	dm := 0
	if match(world.StarPort(), "E") {
		dm -= 2
	}
	if match(world.StarPort(), "C", "D") {
		dm--
	}
	if match(world.StarPort(), "A") {
		dm++
	}
	if match(govr(world), 7) {
		dm -= 2
	}
	if match(govr(world), 1, 9, 10, 12) {
		dm--
	}
	if match(govr(world), 6) {
		dm += 2
	}
	if match(world.TradeCodes(), "Lo", "Po") {
		dm -= 2
	}
	if match(world.TradeCodes(), "Lt", "Ni") {
		dm--
	}
	if match(world.TradeCodes(), "Ri") {
		dm++
	}
	pbg := []byte(world.PBG())
	if match(string(pbg[2]), "0") {
		dm -= 2
	}
	roll := TrvCore.Roll2D(dm) + orbPrez - 7
	if roll < 0 {
		roll = 0
	}
	return roll
}

func calculateStanse(world *world.World) int {
	dm := 0
	if match(world.StarPort(), "X") {
		dm += 2
	}
	if match(atmo(world), 1, 10) {
		dm++
	}
	if match(atmo(world), 0, 11, 12) {
		dm += 2
	}
	if match(govr(world), 2, 12) {
		dm -= 2
	}
	if match(govr(world), 10) {
		dm--
	}
	if match(govr(world), 1, 5, 11) {
		dm++
	}
	if match(govr(world), 6, 13, 14, 15) {
		dm += 2
	}
	if match(world.TradeCodes(), "Hi") {
		dm -= 2
	}
	if match(world.TradeCodes(), "Ht") {
		dm--
	}
	if match(world.TradeCodes(), "Lt") {
		dm++
	}
	roll := TrvCore.Roll2D(dm) + laws(world)
	return roll
}

func match(val interface{}, chck ...interface{}) bool {
	switch val.(type) {
	default:
		return false
	case int, string:
		for _, checkVal := range chck {
			if val == checkVal {
				return true
			}
		}
		return false
	case []string:
		valSl := val.([]string)
		for i := range valSl {
			for _, checkVal := range chck {
				if valSl[i] == checkVal {
					return true
				}
			}
		}
		return false
	}
}

func assignSecurityCodes(world *world.World, plpres int) (codes []string) {
	if match(govr(world), 1, 3, 5, 6, 7, 8, 9, 11, 13, 14, 15) &&
		pops(world) >= 4 &&
		match(world.TradeCodes(), "Po", "Ri") &&
		match(plpres, 1, 2, 3, 4, 5) &&
		TrvCore.Roll2D() == 12 {
		codes = append(codes, "Cr")
	}
	if match(govr(world), 1, 3, 6, 8, 9, 11, 13, 14, 15) &&
		pops(world) >= 6 &&
		match(plpres, 1, 2, 3, 4, 5) &&
		TrvCore.Roll2D() >= 10 {
		codes = append(codes, "Co")
	}
	if match(govr(world), 4, 5, 6, 9, 11, 12, 13, 14, 15) &&
		pops(world) >= 5 &&
		//match(world.TradeCodes(), "Po", "Ri") &&
		plpres >= 5 &&
		TrvCore.Roll2D() >= 10 {
		codes = append(codes, "Fa")
	}
	if match(govr(world), 1, 6, 9, 10, 11, 12, 13, 14, 15) &&
		pops(world) >= 8 &&
		//match(world.TradeCodes(), "Po", "Ri") &&
		match(plpres, 1, 2, 3, 4, 5, 6) {
		//TrvCore.Roll2D() >= 10 {
		codes = append(codes, "Fo")
	}
	if match(govr(world), 1, 3, 6, 9, 13, 14, 15) &&
		pops(world) >= 5 {
		tn := 10
		if govr(world) == 9 {
			tn = 5
		}
		if TrvCore.Roll2D() >= tn {
			codes = append(codes, "Ip")
		}
	}
	if match(govr(world), 3, 5, 6, 7, 11, 15) &&
		pops(world) >= 4 &&
		TrvCore.Roll2D() >= 10 {
		codes = append(codes, "Mi")
	}
	if match(govr(world), 1, 5, 6, 8, 9, 11, 13, 14, 15) &&
		match(pops(world), 1, 2, 3, 4, 5, 6, 7, 8, 9) &&
		plpres >= 7 {
		codes = append(codes, "Pe")
	}
	if tl(world) >= 12 {
		codes = append(codes, "Te")
	}
	if match(govr(world), 2, 4, 7, 10, 12) &&
		match(pops(world), 1, 2) &&
		TrvCore.Roll2D() >= 5 {
		codes = append(codes, "Vo")
	}

	return codes
}

func (sp *Security) formProfile(world *world.World) {
	sp.profile = "S"
	sp.profile += TrvCore.DigitToEhex(sp.planetaryPresence)
	sp.profile += TrvCore.DigitToEhex(sp.orbitalPresence)
	sp.profile += TrvCore.DigitToEhex(sp.systemPresence)
	sp.profile += "-"
	sp.profile += TrvCore.DigitToEhex(sp.stance)
	if govr(world) == 7 {
		sp.profile += "B"
	}
	for i := range sp.securityCodes {
		sp.profile += " "
		sp.profile += sp.securityCodes[i]
	}
}
