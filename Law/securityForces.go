package law

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/utils"

	"github.com/Galdoba/TR_Dynasty/constant"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/world"
)

var dicepool *dice.Dicepool

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
	dicepool = dice.New(utils.SeedFromString(world.UWP()))
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

func NewSecurityFromUWP(uwp string) *Security {
	w := world.FromUWP(uwp)
	return NewSecurity(&w)
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
	//roll := TrvCore.Roll2D(dm) + laws(world) - 7
	roll := dicepool.RollNext("2d6").DM(dm).Sum() + laws(world) - 7
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
	//roll := TrvCore.Roll2D(dm) + laws(world) - 7
	roll := dicepool.RollNext("2d6").DM(dm).Sum() + laws(world) - 7
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
	fmt.Println("Punk------------------------------", dm)
	if match(govr(world), 1, 9, 10, 12) {
		dm--
	}
	fmt.Println("Punk------------------------------", dm)
	if match(govr(world), 6) {
		dm += 2
	}
	fmt.Println("Punk------------------------------", dm)
	if match(world.TradeCodes(), "Lo", "Po") {

		dm -= 2
	}
	fmt.Println("Punk------------------------------", dm)
	if match(world.TradeCodes(), "Lt", "Ni") {
		dm--
	}
	fmt.Println("Punk------------------------------", dm)
	if match(world.TradeCodes(), "Ri") {
		dm++
	}
	fmt.Println("Punk------------------------------", dm)
	pbg := []byte(world.PBG())
	if match(string(pbg[2]), "0") {
		dm -= 2
	}
	fmt.Println("Punk------------------------------", dm)
	//roll := TrvCore.Roll2D(dm) + orbPrez - 7
	fmt.Println(dicepool)
	fmt.Println("Punk------------------------------", dm)

	roll := dicepool.RollNext("2d6").DM(dm).Sum() + orbPrez - 7
	fmt.Println(dicepool)
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
	//roll := TrvCore.Roll2D(dm) + laws(world)
	roll := dicepool.RollNext("2d6").DM(dm).Sum() + laws(world) - 7
	if roll < 0 {
		roll = 0
	}
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
		dicepool.RollNext("2d6").Sum() == 12 {
		codes = append(codes, "Cr")
	}
	if match(govr(world), 1, 3, 6, 8, 9, 11, 13, 14, 15) &&
		pops(world) >= 6 &&
		match(plpres, 1, 2, 3, 4, 5) &&
		dicepool.RollNext("2d6").Sum() >= 10 {
		codes = append(codes, "Co")
	}
	if match(govr(world), 4, 5, 6, 9, 11, 12, 13, 14, 15) &&
		pops(world) >= 5 &&
		//match(world.TradeCodes(), "Po", "Ri") &&
		plpres >= 5 &&
		dicepool.RollNext("2d6").Sum() >= 10 {
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
		if dicepool.RollNext("2d6").Sum() >= tn {
			codes = append(codes, "Ip")
		}
	}
	if match(govr(world), 3, 5, 6, 7, 11, 15) &&
		pops(world) >= 4 &&
		dicepool.RollNext("2d6").Sum() >= 10 {
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
		dicepool.RollNext("2d6").Sum() >= 5 {
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

func (sp *Security) String() string {
	str := "World Security Profile: " + sp.Profile() + "\n"
	str += "-----------------------\n"
	str += "Planetary presence: " + strconv.Itoa(sp.planetaryPresence) + "\n"
	str += "  Orbital presence: " + strconv.Itoa(sp.orbitalPresence) + "\n"
	str += "   System presence: " + strconv.Itoa(sp.systemPresence) + "\n"
	str += "   Security Stance: " + strconv.Itoa(sp.stance) + "\n"
	str += "-----------------------\n"
	str += "Codes:\n"
	if len(sp.securityCodes) == 0 {
		str += "NONE\n"
	}
	for i := range sp.securityCodes {
		str += fullNameCode(sp.securityCodes[i]) + "\n"
	}
	return str
}

func fullNameCode(code string) string {
	switch code {
	default:
		return "Unknown"
	case "Cr":
		return "Corrupt: \nGraft, bribery, and self-interest are extremely common in the ranks of the security officers. Travellers should expect fair treatment only if it benefits the officers – or if they can pay for it."
	case "Co":
		return "Covert: \nWhilst most worlds have small covert security forces, this world’s security is predominantly hidden and consists of extensive surveillance, and in some societies a network of citizen informants."
	case "Fa":
		return "Factionalised: \nSecurity forces are numerous and often hold very specific mandates. This can lead to inefficiency and bureaucratic infighting that can inconvenience (or be exploited by) the Travellers."
	case "Fo":
		return "Focussed: \nThe strongest security and enforcement is found around key locations and people, with the rest of the world or system having much less. High Presence values with the Focussed code can mean extensive passive monitoring, with significant resources available when needed."
	case "Ip":
		return "Impersonal: \nThe security forces are less concerned with individual rights and justice, and more with the laws themselves and public order. A Difficult (10+) Advocate check can reverse the negative DM on sentencing rolls on these worlds, as Travellers use the letter of the law to their favour."
	case "Mi":
		return "Militarised: \nAll key security forces are military in nature. Typically more heavily armed and armoured than civilian security forces, they will normally be granted significant latitude by the government."
	case "Pe":
		return "Pervasive: \nSecurity apparatus is wide-ranging and common. This can vary from constant data-mining of computer networks, to a panopticon of cameras and gunshot sensors, to guards on every door, depending on the Tech Level. Pervasive security may be limited to the planet alone, or reach beyond it."
	case "Te":
		return "Technological: \nMain security functions are automated, or heavily reliant on hardware and software. Fewer officers will be present, but cameras, drones, and other devices will be very common. "
	case "Vo":
		return "Volunteer: \nSecurity forces are made up of volunteers, perhaps led by one or two paid full-time officer(s). They will typically be less well-trained but are dedicated to their community."
	}
}
