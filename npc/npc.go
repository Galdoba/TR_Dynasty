package npc

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/devtools/cli"
	"github.com/Galdoba/utils"
)

const (
	occAverageCitizen    = "Average Citizen"
	occAdventurer        = "Adventurer"
	occArts              = "Arts"
	occBelter            = "Belter"
	occBountyHunter      = "Bounty Hunter"
	occCelebrity         = "Celebrity"
	occClergy            = "Clergy"
	occColonist          = "Colonist"
	occCorporateShipper  = "Corporate Shipper"
	occCraftperson       = "Craftperson"
	occDiplomaticService = "Diplomatic Service"
	occExplorer          = "Explorer"
	occFreeTrader        = "Free Trader"
	occFringeMarketer    = "Fringe Marketer"
	occGambler           = "Gambler"
	occGroundForces      = "Ground Forces"
	occInstructor        = "Instructor"
	occInvestigator      = "Investigator"
	occJournalist        = "Journalist"
	occMedic             = "Medic"
	occNavy              = "Navy"
	occOrganizedCrime    = "Organized Crime"
	occPirate            = "Pirate"
	occPolice            = "Police"
	occPolitican         = "Politican"
	occPortAuthority     = "Port Authority"
	occProstitute        = "Prostitute"
	occScavenger         = "Scavenger"
	occSports            = "Sports"
	occSpy               = "Spy"
	occThief             = "Thief"
	occVagabond          = "Vagabond"
)

// type NPCbuilder interface {
// 	SetRace(string)
// 	SetOccupation(string)
// 	GetNPC() *NPCensembleCast
// }

/*
-Race Human -Occupation Spy

*/

type NPCensembleCast struct {
	name           string
	race           string
	upp            string
	relations      int
	relationsBonus int
	quirk1         string
	quirk2         string
	occupation     string
	skillList      int
	skills         string
	age            int
}

func (npc *NPCensembleCast) String() string {
	str := ""
	str += "      Name: " + npc.name
	str += "\n       UPP: " + npc.upp
	str += "\n       Age: " + strconv.Itoa(18+4*(npc.skillList))
	str += "\n      Race: " + npc.race
	str += "\nOccupation: " + npc.occupation
	str += "\n    Quirk1: " + npc.quirk1
	if npc.quirk2 != "" {
		str += "\n    Quirk2: " + npc.quirk2
	}
	str += "\n    Skills: " + npc.skills
	//str += "\n Rel Bonus: " + strconv.Itoa(npc.relationsBonus)
	relation := dice.Roll("2d6").DM(npc.relationsBonus / 5).Sum()
	str += "\n  Reaction: " + describeReaction(relation)
	return str
}

func greetMsg() {
	fmt.Println("Valid Args:")
	fmt.Println("  '-help'      	Prints implemented keys of the program")
	fmt.Println("  '-help race' 	Prints implemented races for generator")
	fmt.Println("  '-help occ'  	Prints implemented occupations for generator")
	fmt.Println("  '-race [arg]'	Force NPC to have that specific race")
	fmt.Println("  '-occ [arg]' 	Force NPC to have that specific occupation")
	fmt.Println("***************************************************************")
	fmt.Println("")

}

func racesMsg() {
	fmt.Println("Valid Races:")
	fmt.Println("  'Human'          	No stat changes. Most Common")
	fmt.Println("  'Aslan'          	+2 STR; -2 DEX")
	fmt.Println("  'Vargr'          	-1 STR; +1 DEX; -1 END")
	fmt.Println("  'Floriani(Feskal)' 	+2 STR; +2 END; -2 INT; -2 EDU; -2 SOC")
	fmt.Println("  'Floriani(Barnai)'  	-2 STR; -2 END; +2 INT; +2 EDU")
	fmt.Println("***************************************************************")
	fmt.Println("")
	fmt.Println("")
}

func occMsg() {
	fmt.Println("Valid Occupations:")
	fmt.Println("   Average Citizen")
	fmt.Println("   Adventurer")
	fmt.Println("   Arts")
	fmt.Println("   Belter")
	fmt.Println("   Bounty Hunter")
	fmt.Println("   Celebrity")
	fmt.Println("   Clergy")
	fmt.Println("   Colonist")
	fmt.Println("   Corporate Shipper")
	fmt.Println("   Craftperson")
	fmt.Println("   Diplomatic Service")
	fmt.Println("   Explorer")
	fmt.Println("   Free Trader")
	fmt.Println("   Fringe Marketer")
	fmt.Println("   Gambler")
	fmt.Println("   Ground Forces")
	fmt.Println("   Instructor")
	fmt.Println("   Investigator")
	fmt.Println("   Journalist")
	fmt.Println("   Medic")
	fmt.Println("   Navy")
	fmt.Println("   Organized Crime")
	fmt.Println("   Pirate")
	fmt.Println("   Police")
	fmt.Println("   Politican")
	fmt.Println("   Port Authority")
	fmt.Println("   Prostitute")
	fmt.Println("   Scavenger")
	fmt.Println("   Sports")
	fmt.Println("   Spy")
	fmt.Println("   Thief")
	fmt.Println("   Vagabond")
	fmt.Println("***************************************************************")
}

func RandomNPC() NPCensembleCast {
	npc := NPCensembleCast{}
	npcArgs := cli.ArgsMap()
	npc.name = firstName() + " " + familyName()
	occ := ""
	if val, ok := npcArgs["-occ"]; ok {
		if len(val) > 0 {
			occ = val[0]
		}
	}
	if keys, ok := npcArgs["-help"]; ok {
		fmt.Println("***************************************************************")
		fmt.Println("Ensemble Cast: rollnpc v1.0.1")
		if len(keys) < 1 {
			greetMsg()
		}
		for i := range keys {
			if keys[i] == "race" {
				racesMsg()
			}
			if keys[i] == "occ" {
				occMsg()
			}
		}
		os.Exit(0)
	}
	npc.occupation = pickOccupation(occ)
	npc.rollStats()

	npc.rollRace()

	npc.skillTable()
	npc.quirk1 = TrvCore.RollD66()
	npc.quirk1 = npc.quirk(dice.RollD66())
	npc.quirk2 = ""
	if utils.RollDiceRandom("d2") > 1 {
		npc.quirk2 = npc.quirk(dice.RollD66())
	}
	npc.age = 18 + utils.RollDiceRandom("6d6")
	fmt.Print(npc.String())
	return npc
}

func (npc *NPCensembleCast) rollRace() {
	////////////////////////////////////Define races
	raceWeight := make(map[string]int)
	raceWeight["Human"] = 200
	raceWeight["Aslan"] = 40
	raceWeight["Vargr"] = 25
	raceWeight["Floriani(Feskal)"] = 15
	raceWeight["Floriani(Barnai)"] = 15
	totalWeight := 0
	var raceList []string
	for k, v := range raceWeight {
		for i := 0; i < v; i++ {
			raceList = append(raceList, k)
		}
		totalWeight = totalWeight + v
	}
	r := dice.Roll("1d" + strconv.Itoa(totalWeight)).DM(-1).Sum()
	race := raceList[r]
	/////////////////////////////Check os.args
	if cli.ArgExist("-race") {
		argmap := cli.ArgsMap()
		if vals, ok := argmap["-race"]; ok {
			check := ""
			if len(vals) > 0 {
				check = vals[0]
			}
			if _, ok := raceWeight[check]; ok {
				race = check
			}
		}
	}
	/////////////////////////////////Set values
	switch race {
	case "Human":
		npc.race = "Human"
	case "Aslan":
		npc.changeStat(1, 2)
		npc.changeStat(2, -2)
	case "Vargr":
		npc.changeStat(1, -1)
		npc.changeStat(2, 1)
		npc.changeStat(3, -1)
	case "Floriani(Feskal)":
		npc.changeStat(1, 2)
		npc.changeStat(3, 2)
		npc.changeStat(4, -2)
		npc.changeStat(5, -2)
		npc.changeStat(6, -2)
	case "Floriani(Barnai)":
		npc.changeStat(1, -2)
		npc.changeStat(3, -2)
		npc.changeStat(4, 2)
		npc.changeStat(5, 2)
	}
	npc.race = race
}

func FromTable(table table, keys ...string) string {
	return ""
}

type table struct {
	occ map[string][][]string
}

func newTable() table {
	tab := table{}
	tab.occ = make(map[string][][]string)
	tab.occ[occAverageCitizen] = [][]string{
		[]string{"2", "6", "7", "6", "6", "6", "6"},
		[]string{"3", "7", "6", "7", "6", "7", "6"},
		[]string{"4", "6", "7", "7", "6", "6", "6"},
		[]string{"5", "7", "7", "6", "6", "6", "7"},
		[]string{"6", "7", "7", "7", "6", "6", "6"},
		[]string{"7", "7", "7", "7", "7", "7", "7"},
		[]string{"8", "8", "7", "7", "8", "8", "7"},
		[]string{"9", "7", "8", "8", "7", "7", "7"},
		[]string{"A", "8", "8", "8", "7", "7", "7"},
		[]string{"B", "7", "7", "7", "8", "8", "8"},
		[]string{"C", "8", "8", "8", "8", "8", "8"},
	}
	tab.occ[occAdventurer] = tabConstructor([]string{"2", "72", "81", "73", "71", "61", "61"})
	tab.occ[occArts] = tabConstructor([]string{"2", "33", "43", "32", "73", "71", "72"})
	tab.occ[occBelter] = tabConstructor([]string{"2", "63", "71", "72", "71", "51", "31"})
	tab.occ[occBountyHunter] = tabConstructor([]string{"2", "63", "73", "71", "72", "62", "31"})
	tab.occ[occCelebrity] = [][]string{
		[]string{"2", "6", "7", "6", "6", "6", "8"},
		[]string{"3", "7", "6", "7", "6", "7", "8"},
		[]string{"4", "6", "7", "7", "6", "6", "8"},
		[]string{"5", "7", "7", "6", "6", "6", "8"},
		[]string{"6", "7", "7", "7", "6", "6", "9"},
		[]string{"7", "7", "7", "7", "7", "7", "9"},
		[]string{"8", "8", "7", "7", "8", "8", "9"},
		[]string{"9", "7", "8", "8", "7", "7", "9"},
		[]string{"A", "8", "8", "8", "7", "7", "A"},
		[]string{"B", "7", "7", "7", "8", "8", "A"},
		[]string{"C", "8", "8", "8", "8", "8", "B"},
	}
	tab.occ[occClergy] = tabConstructor([]string{"2", "62", "53", "61", "71", "71", "61"})
	tab.occ[occColonist] = [][]string{
		[]string{"2", "6", "7", "7", "6", "6", "5"},
		[]string{"3", "7", "6", "7", "7", "6", "5"},
		[]string{"4", "6", "7", "8", "7", "6", "5"},
		[]string{"5", "7", "7", "8", "8", "7", "6"},
		[]string{"6", "7", "7", "8", "8", "7", "6"},
		[]string{"7", "7", "7", "9", "8", "7", "6"},
		[]string{"8", "8", "7", "9", "9", "8", "7"},
		[]string{"9", "7", "8", "9", "9", "8", "7"},
		[]string{"A", "8", "8", "A", "9", "8", "7"},
		[]string{"B", "7", "7", "A", "A", "9", "8"},
		[]string{"C", "8", "8", "A", "A", "9", "8"},
	}
	tab.occ[occCorporateShipper] = tabConstructor([]string{"2", "51", "63", "52", "71", "61", "61"})
	tab.occ[occCraftperson] = [][]string{
		[]string{"2", "7", "7", "7", "6", "5", "4"},
		[]string{"3", "8", "7", "8", "6", "5", "4"},
		[]string{"4", "8", "7", "8", "6", "5", "4"},
		[]string{"5", "8", "8", "8", "7", "6", "4"},
		[]string{"6", "9", "8", "9", "7", "6", "5"},
		[]string{"7", "9", "8", "9", "7", "6", "5"},
		[]string{"8", "9", "9", "9", "8", "7", "5"},
		[]string{"9", "A", "9", "A", "8", "7", "6"},
		[]string{"A", "A", "9", "A", "8", "7", "6"},
		[]string{"B", "A", "A", "A", "9", "8", "7"},
		[]string{"C", "A", "A", "A", "9", "8", "7"},
	}
	tab.occ[occDiplomaticService] = [][]string{
		[]string{"2", "6", "7", "6", "7", "7", "7"},
		[]string{"3", "7", "6", "7", "8", "7", "8"},
		[]string{"4", "6", "7", "7", "8", "8", "8"},
		[]string{"5", "7", "7", "6", "8", "8", "8"},
		[]string{"6", "7", "7", "7", "9", "8", "9"},
		[]string{"7", "7", "7", "7", "9", "9", "9"},
		[]string{"8", "8", "7", "7", "9", "9", "9"},
		[]string{"9", "7", "8", "8", "A", "9", "A"},
		[]string{"A", "8", "8", "8", "A", "A", "B"},
		[]string{"B", "7", "7", "7", "A", "A", "B"},
		[]string{"C", "8", "8", "8", "B", "B", "C"},
	}
	tab.occ[occExplorer] = tabConstructor([]string{"2", "62", "72", "72", "63", "62", "52"})
	tab.occ[occFreeTrader] = tabConstructor([]string{"2", "62", "72", "61", "71", "63", "42"})
	tab.occ[occFringeMarketer] = tabConstructor([]string{"2", "41", "43", "42", "61", "51", "31"})
	tab.occ[occGambler] = tabConstructor([]string{"2", "51", "53", "61", "71", "52", "53"})
	tab.occ[occGroundForces] = tabConstructor([]string{"2", "72", "73", "73", "61", "53", "42"})
	tab.occ[occInstructor] = [][]string{
		[]string{"2", "4", "4", "4", "7", "7", "6"},
		[]string{"3", "4", "4", "4", "7", "7", "7"},
		[]string{"4", "4", "4", "5", "8", "8", "7"},
		[]string{"5", "5", "4", "5", "8", "8", "7"},
		[]string{"6", "5", "5", "5", "8", "8", "8"},
		[]string{"7", "5", "5", "6", "9", "9", "8"},
		[]string{"8", "6", "5", "6", "9", "9", "8"},
		[]string{"9", "6", "6", "6", "9", "9", "9"},
		[]string{"A", "6", "6", "7", "A", "A", "9"},
		[]string{"B", "7", "7", "7", "A", "A", "9"},
		[]string{"C", "7", "7", "7", "B", "B", "9"},
	}
	tab.occ[occInvestigator] = tabConstructor([]string{"2", "61", "61", "63", "72", "71", "63"})
	tab.occ[occJournalist] = tabConstructor([]string{"2", "42", "51", "51", "61", "63", "51"})
	tab.occ[occMedic] = [][]string{
		[]string{"2", "5", "6", "6", "7", "8", "7"},
		[]string{"3", "5", "7", "6", "7", "8", "7"},
		[]string{"4", "5", "7", "7", "7", "8", "7"},
		[]string{"5", "6", "7", "7", "8", "9", "8"},
		[]string{"6", "6", "8", "7", "8", "9", "8"},
		[]string{"7", "6", "8", "8", "8", "A", "9"},
		[]string{"8", "7", "8", "8", "9", "A", "9"},
		[]string{"9", "7", "9", "8", "9", "B", "A"},
		[]string{"A", "7", "9", "9", "9", "B", "A"},
		[]string{"B", "8", "9", "9", "A", "B", "B"},
		[]string{"C", "8", "A", "9", "A", "C", "B"},
	}
	tab.occ[occNavy] = [][]string{
		[]string{"2", "6", "7", "7", "6", "5", "4"},
		[]string{"3", "7", "7", "7", "6", "5", "4"},
		[]string{"4", "7", "8", "7", "6", "6", "4"},
		[]string{"5", "7", "8", "8", "7", "6", "5"},
		[]string{"6", "8", "8", "8", "7", "6", "5"},
		[]string{"7", "8", "9", "8", "7", "7", "6"},
		[]string{"8", "8", "9", "9", "8", "7", "6"},
		[]string{"9", "9", "9", "9", "8", "8", "7"},
		[]string{"A", "9", "A", "9", "8", "8", "8"},
		[]string{"B", "9", "A", "A", "9", "9", "9"},
		[]string{"C", "A", "A", "A", "9", "A", "A"},
	}
	tab.occ[occOrganizedCrime] = [][]string{
		[]string{"2", "6", "6", "7", "5", "4", "3"},
		[]string{"3", "6", "7", "7", "5", "4", "3"},
		[]string{"4", "7", "7", "7", "5", "4", "3"},
		[]string{"5", "7", "7", "8", "6", "5", "4"},
		[]string{"6", "7", "8", "8", "6", "5", "4"},
		[]string{"7", "8", "8", "8", "6", "5", "4"},
		[]string{"8", "8", "8", "9", "7", "6", "5"},
		[]string{"9", "8", "9", "9", "7", "6", "5"},
		[]string{"A", "9", "9", "9", "7", "7", "6"},
		[]string{"B", "9", "9", "A", "8", "7", "6"},
		[]string{"C", "9", "A", "A", "8", "8", "7"},
	}
	tab.occ[occPirate] = tabConstructor([]string{"2", "62", "72", "71", "41", "43", "23"})
	tab.occ[occPolice] = tabConstructor([]string{"2", "71", "72", "73", "51", "51", "51"})
	tab.occ[occPolitican] = [][]string{
		[]string{"2", "4", "4", "4", "7", "7", "8"},
		[]string{"3", "4", "4", "4", "7", "8", "8"},
		[]string{"4", "4", "4", "4", "7", "8", "9"},
		[]string{"5", "4", "5", "5", "8", "8", "9"},
		[]string{"6", "5", "5", "5", "8", "9", "9"},
		[]string{"7", "5", "5", "6", "8", "9", "9"},
		[]string{"8", "5", "6", "6", "9", "9", "9"},
		[]string{"9", "6", "6", "6", "9", "A", "A"},
		[]string{"A", "6", "6", "7", "9", "A", "A"},
		[]string{"B", "6", "6", "7", "A", "A", "B"},
		[]string{"C", "7", "7", "7", "A", "B", "B"},
	}
	tab.occ[occPortAuthority] = [][]string{
		[]string{"2", "5", "5", "5", "5", "5", "3"},
		[]string{"3", "6", "5", "6", "5", "5", "3"},
		[]string{"4", "6", "6", "6", "5", "5", "3"},
		[]string{"5", "7", "6", "7", "6", "6", "4"},
		[]string{"6", "7", "6", "7", "6", "6", "4"},
		[]string{"7", "8", "7", "8", "6", "6", "4"},
		[]string{"8", "8", "7", "8", "7", "7", "5"},
		[]string{"9", "8", "7", "9", "7", "7", "5"},
		[]string{"A", "9", "8", "9", "7", "7", "5"},
		[]string{"B", "9", "8", "9", "8", "8", "6"},
		[]string{"C", "9", "8", "A", "8", "8", "6"},
	}
	tab.occ[occProstitute] = tabConstructor([]string{"2", "51", "61", "63", "51", "41", "32"})
	tab.occ[occScavenger] = tabConstructor([]string{"2", "62", "71", "63", "61", "61", "41"})
	tab.occ[occSports] = [][]string{
		[]string{"2", "8", "8", "8", "5", "5", "4"},
		[]string{"3", "8", "8", "8", "5", "5", "5"},
		[]string{"4", "9", "9", "9", "5", "5", "6"},
		[]string{"5", "9", "9", "9", "6", "6", "7"},
		[]string{"6", "9", "9", "9", "6", "7", "7"},
		[]string{"7", "A", "A", "A", "6", "7", "7"},
		[]string{"8", "A", "A", "A", "7", "7", "8"},
		[]string{"9", "A", "A", "A", "7", "8", "8"},
		[]string{"A", "A", "B", "B", "7", "8", "9"},
		[]string{"B", "B", "B", "B", "8", "8", "A"},
		[]string{"C", "B", "B", "B", "8", "8", "B"},
	}
	tab.occ[occSpy] = [][]string{
		[]string{"2", "6", "7", "7", "7", "6", "5"},
		[]string{"3", "7", "7", "7", "7", "6", "5"},
		[]string{"4", "7", "8", "7", "7", "6", "5"},
		[]string{"5", "7", "8", "8", "8", "7", "6"},
		[]string{"6", "8", "8", "8", "8", "7", "7"},
		[]string{"7", "8", "9", "8", "8", "7", "7"},
		[]string{"8", "8", "9", "9", "9", "8", "7"},
		[]string{"9", "9", "9", "9", "9", "8", "8"},
		[]string{"A", "9", "A", "9", "9", "8", "8"},
		[]string{"B", "9", "A", "A", "A", "9", "8"},
		[]string{"C", "A", "A", "A", "A", "9", "9"},
	}
	tab.occ[occThief] = tabConstructor([]string{"2", "41", "71", "43", "71", "51", "32"})
	tab.occ[occVagabond] = [][]string{
		[]string{"2", "4", "4", "7", "6", "4", "2"},
		[]string{"3", "4", "4", "7", "6", "4", "3"},
		[]string{"4", "4", "4", "8", "7", "4", "3"},
		[]string{"5", "5", "5", "8", "7", "4", "3"},
		[]string{"6", "5", "5", "8", "7", "5", "3"},
		[]string{"7", "5", "6", "8", "7", "5", "3"},
		[]string{"8", "6", "6", "8", "7", "5", "3"},
		[]string{"9", "6", "6", "9", "8", "5", "3"},
		[]string{"A", "6", "7", "9", "8", "5", "3"},
		[]string{"B", "7", "7", "9", "8", "6", "3"},
		[]string{"C", "7", "7", "A", "8", "7", "3"},
	}
	return tab
}

func tabConstructor(st []string) [][]string {
	var res [][]string
	str := fromCode(st[1])
	dex := fromCode(st[2])
	end := fromCode(st[3])
	inn := fromCode(st[4])
	edu := fromCode(st[5])
	soc := fromCode(st[6])
	for i := 0; i <= 10; i++ {
		res = append(res, []string{strconv.Itoa(i + 2), str[i], dex[i], end[i], inn[i], edu[i], soc[i]})
	}

	// for i := range st {
	// 	switch i {
	// 	case 0:
	// 		res = append(res, []string{"2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C"})
	// 	case 1, 2, 3, 4, 5, 6:
	// 		res = append(res, fromCode(st[i]))
	// 	}
	// }
	return res
}

func fromCode(s string) []string {
	var res []string
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("-----------ERROR")
	}
	val := i / 10
	it := i - (val * 10)
	for l := 0; l <= 10; l++ {
		if it > 3 {
			it = 1
			val++
		}
		res = append(res, TrvCore.DigitToEhex(val))
		it++
	}
	return res
}

func (npc *NPCensembleCast) rollStats() {
	tab := newTable()
	var upp string
	for i := 0; i < 6; i++ {
		r := dice.Roll("2d6").Sum() - 2
		c := dice.Roll("d6").Sum()
		time.Sleep(time.Millisecond * 1)
		//fmt.Println(npc.occupation, r, c, tab.occ[npc.occupation][r][c])
		upp = upp + tab.occ[npc.occupation][r][c]

	}
	npc.upp = upp

}

func pickOccupation(occ string) string {
	occupation := []string{
		occAverageCitizen,
		occAdventurer,
		occArts,
		occBelter,
		occBountyHunter,
		occCelebrity,
		occClergy,
		occColonist,
		occCorporateShipper,
		occCraftperson,
		occDiplomaticService,
		occExplorer,
		occFreeTrader,
		occFringeMarketer,
		occGambler,
		occGroundForces,
		occInstructor,
		occInvestigator,
		occJournalist,
		occMedic,
		occNavy,
		occOrganizedCrime,
		occPirate,
		occPolice,
		occPolitican,
		occPortAuthority,
		occProstitute,
		occScavenger,
		occSports,
		occSpy,
		occThief,
		occVagabond,
	}
	for i := range occupation {
		if occ != occupation[i] {
			continue
		}
		return occupation[i]
	}
	return occupation[rand.Intn(len(occupation))]
}

func (npc *NPCensembleCast) quirk(code string) string {
	//r := dice.RollD66()
	quirk := ""
	switch code {
	default:
		quirk = "No Quirk"
	case "11":
		quirk = "Abrasive. The NPC is annoying and tends to cause ill will among those around them. The sort of person who nitpicks holovids and trolls people in worldnet forums."
		npc.relationsBonus += -10
	case "12":
		quirk = "Egotistical. The NPC is vain and boasts about themselves a great deal. They are constantly speaking about themselves and their accomplishments."
		npc.relationsBonus += -10
	case "13":
		quirk = "Violent. The NPC is prone to violence. They will lash out physically against any perceived problem."
		npc.relationsBonus += -10
	case "14":
		quirk = "Cruel. The NPC willingly and knowingly causes pain and distress in others. They will often detect a weakness in a person’s psyche and exploit it for their own amusement."
		npc.relationsBonus += -10
	case "15":
		quirk = "Greedy. The NPC has a strong desire for wealth or profit. While all of us require money to survive and prosper, this NPC has placed the acquisition of wealth above all else."
		npc.relationsBonus += -5
	case "16":
		quirk = "Envious. The NPC has a constant discontent with themselves and their situation and an obsession with the advantages, success, or possessions of other people."
		npc.relationsBonus += -5
	case "21":
		quirk = "Lustful. The NPC is motivated by lust and will have strong sexual desires that will override other aspects of life."
		npc.relationsBonus += -5
	case "22":
		quirk = "Liar. The NPC tells a great many falsehoods. They will often make intentionally inaccurate statements about themselves or others."
		npc.relationsBonus += -5
	case "23":
		quirk = "Skeptical. The NPC approaches all subjects of interest with an attitude of doubt. They will find it difficult to believe anything without proof."
	case "24":
		quirk = "Superstitious. The NPC has a great many beliefs which are not based in reason or knowledge. They will often assign ominous significance to seemingly innocuous things, circumstances, or occurrences."
	case "25":
		quirk = "Lazy. The NPC is often idle and sluggish. They will frequently display an aversion to work or exertion."
	case "26":
		quirk = "Incapable. The NPC is incompetent at the career or job in which they find themselves. They will often cause problems due to this lack of talent."
	case "31":
		quirk = "Ambitious. The NPC is obsessed with achieving or obtaining success, power, or a specific goal often to the detriment of those around them."
	case "32":
		quirk = "Blunt. The NPC is abrupt in the manner and often shuns normal societal expectations of decorum."
		npc.relationsBonus += -5
	case "33":
		quirk = "Cautious. The NPC is concerned about the danger of many situations or activities which would normally not be considered a problem. They will often hesitate in situations which they feel are dangerous no matter if the situation is dangerous or not."
	case "34":
		quirk = "Ignorant. The NPC is lacking in knowledge in their field of endeavor or study. They are often uniformed or unaware of the situation which faces them."
		npc.changeStat(5, -2)
	case "35":
		quirk = "Naïve. The NPC has a lack of sophistication and judgement. They often do not understand the complexities of reality and do not understand that some may have ulterior motives."
		npc.changeStat(6, -2)
	case "36":
		quirk = "Introverted. The NPC is shy and often finds happiness in solitude or small groups. They are often concerned only with their own thoughts and feelings which can lead to a lack of interest in the thoughts and feelings of others."
		npc.changeStat(6, -2)
	case "41":
		quirk = "Energetic. The NPC has an abundance of energy. This can cause the NPC to not note dangers which might await them."
	case "42":
		quirk = "Intelligent. The NPC has a good understanding of ideas and situations and is quick to comprehend new ones which they might encounter. They are often fast witted and able to easily grasp solutions to problems."
		npc.changeStat(4, 2)
	case "43":
		quirk = "Extraverted. The NPC is outgoing and gregarious. They are often concerned mainly with the social environment and the dynamic of groups in which they belong or desire to belong."
		npc.changeStat(6, 2)
	case "44":
		quirk = "Conservative. The NPC desires to preserve existing conditions and institutions or wishes to restore older or traditional ones. They will dislike change."
	case "45":
		quirk = "Liberal. The NPC wishes to reform existing conditions and institutions and often wishes to create new ones. They will often be dissatisfied with the status quo."
	case "46":
		quirk = "Handy. The NPC is skillful with their hands. They will often be knowledgeable or experienced with physical labor or repair of devices."
		npc.skills = npc.skills + ", Mechanics 0"
	case "51":
		quirk = "Considerate. The NPC shows kindness and has a deep regard for the feelings or circumstances of others."
		npc.relationsBonus += 5
	case "52":
		quirk = "Peaceful. The NPC is not violent. They will often shun argumentative or hostile situations."
		npc.relationsBonus += 5
	case "53":
		quirk = "Imaginative. The NPC has an exceptional imagination. They will be capable of forming mental images of things which they have created or originated."
	case "54":
		quirk = "Dependable. The NPC is trustworthy and reliable. They are often loyal to their friends, family, and colleagues."
		npc.relationsBonus += 5
	case "55":
		quirk = "Diligent. The NPC has the will and stamina to accomplish a given task. They will often not be deterred by minor difficulties or monotony."
	case "56":
		quirk = "Austere. The NPC is self-disciplined and serious. They will shun luxury and excess."
	case "61":
		quirk = "Forgiving. The NPC is capable of easily forgiving. They will attempt to rid themselves of resentment and ill will concerning past transgressions."
		npc.relationsBonus += 10
	case "62":
		quirk = "Generous. The NPC is unselfish and often gives of their wealth, time, or abilities. They will often be willing to help those in need."
		npc.relationsBonus += 10
	case "63":
		quirk = "Honest. The NPC will avoid untruths and will be honorable in their principles and intentions. They will attempt to be fair to those around them."
		npc.relationsBonus += 10
	case "64":
		quirk = "Humble. The NPC is neither proud nor boastful. They are often courteous and respectful."
		npc.relationsBonus += 10
	case "65":
		quirk = "Courageous. The NPC is brave. They will often be able to face difficulties and dangers that others would seek to avoid."
		npc.relationsBonus += 10
	case "66":
		quirk = "Friendly. The NPC will be amicable and kind. They will often be willing to enter a positive relationship with those with which they are unfamiliar."
		npc.relationsBonus += 10
	}
	return quirk
}

func (npc *NPCensembleCast) changeStat(pos, val int) {
	bt := []byte(npc.upp)
	hex := string(bt[pos-1])
	hexInt := TrvCore.EhexToDigit(hex) + val
	hexInt = utils.BoundInt(hexInt, 1, 20)
	newUPP := ""
	for i := range bt {
		if i == pos-1 {
			newUPP = newUPP + TrvCore.DigitToEhex(hexInt)
		} else {
			newUPP = newUPP + string(bt[i])
		}
	}
	npc.upp = newUPP
}

func (npc *NPCensembleCast) skillTable() {
	npc.skillList = dice.Roll("2d6").Sum()
	switch npc.occupation {
	case occAverageCitizen:
		switch npc.skillList {
		case 2:
			npc.skills = "Electronics 0, Mechanic 0, Survival 0"
		case 3:
			npc.skills = "Broker 1, Trade (Any) 1"
		case 4:
			npc.skills = "Broker 1, Advocate 0"
		case 5:
			npc.skills = "Electronics (Computer)-1, Trade (Any) 1, Flyer 0"
		case 6:
			npc.skills = "Carouse 1, Flyer (Grav) 1, Persuade 1"
		case 7:
			npc.skills = "Electronics (Computers) 1, Flyer (Grav) 1, Trade (Any) 1"
		case 8:
			npc.skills = "Admin 1, Broker 1, Persuade 1"
		case 9:
			npc.skills = "Admin 1, Broker 0, Electronics 0"
		case 10:
			npc.skills = "Admin 1, Broker 1, Carouse 1"
		case 11:
			npc.skills = "Mechanic 1, Electronics 0, Trade 0"
		case 12:
			npc.skills = "Admin 1, Broker 1, Flyer (Grav) 1, Trade (Any) 1"
		}
	case occAdventurer:
		switch npc.skillList {
		case 2:
			npc.skills = "Drive (Tracked) 1, Investigate 1, Recon 1, Survival 0, Navigation 0, Flyer 0"
		case 3:
			npc.skills = "Animals (Riding) 1, Gun Combat (Slug) 1, Recon 1, Stealth 1, Survival (Any) 1, Drive 0, Navigation 0"
		case 4:
			npc.skills = "Drive (Wheeled) 1, Gun Combat (Slug) 1, Investigate 1, Navigation 1, Recon 1, Survival (Any) 1, Survival (Any) 1"
		case 5:
			npc.skills = "Carouse 1, Flyer (Grav) 1, Gun Combat (Slug) 1, Investigate 1, Navigation 1, Recon 1, Science (Archaeology) 1, Survival (Any) 1, Survival (Any) 1"
		case 6:
			npc.skills = "Investigate 2, Recon 2, Drive (Wheeled) 1, Flyer (Grav) 1, Navigation 1, Science (Planetology) 1, Survival (Any) 1, Survival (Any) 1, Gun Combat 0"
		case 7:
			npc.skills = "Drive (Any) 2, Flyer (Grav) 2, Investigate 2, Navigation 1, Recon 1, Science (Planetology) 1, Survival (Any) 1, Survival (Any) 1"
		case 8:
			npc.skills = "Gun Combat (Slug) 2, Tactics (Military) 2, Animals (Riding) 1, Navigation 1, Recon 1, Stealth 1, Survival (Any) 1, Survival (Any) 1"
		case 9:
			npc.skills = "Navigation 2, Recon 2, Survival (Any) 2, Gun Combat (Slug) 1, Survival (Any) 1, Survival (Any) 1, Seafarer (Any) 1, Drive 0, Flyer 0"
		case 10:
			npc.skills = "Science (Archaeology) 2, Science (Planetology) 2, Carouse 1, Electronics (Sensors) 1, Investigate 1, Recon 1, Gun Combat 0, Survival 0"
		case 11:
			npc.skills = "Survival (Any) 3, Survival (Any) 3, Navigation 2, Recon 2, Survival (Any) 2, Electronics (Sensors) 1, Jack of All Trades 1, Medic (First Aid) 1, Survival (Any) 1, Gun Combat 0"
		case 12:
			npc.skills = "Gun Combat (Slug) 3, Stealth 3, Carouse 2, Navigation 2, Recon 2, Survival (Any) 2, Survival (Any) 2, Electronics 0"
		}
	case occArts:
		switch npc.skillList {
		case 2:
			npc.skills = "Art (Any) 1, Carouse 1, Persuade 1, Advocate 0, Language 0, Investigate 0"
		case 3:
			npc.skills = "Art (Painting) 1, Art (Sculpting) 1, Carouse 1, Persuade 1, Broker 0, Language 0"
		case 4:
			npc.skills = "Admin 1, Broker 1, Advocate (Legal) 1, Investigate 1, Persuade 1, Carouse 0"
		case 5:
			npc.skills = "Art (Writing) 1, Advocate (Oratory) 1, Investigate 1, Persuade 1, Broker 0, Carouse 0"
		case 6:
			npc.skills = "Art (Painting) 1, Carouse 1, Persuade 1, Language (Any) 1, Investigate 1, Advocate 0"
		case 7:
			npc.skills = "Art (Painting) 2, Carouse 2, Art (Sculpting) 1, Broker 1, Persuade 1, Language (Any) 1"
		case 8:
			npc.skills = "Broker 2, Investigation 2, Science (Art History) 2, Carouse 1, Language (Any) 1, Language (Any) 1, Persuade 1"
		case 9:
			npc.skills = "Advocate (Oratory) 2, Art (Writing) 2, Investigate 2, Persuade 2, Carouse 1, Language (Latin) 1, Language (Any) 1"
		case 10:
			npc.skills = "Art (Sculpting) 2, Art (Painting) 2, Carouse 2, Persuade 2, Broker 1, Science (Art History) 1, Language (Any) 1, Electronics 0"
		case 11:
			npc.skills = "Broker 3, Carouse 2, Investigation 2, Persuade 2, Art (Writing) 1, Art (Painting) 1, Science (Art History) 1, Language (Any) 1"
		case 12:
			npc.skills = "Art (Painting) 3, Carouse 3, Art (Sculpting) 2, Persuade 2, Jack of All Trades 1, Diplomat 1, Electronics 0, Investigate 0"
		}
	case occBelter:
		switch npc.skillList {
		case 2:
			npc.skills = "Drive (Mole) 1, Explosives 1, Trade (Prospector) 1, Vacc Suit 1, Zero-G 1, Pilot 0"
		case 3:
			npc.skills = "Trade (Prospector) 2, Carouse 1, Explosives 1, Vacc Suit 1, Zero-G 1, Drive 0, Pilot 0"
		case 4:
			npc.skills = "Trade (Prospector) 2, Admin 1, Electronics (Sensors) 1, Science (Geology) 1, Vacc Suit 1, Zero-G 1, Drive 0, Pilot 0"
		case 5:
			npc.skills = "Admin 1, Carouse 1, Steward 1, Mechanic 1, Vacc Suit 1, Zero-G 1, Electronics 0, Science 0"
		case 6:
			npc.skills = "Drive (Mole) 2, Explosives 2, Pilot (Small Craft) 1, Trade (Prospector) 1, Vacc Suit 1, Zero-G 1"
		case 7:
			npc.skills = "Trade (Prospector) 2, Explosives 2, Science (Geology) 2, Carouse 1, Drive (Mole) 1, Electronics (Sensors) 1, Vacc Suit 1, Zero-G 1"
		case 8:
			npc.skills = "Admin 2, Electronics (Sensors) 2, Trade (Prospector) 2, Science (Geology) 1, Vacc Suit 1, Zero-G 1"
		case 9:
			npc.skills = "Trade (Prospector) 2, Advocate (Oratory) 2, Carouse 2, Persuade 2, Advocate (Politics) 1, Language (Any) 1, Leadership 1, Science (Geology) 1, Vacc Suit 1, Zero-G 1"
		case 10:
			npc.skills = "Trade (Prospector) 2, Streetwise 2, Science (Geology) 1, Explosives 1, Vacc Suit 1, Zero-G 1, Pilot 0"
		case 11:
			npc.skills = "Trade (Prospector) 3, Science (Geology) 3, Language (Any) 2, Leadership 2, Persuade 2, Zero-G 2, Vacc Suit 1, Diplomat 1, Electronics 0, Pilot 0"
		case 12:
			npc.skills = "Trade (Prospector) 3, Science (Geology) 3, Explosives 2, Zero-G 2, Vacc Suit 2, Persuade 1, Streetwise 1, Electronics (Sensors) 1, Mechanic 0, Pilot 0"
		}
	case occBountyHunter:
		switch npc.skillList {
		case 2:
			npc.skills = "Deception (Intrusion) 1, Gun Combat (Slug) 1, Investigate 1, Recon 1, Streetwise 1, Melee 0, Tactics 0"
		case 3:
			npc.skills = "Investigate 1, Recon 1, Streetwise 1, Stealth 1, Melee (Unarmed Combat) 1, Tactics (Military) 1, Gun Combat 0, Survival 0"
		case 4:
			npc.skills = "Deception (Intrusion) 1, Electronics (Sensors) 1, Flyer (Grav) 1, Pilot (Spacecraft) 1, Streetwise 1, Gun Combat 0, Melee 0, Tactics 0"
		case 5:
			npc.skills = "Advocate (Legal) 1, Broker 1, Electronics (Comms) 1, Electronics (Computers) 1, Electronics (Sensors) 1, Flyer (Grav) 1, Streetwise 1, Gun Combat 0"
		case 6:
			npc.skills = "Investigate 2, Recon 2, Streetwise 2, Deception (Intrusion) 1, Deception 1, Melee (Unarmed Combat) 1, Gun Combat (Slug) 1, Tactics (Military) 1"
		case 7:
			npc.skills = "Recon 2, Stealth 2, Streetwise 2, Melee (Unarmed Combat) 2, Investigate 1, Flyer (Grav) 1, Tactics (Military) 1, Advocate (Legal) 1, Electronics 0, Medic 0"
		case 8:
			npc.skills = "Deception (Intrusion) 2, Drive (Wheeled) 2, Flyer (Grav) 2, Streetwise 2, Stealth 1, Investigate 1, Advocate (Legal) 1, Electronics (Computers) 1, Electronics (Sensors) 1, Gun Combat 0, Melee 0"
		case 9:
			npc.skills = "Admin 2, Advocate (Legal) 2, Broker 2, Electronics (Comms) 2, Electronics (Sensors) 2, Investigate 2, Recon 1, Stealth 1, Gun Combat 0, Melee 0, Flyer 0"
		case 10:
			npc.skills = "Investigate 2, Advocate (Legal) 2, Diplomat 1, Electronics (Computers) 1, Streetwise 1, Medic (First Aid) 1, Melee (Unarmed Combat) 1, Gun Combat (Slug) 1, Stealth 0, Tactics 0"
		case 11:
			npc.skills = "Gun Combat (Slug) 3, Advocate (Legal) 2, Carouse 2, Interrogation (Questioning) 2, Investigate 2, Persuade 2, Recon 2, Admin 1, Melee 0, Stealth 0, Deception 0, Streetwise 0"
		case 12:
			npc.skills = "Carouse 2, Gambler 2, Investigate 2, Melee (Blade) 2, Stealth 2, Streetwise 2, Advocate (Legal) 1, Broker 1, Deception (Intrusion) 1, Flyer (Grav) 1, Gun Combat (Slug) 1, Persuade 1, Pilot 0, Tactics 0"
		}
	case occCelebrity:
		switch npc.skillList {
		case 2:
			npc.skills = "Art (Holography) 1, Art (Writing) 1, Electronics 0"
		case 3:
			npc.skills = "Art (Instrument) 1, Carouse 1, Persuade 1, Broker 0, Electronics 0"
		case 4:
			npc.skills = "Carouse 2, Diplomat 1, Persuade 1, Steward 1"
		case 5:
			npc.skills = "Art (Writing) 2, Electronics (Comms) 2, Electronics (Computers) 2, Advocate 0"
		case 6:
			npc.skills = "Art (Instrument) 2, Carouse 2, Persuade 1, Electronics 0"
		case 7:
			npc.skills = "Art (Any) 2, Carouse 2, Deception (Disguise) 2, Persuade 2, Art (Any) 1, Electronics 0"
		case 8:
			npc.skills = "Art (Acting) 2, Deception (Disguise) 2, Admin 1, Carouse 1, Advocate 0, Electronics 0"
		case 9:
			npc.skills = "Persuade 2, Art (Writing) 1, Carouse 1, Advocate 0, Science 0"
		case 10:
			npc.skills = "Art (Holography) 2, Carouse 2, Electronics (Computers) 2, Admin 1, Broker 1, Persuade 0"
		case 11:
			npc.skills = "Art (Writing) 2, Science (History) 2, Persuade 2, Advocate 0"
		case 12:
			npc.skills = "Art (Acting) 3, Art (Writing) 2, Art (Holography) 1, Carouse 1, Persuade 1, Advocate 0"
		}
	case occClergy:
		switch npc.skillList {
		case 2:
			npc.skills = "Admin 1, Advocate (Oratory) 1, Diplomat 1, Persuade 1, Language 0, Streetwise 0"
		case 3:
			npc.skills = "Advocate (Oratory) 1, Carouse 1, Electronics (Comms) 1, Persuade 1, Drive 0, Flyer 0, Language 0"
		case 4:
			npc.skills = "Admin 1, Advocate (Oratory) 1, Carouse 1, Electronics (Comms) 1, Diplomat 1, Persuade 1, Science (Psychology) 1, Streetwise 1"
		case 5:
			npc.skills = "Admin 1, Art (Writing) 1, Language (Latin) 1, Diplomat 1, Investigate 1, Persuade 1, Steward 1 "
		case 6:
			npc.skills = "Admin 2, Advocate (Oratory) 2, Diplomat 2, Language (Any) 1, Persuade 1, Streetwise 1, Steward 1"
		case 7:
			npc.skills = "Advocate (Oratory) 2, Carouse 2, Persuade 2, Admin 1, Art (Writing) 1, Diplomat 1, Electronics (Comms) 1, Language (Any) 1, Language (Any) 1, Streetwise 1"
		case 8:
			npc.skills = "Admin 2, Advocate (Oratory) 2, Advocate (Politics) 2, Carouse 2, Persuade 2, Language (Any) 2, Language (Any) 2, Diplomat 1, Electronics (Comms) 1, Streetwise 1, Steward 1"
		case 9:
			npc.skills = "Art (Writing) 2, Science (History) 2, Investigate 2, Language (Greek) 2, Language (Hebrew) 2, Steward 1"
		case 10:
			npc.skills = "Advocate (Oratory) 3, Science (Psychology) 3, Persuade 3, Carouse 2, Diplomat 2, Language (Any) 2, Language (Any) 2, Streetwise 1, Steward 0, Pilot 0"
		case 11:
			npc.skills = "Admin 3, Persuade 3, Diplomat 3, Advocate (Oratory) 2, Science (Psychology) 2, Language (Any) 2, Drive 0"
		case 12:
			npc.skills = "Admin 3, Science (History) 3, Investigate 3, Melee (Unarmed Combat) 2, Art (Painting) 2, Electronics (Computers) 2, Advocate 0, Streetwise 0"
		}
	case occColonist:
		switch npc.skillList {
		case 2:
			npc.skills = "Animals (Any) 1, Electronics (Comms) 1, Survival (Any) 1, Science 0, Tactics 0"
		case 3:
			npc.skills = "Drive (Wheeled) 1, Electronics (Computers) 1, Mechanic 1, Animals 0, Navigation 0"
		case 4:
			npc.skills = "Mechanic 1, Survival (Any) 1, Science 0"
		case 5:
			npc.skills = "Carouse 1, Diplomat 1, Persuade 1, Survival (Any) 1, Leadership 0"
		case 6:
			npc.skills = "Admin 1, Broker 1, Electronics (Computers) 1, Survival (Any) 1"
		case 7:
			npc.skills = "Drive (Wheeled) 2, Gun Combat (Slug) 2, Survival (Any) 2, Navigation 1, Animals 0"
		case 8:
			npc.skills = "Admin 2, Broker 2, Carouse 1, Electronics (Computers) 1, Survival (Any) 1"
		case 9:
			npc.skills = "Gun Combat (Slug) 2, Survival (Any) 2, Recon 1, Tactics (Military) 1"
		case 10:
			npc.skills = "Animals (Farming) 2, Survival (Any) 2, Gun Combat (Slug) 1, Drive 0, Science 0"
		case 11:
			npc.skills = "Animals (Farming) 2, Animals (Riding) 2, Survival (Any) 2, Jack of All Trades 1, Gun Combat (Slug) 1, Navigation 1, Medic 0, Electronics 0"
		case 12:
			npc.skills = "Animals (Farming) 3, Animals (Riding) 2, Survival (Any) 2, Navigation 2, Jack of All Trades 2, Gun Combat (Slug) 1, Medic (First Aid) 1, Electronics 0, Flyer 0"
		}
	case occCorporateShipper:
		switch npc.skillList {
		case 2:
			npc.skills = "Admin 1, Broker 1, Carouse 1, Persuade 1, Vacc Suit 0, Zero-G 0"
		case 3:
			npc.skills = "Astrogation 1, Carouse 1, Electronics (Sensors) 1, Pilot (Spacecraft) 1, Drive 0, Flyer 0, Vacc Suit 	0, Zero-G 0"
		case 4:
			npc.skills = "Carouse 1, Engineer (M-Drive) 1, Engineer (Z-Drive) 1, Mechanic 1, Vacc Suit 0, Zero-G 0"
		case 5:
			npc.skills = "Gunner (Turrets) 1, Gun Combat (Energy) 1, Electronics (Sensors) 1, Tactics (Naval) 1, Persuade 1, Vacc Suit 0, Zero-G 0"
		case 6:
			npc.skills = "Admin 1, Astrogation 1, Broker 1, Carouse 1, Persuade 1, Pilot 0, Vacc Suit 0, Zero-G 0"
		case 7:
			npc.skills = "Pilot (Spacecraft) 2, Broker 2, Astrogation 1, Electronics (Sensors) 1, Mechanic 1, Vacc Suit 0, Zero-G 0"
		case 8:
			npc.skills = "Carouse 2, Engineer (M-Drive) 2, Engineer (Z-Drive) 2, Electronics (Computers) 1, Engineer (Life Support) 1, Mechanic 1, Vacc Suit 0, Zero-G 0"
		case 9:
			npc.skills = "Gunner (Turrets) 2, Gun Combat (Energy) 2, Carouse 1, Electronics (Sensors) 1, Tactics (Naval) 1, Vacc Suit 0, Zero-G 0"
		case 10:
			npc.skills = "Engineer (M-Drive) 3, Engineer (Z-Drive) 3, Carouse 2, Mechanic 2, Electronics (Computers) 2, Engineer (Life Support) 1, Engineer (Power) 1, Persuade 1, Diplomat 0, Vacc Suit 0, Zero-G 0"
		case 11:
			npc.skills = "Gunner (Turrets) 3, Carouse 2, Electronics (Sensors) 2, Gun Combat (Energy) 2 Tactics (Naval) 2, Persuade 1, Mechanic 1, Vacc Suit 0, Zero-G 0"
		case 12:
			npc.skills = "Broker 3, Admin 2, Astrogation 2, Pilot (Spacecraft) 2, Electronics (Sensors) 2, Persuade 2, Diplomat 1, Language (Any) 1, Language (Any) 1, Advocate (Legal) 1, Engineer 0, Tactics 0, Vacc Suit 0, Zero-G 0"
		}
	case occCraftperson:
		switch npc.skillList {
		case 2:
			npc.skills = "Carouse 1, Electronics (Computers) 1, Trade (Construction) 1, Engineer 0, Drive 0, Flyer 0"
		case 3:
			npc.skills = "Broker 1, Electronics (Computers) 1, Trade (Architect) 1, Persuade 1, Drive 0, Engineer 0"
		case 4:
			npc.skills = "Carouse 1, Trade (Construction) 1, Mechanic 1, Electronics (Electrical Repair) 1, Engineer (Power) 1, Persuade 1, Broker 0, Language 0"
		case 5:
			npc.skills = "Carouse 1, Drive (Wheeled) 1, Trade (Construction) 1, Jack of All Trades 1, Mechanic 0, Survival 0"
		case 6:
			npc.skills = "Trade (Architect) 2, Electronics (Computers) 2, Art (Writing) 1, Broker 1, Carouse 1, Engineer (Power) 1, Language (Any) 1, Persuade 1, Mechanic 0"
		case 7:
			npc.skills = "Trade (Construction) 2, Broker 1, Carouse 1, Electronics (Electrical Repair) 1, Engineer (Power) 1, Language (Any) 1, Mechanic 1, Advocate 0, Survival 0"
		case 8:
			npc.skills = "Trade (Construction) 2, Athletics (Endurance) 1, Broker 1, Carouse 1, Jack of All Trades 1, Language (Any) 1, Persuade 1, Flyer 0"
		case 9:
			npc.skills = "Broker 2, Trade (Construction) 2, Admin 1, Carouse 1, Language (Any) 1, Diplomat 1, Persuade 1, Streetwise 1"
		case 10:
			npc.skills = "Trade (Architect) 2, Art (Writing) 2, Broker 2, Carouse 2, Persuade 2, Diplomat 1, Language (Any) 1"
		case 11:
			npc.skills = "Trade (Construction) 3, Athletics (Endurance) 2, Broker 2, Carouse 2, Persuade 2, Jack of All Trades 2, Diplomat 1, Advocate (Politics) 1, Survival (Any) 1, Mechanic 1"
		case 12:
			npc.skills = "Trade (Construction) 3, Broker 3, Carouse 2, Persuade 2, Jack of All Trades 2, Mechanic 2, Language (Any) 1, Language (Any) 1, Flyer (Grav) 1, Survival (Any) 1"
		}
	case occDiplomaticService:
		switch npc.skillList {
		case 2:
			npc.skills = "Carouse 2, Diplomat 1, Electronics (Computers) 1, Language (Any) 1, Persuade 1"
		case 3:
			npc.skills = "Carouse 1, Diplomat 1, Electronics (Comms) 1, Persuade 1, Advocate 0, Language 0"
		case 4:
			npc.skills = "Stealth 2, Electronics (Comms) 1, Gun Combat (Energy) 1, Investigate 1"
		case 5:
			npc.skills = "Diplomat 2, Carouse 1, Language (Any) 2, Advocate 0, Electronics 0"
		case 6:
			npc.skills = "Carouse 1, Diplomat 1, Persuade 1, Advocate 0, Electronics 0"
		case 7:
			npc.skills = "Diplomat 2, Persuade 2, Language (Any) 2, Carouse 1, Electronics (Comms) 1, Advocate (Oratory) 1, Language (Any) 1, Language (Any) 1"
		case 8:
			npc.skills = "Advocate (Politics) 2, Diplomat 2, Language (Any) 2, Language (Any) 2, Carouse 1, Advocate (Oratory) 1, Electronics 0"
		case 9:
			npc.skills = "Advocate (Politics) 3, Language (Any) 3, Advocate (Oratory) 2, Carouse 2, Diplomat 2, Persuade 2, Science 0, Electronics 0, Survival 0"
		case 10:
			npc.skills = "Diplomat 2, Advocate (Oratory) 2, Advocate (Politics) 2, Carouse 2, Language (Any) 2, Language (Any) 2, Language (Any) 2, Language (Any) 2, Persuade 1"
		case 11:
			npc.skills = "Diplomat 2, Language (Any) 2, Language (Any) 2, Language (Any) 2, Language (Any) 2, Persuade 2. Advocate (Politics) 1, Advocate (Oratory) 1"
		case 12:
			npc.skills = "Diplomat 3, Advocate (Oratory) 2, Advocate (Politics) 2, Language (Any) 2, Language (Any) 2, Language (Any) 2, Language (Any) 2, Persuade 2, Carouse 1, Electronics 0"
		}
	case occExplorer:
		switch npc.skillList {
		case 2:
			npc.skills = "Astrogation 1, Electronics (Sensors) 1, Survival (Any) 1, Pilot (Spacecraft) 1, Navigation 0, Vacc Suit 0"
		case 3:
			npc.skills = "Electronics (Sensors) 1, Navigation 1, Survival (Any) 1, Gun Combat 0, Science 0, Vacc Suit 0"
		case 4:
			npc.skills = "Animals (Riding) 1, Navigation 1, Science (Any) 1, Survival (Any) 1, Flyer 0, Vacc Suit 0"
		case 5:
			npc.skills = "Astrogation 2, Pilot (Spacecraft) 1, Gunner (Turrets) 1, Gun Combat (Slug) 1, Science (Any) 1, Survival (Any) 1, Vacc Suit 0"
		case 6:
			npc.skills = "Electronics (Sensors) 2, Electronics (Computers) 1, Astrogation 1, Navigation 1, Gun Combat 0, Science 0, Vacc Suit 0, Zero-G 0"
		case 7:
			npc.skills = "Pilot (Spacecraft) 2, Electronics (Sensors) 2, Astrogation 1, Survival (Any) 1, Navigation 0, Gun Combat 0, Vacc Suit 0"
		case 8:
			npc.skills = "Electronics (Sensors) 2, Gun Combat (Energy) 2, Navigation 1, Recon 1, Gunner (Turrets) 1, Survival 0, Vacc Suit 0"
		case 9:
			npc.skills = "Electronics (Sensors) 2, Science (Planetology) 2, Science (Orbital Mechanics) 2, Pilot (Spacecraft) 1, Astrogation 1, Survival (Any) 1, Recon 1, Gun Combat 0, Tactics 0, Vacc Suit 0"
		case 10:
			npc.skills = "Astrogation 2, Science (Orbital Mechanics) 2, Pilot (Spacecraft) 2, Vacc Suit 1, Electronics 0"
		case 11:
			npc.skills = "Gun Combat (Energy) 2, Recon 2, Stealth 2, Navigation 1, Flyer (Grav) 1, Pilot (Small Craft) 1, Animals 0, Vacc Suit 0"
		case 12:
			npc.skills = "Pilot (Spacecraft) 3, Electronics (Sensors) 2, Astrogation 2, Survival (Any) 1, Navigation 0, Gun Combat 0, Vacc Suit 0"
		}
	case occFreeTrader:
		switch npc.skillList {
		case 2:
			npc.skills = "Astrogation 1, Broker 1, Persuade 1, Pilot (Spacecraft) 1, Zero-G 0, Vacc Suit 0"
		case 3:
			npc.skills = "Electronics (Sensors) 1, Broker 1, Persuade 1, Gun Combat 0, Science 0, Vacc Suit 0"
		case 4:
			npc.skills = "Broker 1, Persuade 1, Pilot (Spacecraft) 1, Survival (Any) 1, Flyer 0, Vacc Suit 0"
		case 5:
			npc.skills = "Broker 2, Carouse 2, Pilot (Spacecraft) 1, Persuade 1, Vacc Suit 0, Zero-G 0"
		case 6:
			npc.skills = "Admin 1, Broker 1, Carouse 1, Electronics (Computers) 1, Pilot 0, Vacc Suit 0, Zero-G 0"
		case 7:
			npc.skills = "Broker 2, Admin 1, Astrogation 1, Electronics 0, Pilot 0, Vacc Suit 0, Zero-G 0"
		case 8:
			npc.skills = "Admin 2, Broker 2, Carouse 1, Electronics (Comms) 1, Persuade 1, Astrogation 0, Vacc Suit 0, Zero-G 0"
		case 9:
			npc.skills = "Broker 2, Carouse 2, Admin 1, Advocate (Legal) 1, Persuade 1, Vacc Suit 0, Zero-G 0"
		case 10:
			npc.skills = "Broker 2, Carouse 2, Pilot (Spacecraft) 2, Astrogation 1, Electronics (Sensors) 1, Persuade 0, Vacc Suit 0"
		case 11:
			npc.skills = "Admin 3, Broker 2, Carouse 2, Persuade 2, Zero-G 0, Vacc Suit 0"
		case 12:
			npc.skills = "Broker 3, Persuade 3, Admin 2, Advocate (Legal) 2, Astrogation 1, Pilot (Spacecraft) 1, Zero-G 0, Vacc Suit 0"
		}
	case occFringeMarketer:
		switch npc.skillList {
		case 2:
			npc.skills = "Admin 1, Broker 1, Carouse 1, Persuade 1, Streetwise 1, Deception 0"
		case 3:
			npc.skills = "Admin 1, Broker 1, Deception (Forgery) 1, Gun Combat (Slug) 1, Streetwise 1, Recon 0"
		case 4:
			npc.skills = "Broker 1, Carouse 1, Deception (Lie) 1, Streetwise 1, Recon 1, Stealth 0"
		case 5:
			npc.skills = "Admin 1, Broker 1, Streetwise 1, Advocate 0, Gun Combat 0, Recon 0,"
		case 6:
			npc.skills = "Broker 2, Carouse 2, Persuade 2, Admin 1, Streetwise 1, Deception 1"
		case 7:
			npc.skills = "Broker 2, Streetwise 2, Persuade 2, Admin 1, Carouse 1, Deception (Forgery) 1, Deception 1"
		case 8:
			npc.skills = "Admin 2, Advocate (Legal) 2, Broker 2, Gun Combat (Slug) 1, Investigate 1, Deception 0, Electronics 0"
		case 9:
			npc.skills = "Admin 2, Broker 2, Carouse 2, Deception (Forgery) 2, Persuade 2, Deception 1, Gun Combat (Slug) 1, Investigate 0, Melee 0"
		case 10:
			npc.skills = "Broker 3, Persuade 2, Admin 2, Investigate 2, Electronics (Sensors) 1, Electronics (Computers) 1, Gun Combat (Slug) 1, Streetwise 1, Melee 0, Diplomat 0"
		case 11:
			npc.skills = "Broker 3, Streetwise 3, Admin 2, Recon 2, Deception (Forgery) 2, Deception 2, Persuade 2, Carouse 1, Gun Combat (Slug) 1, Melee (Blade) 1, Mechanic 0"
		case 12:
			npc.skills = "Broker 3, Carouse 3, Persuade 3, Streetwise 3, Deception 2, Admin 1, Gun Combat (Slug) 1, Melee (Blade) 1"
		}
	case occGambler:
		switch npc.skillList {
		case 2:
			npc.skills = "Broker 1, Carouse 1, Deception 1, Gambler 1, Recon 1, Investigate 0"
		case 3:
			npc.skills = "Deception 1, Gambler 1, Interrogation (Questioning) 1, Persuade 1, Science (Probability) 1, Streetwise 0"
		case 4:
			npc.skills = "Gambler 1, Investigate 1, Recon 1, Tactics (Sport) 1, Persuade 1, Streetwise 0"
		case 5:
			npc.skills = "Carouse 1, Gambler 1, Persuade 1, Streetwise 1, Animals 0, Mechanic 0"
		case 6:
			npc.skills = "Carouse 2, Gambler 2, Broker 1, Deception 1, Persuade 1, Recon 1, Investigate 0"
		case 7:
			npc.skills = "Gambler 2, Deception 2, Science (Probability) 2, Persuade 2, Science (Sociology) 1, Streetwise 1, Diplomat 0"
		case 8:
			npc.skills = "Gambler 2, Recon 2, Tactics (Sport) 2, Persuade 2, Investigate 1, Streetwise 1, Deception 0"
		case 9:
			npc.skills = "Gambler 2, Persuade 2, Streetwise 2, Carouse 1, Animals (Training) 1, Mechanic 1, Recon 1"
		case 10:
			npc.skills = "Gambler 3, Carouse 3, Deception 3, Interrogation (Questioning) 2, Persuade 2, Recon 2, Science (Probability) 2, Diplomat 1, Investigate 1, Science (Psychology) 1, Tactics 0"
		case 11:
			npc.skills = "Gambler 3, Carouse 2, Investigate 2, Recon 2, Persuade 2, Streetwise 2, Diplomat 1, Tactics (Sport) 1"
		case 12:
			npc.skills = "Carouse 3, Gambler 3, Animals (Training) 2, Investigate 2, Recon 2, Persuade 2, Streetwise 2, Tactics (Sport) 1, Gun Combat 0, Melee 0"
		}
	case occGroundForces:
		switch npc.skillList {
		case 2:
			npc.skills = "Carouse 1, Electronics (Computers) 1, Gun Combat (Slug) 1, Medic (First Aid) 1, Melee (Unarmed Combat) 1, Survival (Any) 1"
		case 3:
			npc.skills = "Carouse 1, Gun Combat (Slug) 1, Gun Combat (Energy) 1, Navigation 1, Melee (Blade) 1, Survival (Any) 1, Discipline 0, Battle Armor 0"
		case 4:
			npc.skills = "Forward Observer 1, Gun Combat (Slug) 1, Navigation 1, Persuade 1, Recon 1, Stealth 1, Medic 0"
		case 5:
			npc.skills = "Battle Armor 1, Carouse 1, Electronics (Sensors) 1, Explosives 1, Gun Combat (Slug) 1, Heavy Weapons (Launchers) 1, Recon 1"
		case 6:
			npc.skills = "Gun Combat (Slug) 2, Melee (Unarmed Combat) 2, Survival (Any) 2, Carouse 1, Discipline 1, Electronics (Computers) 1, Gun Combat (Energy) 1, Navigation 1, Recon 1, Medic (First Aid) 1, Melee (Blade) 1"
		case 7:
			npc.skills = "Carouse 2, Gun Combat (Slug) 2, Gun Combat (Energy) 2, Navigation 2, Survival (Any) 2, Battle Armor 1, Discipline 1, Melee (Blade) 1, Recon 1, Stealth 1, Medic 0"
		case 8:
			npc.skills = "Navigation 2, Recon 2, Stealth 2, Carouse 1, Discipline 1, Forward Observer 1, Medic (First Aid) 1, Persuade 1"
		case 9:
			npc.skills = "Battle Armor 2, Carouse 2, Gun Combat (Energy) 2, Heavy Weapons (Launchers) 2, Recon 2, Discipline 1, Electronics (Sensors) 1, Heavy Weapons (Field Artillery) 1, Melee 0, Medic 0"
		case 10:
			npc.skills = "Gun Combat (Slug) 3, Survival (Any) 3, Carouse 2, Melee (Unarmed Combat) 2, Recon 2, Survival (Any) 2, Discipline 1, Melee (Blade) 1, Medic (First Aid) 1, Navigation 1, Survival (Any) 1"
		case 11:
			npc.skills = "Navigation 3, Recon 3, Stealth 3, Discipline 2, Forward Observer 2, Melee (Blade) 2, Persuade 2, Medic 0, Electronics 0"
		case 12:
			npc.skills = "Battle Armor 3, Carouse 3, Gun Combat (Energy) 2, Heavy Weapons (Launchers) 2, Recon 2, Discipline 1, Electronics (Sensors) 1, Heavy Weapons (Field Artillery) 1, Melee 0, Medic 0"
		}
	case occInstructor:
		switch npc.skillList {
		case 2:
			npc.skills = "Admin 1, Art (Writing) 1, Instruction 1, Language (Any) 1, Persuade 1, Science (Any) 1"
		case 3:
			npc.skills = "Advocate (Oratory) 1, Art (Writing) 1, Carouse 1, Instruction 1, Persuade 1, Science (Any) 1"
		case 4:
			npc.skills = "Admin 1, Art (Any) 1, Electronics (Computers) 1, Electronics (Comms) 1, Instruction 1, Persuade 1, Science (Any)"
		case 5:
			npc.skills = "Admin 1, Art (Any) 1, Broker 1, Instruction 1, Persuade 1, Science (Any) 1"
		case 6:
			npc.skills = "Instruction 2, Persuade 2, Admin 1, Art (Writing) 1, Language (Any) 1, Language (Any) 1, Science (Any) 1"
		case 7:
			npc.skills = "Advocate (Oratory) 2, Art (Writing) 2, Instruction 2, Persuade 2, Science (Any) 1, Diplomat 1, Investigate 1"
		case 8:
			npc.skills = "Art (Any) 2, Instruction 2, Persuade 2, Science (Any) 2, Admin 1, Broker 1, Advocate 0, Diplomat 0"
		case 9:
			npc.skills = "Art (Any) 2, Instruction 2, Persuade 2, Science (Any) 2, Admin 1, Broker 1, Flyer (Grav) 1"
		case 10:
			npc.skills = "Admin 2, Art (Any) 2, Broker 2, Instruction 2, Persuade 2, Science (Any) 2"
		case 11:
			npc.skills = "Instruction 3, Persuade 3, Admin 2, Art (Writing) 2, Language (Any) 2, Science (Any) 2, Language (Any) 1, Language (Any) 1,"
		case 12:
			npc.skills = "Advocate (Oratory) 3, Art (Writing) 3, Instruction 3, Persuade 3, Science (Any) 2, Diplomat 1, Investigate 1"
		}
	case occInvestigator:
		switch npc.skillList {
		case 2:
			npc.skills = "Advocate (Legal) 1, Carouse 1, Investigate 1, Recon 1, Deception 0, Interrogation 0,"
		case 3:
			npc.skills = "Art (Holography) 1, Electronics (Computers) 1, Investigate 1, Medic (Diagnosis) 1, Science (Anatomy) 1, Science (Forensics) 1"
		case 4:
			npc.skills = "Advocate (Legal) 1, Carouse 1, Gun Combat (Slug) 1, Interrogation (Questioning) 1, Investigate 1, Recon 1, Streetwise 1"
		case 5:
			npc.skills = "Carouse 1, Gun Combat (Slug) 1, Interrogation (Questioning) 1, Melee (Unarmed Combat) 1, Recon 1, Stealth 1, Advocate 0, Deception 0"
		case 6:
			npc.skills = "Advocate (Legal) 2, Carouse 2, Interrogation (Questioning) 2, Investigate 2, Recon 1, Streetwise 1, Deception 0. Science 0"
		case 7:
			npc.skills = "Advocate (Legal) 2, Carouse 2, Interrogation (Questioning) 2, Investigate 2, Recon 2, Deception (Intrusion) 1, Gun Combat (Slug) 1, Stealth 1"
		case 8:
			npc.skills = "Investigate 2, Science (Forensics) 2, Science (Anatomy) 2, Medic (Diagnosis) 2, Art (Holography) 1, Recon 1, Science (Psychology) 1, Deception 0"
		case 9:
			npc.skills = "Advocate (Legal) 2, Carouse 2, Interrogation (Questioning) 2, Investigate 2, Recon 2, Deception (Intrusion) 1, Gun Combat (Slug) 1, Streetwise 1, Science 0"
		case 10:
			npc.skills = "Carouse 2, Investigate 2, Interrogation (Questioning) 2, Streetwise 2, Recon 2, Stealth 2, Broker 1, Melee (Unarmed Combat) 1, Gun Combat (Slug) 1, Science (Forensics) 1, Persuade 1"
		case 11:
			npc.skills = "Carouse 3, Advocate (Legal) 2, Investigate 2, Interrogation (Questioning) 2, Recon 2, Stealth 2, Deception (Intrusion) 1, Deception 1, Deception (Disguise) 1, Science (Forensics) 1"
		case 12:
			npc.skills = "Carouse 3, Investigate 3, Streetwise 3, Gun Combat (Slug) 2, Interrogation (Questioning) 2, Recon 2, Stealth 2, Melee (Unarmed Combat) 1, Science (Forensics) 1, Deception 0"
		}
	case occJournalist:
		switch npc.skillList {
		case 2:
			npc.skills = "Art (Holography) 1, Art (Writing) 1, Electronics (Computers) 1, Investigate 1, Language (Any) 1, Persuade 1, Survival 0"
		case 3:
			npc.skills = "Art (Writing) 1, Carouse 1, Electronics (Computers) 1, Investigate 1, Persuade 1, Streetwise 1"
		case 4:
			npc.skills = "Art (Writing) 1, Deception (Intrusion) 1, Carouse 1, Investigate 1, Recon 1, Advocate 0, Jack of All Trades 0"
		case 5:
			npc.skills = "Admin 1, Art (Holography) 1, Broker 1, Diplomat 1, Electronics (Computers) 1, Advocate 0"
		case 6:
			npc.skills = "Art (Writing) 2, Investigate 2, Persuade 2, Art (Holography) 1, Electronics (Computers) 1, Language (Any) 1, Deception 0, Survival 0"
		case 7:
			npc.skills = "Carouse 2, Investigate 2, Persuade 2, Streetwise 2, Advocate (Oratory) 1, Art (Writing) 1, Deception (Intrusion) 1, Deception (Disguise) 1, Electronics (Computers) 1, Recon 1"
		case 8:
			npc.skills = "Carouse 2, Deception (Disguise) 2, Investigate 2, Streetwise 2, Art (Writing) 1, Recon 1, Jack of All Trades 1"
		case 9:
			npc.skills = "Admin 2, Art (Holography) 2, Broker 2, Diplomat 2, Persuade 2, Language (Any) 1, Language (Any) 1, Steward 1"
		case 10:
			npc.skills = "Advocate (Oratory) 2, Art (Writing) 2, Persuade 2, Language (Any) 2, Language (Any) 1, Language (Any) 1, Advocate (Politics) 1, Diplomat 1"
		case 11:
			npc.skills = "Carouse 3, Investigate 3, Persuade 3, Advocate (Oratory) 2, Art (Writing) 2, Deception 2, Deception (Disguise) 2, Deception (Intrusion) 2, Recon 1, Stealth 1"
		case 12:
			npc.skills = "Investigate 3, Streetwise 3, Persuade 3, Advocate (Oratory) 2, Art (Writing) 2, Advocate (Politics) 2, Diplomat 2, Leadership 1, Broker 1"
		}
	case occMedic:
		switch npc.skillList {
		case 2:
			npc.skills = "Admin 1, Investigate 1, Medic (Diagnosis) 1, Medic (First Aid) 1, Persuade 1, Science (Anatomy) 1, Electronics 0"
		case 3:
			npc.skills = "Admin 1, Broker 1, Carouse 1, Diplomat 1, Medic (Diagnosis) 1, Medic (First Aid) 1, Persuade 1, Science 0, Steward 0"
		case 4:
			npc.skills = "Broker 1, Carouse 1, Electronics (Computers) 1, Electronics (Sensors) 1, Medic (Diagnosis) 1, Medic (First Aid) 1, Science 0, Vacc Suit 0"
		case 5:
			npc.skills = "Flyer (Grav) 1, Medic (Cryogenics) 1, Medic (Diagnosis) 1, Medic (First Aid) 1, Streetwise 1, Gun Combat 0, Science 0, Electronics 0"
		case 6:
			npc.skills = "Medic (Diagnosis) 2, Medic (First Aid) 2, Admin 1, Broker 1, Investigate 1, Science (Biology) 1, Persuade 1"
		case 7:
			npc.skills = "Medic (Diagnosis) 2, Medic (Surgery) 2, Admin 1, Broker 1, Carouse 1, Diplomat 1, Medic (Nutrition) 1, Medic (First Aid) 1, Science (Biology) 1, Science (Anatomy) 1"
		case 8:
			npc.skills = "Electronics (Sensors) 2, Medic (Cryogenics) 2, Medic (Diagnosis) 2, Medic (First Aid) 2, Broker 2, Carouse 1, Electronics (Computers) 1, Medic (Surgery) 1, Medic (Nutrition) 1, Science 0, Vacc Suit 0"
		case 9:
			npc.skills = "Medic (Diagnosis) 2, Medic (First Aid) 2, Drive (Wheeled) 2, Flyer (Grav) 2, Medic (Cryogenics) 1, Streetwise 1, Gun Combat (Slug) 1, Science 0, Electronics 0, Pilot 0"
		case 10:
			npc.skills = "Medic (Diagnosis) 3, Admin 2, Broker 2, Carouse 2, Diplomat 2, Medic (First Aid) 2, Medic (Nutrition) 2, Medic (Surgery) 2, Science (Biology) 1, Science (Anatomy) 1, Electronics 0"
		case 11:
			npc.skills = "Electronics (Sensors) 3, Medic (Cryogenics) 3, Medic (Diagnosis) 2, Medic (First Aid) 2, Broker 2, Carouse 1, Electronics (Computers) 1, Medic (Surgery) 1, Medic (Nutrition) 1, Science 0, Vacc Suit 0, Zero-G 0"
		case 12:
			npc.skills = "Medic (Diagnosis) 3, Medic (First Aid) 3, Drive (Wheeled) 2, Flyer (Grav) 2, Medic (Cryogenics) 1, Pilot (Small Craft) 1, Streetwise 1, Gun Combat (Slug) 1, Science 0, Electronics 0"
		}
	case occNavy:
		switch npc.skillList {
		case 2:
			npc.skills = "Electronics (Sensors) 1, Gun Combat (Laser) 1, Mechanic 1, Gunner 0, Pilot 0, Vacc Suit 0"
		case 3:
			npc.skills = "Astrogation 1, Carouse 1, Electronics (Sensors) 1, Pilot (Spacecraft) 1, Vacc Suit 0, Zero-G 0"
		case 4:
			npc.skills = "Carouse 1, Engineer (M-Drive) 1, Engineer (Z-Drive) 1, Mechanic 1, Vacc Suit 0, Zero-G 0"
		case 5:
			npc.skills = "Gunner (Turrets) 1, Gun Combat (Energy) 1, Electronics (Sensors) 1, Tactics (Naval) 1, Persuade 1, Vacc Suit 0, Zero-G 0"
		case 6:
			npc.skills = "Admin 1, Astrogation 1, Carouse 1, Persuade 1, Pilot 0, Vacc Suit 0, Zero-G 0"
		case 7:
			npc.skills = "Pilot (Spacecraft) 2, Astrogation 1, Electronics (Sensors) 1, Mechanic 1, Vacc Suit 0, Zero-G 0"
		case 8:
			npc.skills = "Carouse 2, Engineer (M-Drive) 2, Engineer (Z-Drive) 2, Electronics (Computers) 1, Engineer (Life Support) 1, Mechanic 1, Vacc Suit 0, Zero-G 0"
		case 9:
			npc.skills = "Gunner (Turrets) 2, Gun Combat (Energy) 2, Carouse 1, Electronics (Sensors) 1, Tactics (Naval) 1, Vacc Suit 0, Zero-G 0"
		case 10:
			npc.skills = "Engineer (M-Drive) 3, Engineer (Z-Drive) 3, Carouse 2, Mechanic 2, Electronics (Computers) 2, Engineer (Life Support) 1, Engineer (Power) 1, Persuade 1, Diplomat 0, Vacc Suit 0, Zero-G 0"
		case 11:
			npc.skills = "Gunner (Turrets) 3, Carouse 2, Electronics (Sensors) 2, Gun Combat (Energy) 2 Tactics (Naval) 2, Persuade 1, Mechanic 1, Vacc Suit 0, Zero-G 0"
		case 12:
			npc.skills = "Admin 2, Astrogation 2, Pilot (Spacecraft) 2, Electronics (Sensors) 2, Persuade 2, Diplomat 1, Language (Any) 1, Language (Any) 1, Advocate (Legal) 1, Engineer 0, Tactics 0, Vacc Suit 0, Zero-G 0"
		}
	case occOrganizedCrime:
		switch npc.skillList {
		case 2:
			npc.skills = "Deception (Intrusion) 1, Flyer (Grav) 1, Gun Combat (Slug) 1, Melee (Unarmed Combat) 1, Persuade 1, Streetwise 1"
		case 3:
			npc.skills = "Deception (Intrusion) 1, Gun Combat (Slug) 1, Melee (Unarmed Combat) 1, Streetwise 1, Stealth 1, Tactics (Military) 1"
		case 4:
			npc.skills = "Broker 1, Gun Combat (Slug) 1, Melee (Unarmed Combat) 1, Persuade 1, Recon 1, Streetwise 1"
		case 5:
			npc.skills = "Admin 1, Broker 1, Gun Combat (Slug) 1, Melee (Unarmed Combat) 1, Persuade 1, Streetwise 1"
		case 6:
			npc.skills = "Deception (Intrusion) 2, Gun Combat (Slug) 2, Melee (Unarmed Combat) 2, Persuade 2, Streetwise 2, Flyer (Grav) 1, Diplomat 1"
		case 7:
			npc.skills = "Gun Combat (Slug) 2, Melee (Unarmed Combat) 2, Streetwise 2, Stealth 2, Tactics (Military) 2, Deception (Intrusion) 1, Persuade 1, Investigate 1"
		case 8:
			npc.skills = "Broker 2, Gun Combat (Slug) 2, Melee (Unarmed Combat) 2, Persuade 2, Recon 2, Streetwise 2, Admin 1, Diplomat 1"
		case 9:
			npc.skills = "Admin 2, Broker 2, Gun Combat (Slug) 2, Melee (Unarmed Combat) 2, Persuade 2, Streetwise 2, Recon 1, Diplomat 1"
		case 10:
			npc.skills = "Gun Combat (Slug) 3, Streetwise 3, Melee (Unarmed Combat) 2, Interrogation (Torture) 2, Streetwise 2, Stealth 2, Tactics (Military) 2, Deception (Intrusion) 1, Persuade 1, Investigate 1"
		case 11:
			npc.skills = "Broker 3, Persuade 3, Gun Combat (Slug) 2, Melee (Unarmed Combat) 2, Recon 2, Streetwise 2, Interrogation (Torture) 2, Admin 1, Diplomat 1, Science (Forensics) 1"
		case 12:
			npc.skills = "Admin 3, Broker 3, Gun Combat (Slug) 2, Melee (Unarmed Combat) 2, Persuade 2, Streetwise 2, Recon 2, Diplomat 1"
		}
	case occPirate:
		switch npc.skillList {
		case 2:
			npc.skills = "Deception (Intrusion) 1, Electronics (Sensors) 1, Gun Combat (Energy) 1, Gun Combat (Slug) 1, Streetwise 1, Survival 0, Vacc Suit 0, Zero-G 0"
		case 3:
			npc.skills = "Admin 1, Astrogation 1, Broker 1, Electronics (Comms) 1, Electronics (Sensors) 1, Mechanic 1, Pilot (Spacecraft) 1, Vacc Suit 0, Zero-G 0"
		case 4:
			npc.skills = "Explosives 1, Flyer (Grav) 1, Gun Combat (Slug) 1, Melee (Blade) 1, Recon 1, Tactics (Military) 1, Vacc Suit 0, Zero-G 0"
		case 5:
			npc.skills = "Electronics (Computers) 1, Engineer (M Drive) 1, Engineer (Z Drive) 1, Mechanic 1, Science 0, Vacc Suit 0, Zero-G 0"
		case 6:
			npc.skills = "Electronics (Sensors) 1, Gunner (Turrets) 1, Gun Combat (Energy) 1, Tactics (Naval) 1, Vacc Suit 0, Zero-G 0"
		case 7:
			npc.skills = "Deception (Intrusion) 2, Electronics (Sensors) 2, Gun Combat (Slug) 2, Broker 1, Carouse 1, Explosives 1, Streetwise 1, Vacc Suit 0, Zero-G 0"
		case 8:
			npc.skills = "Admin 2, Broker 2, Astrogation 1, Electronics (Sensors) 1, Mechanic 1, Pilot (Spacecraft) 1, Vacc Suit 0, Zero-G 0"
		case 9:
			npc.skills = "Explosives 2, Gun Combat (Slug) 2, Melee (Blade) 2, Recon 2, Tactics (Military) 2, Carouse 1, Persuade 1, Pilot (Small Craft) 1, Zero-G 1, Medic 0, Vacc Suit 0"
		case 10:
			npc.skills = "Electronics (Sensors) 2, Gunner (Turrets) 2, Tactics (Naval) 2, Carouse 1, Gun Combat (Energy) 1, Mechanic 1, Vacc Suit 0, Zero-G 0"
		case 11:
			npc.skills = "Engineer (M Drive) 2, Engineer (Z Drive) 2, Mechanic 2, Broker 1, Carouse 1, Electronics (Computers) 1, Vacc Suit 0, Zero-G 0"
		case 12:
			npc.skills = "Deception (Intrusion) 3, Broker 2, Carouse 2, Electronics (Sensors) 2, Explosives 2, Gun Combat (Slug) 2, Melee (Blade) 2, Tactics (Military) 2, Persuade 2, Zero-G 2, Mechanic 1, Vacc Suit 0"
		}
	case occPolice:
		switch npc.skillList {
		case 2:
			npc.skills = "Advocate (Legal) 1, Carouse 1, Deception 1, Interrogation (Questioning) 1, Investigate 1, Recon 1"
		case 3:
			npc.skills = "Animals (Riding) 1, Athletics (Endurance) 1, Carouse 1, Gun Combat (Slug) 1, Melee (Bludgeon) 1, Persuade 1"
		case 4:
			npc.skills = "Carouse 1, Flyer (Grav) 1, Gun Combat (Slug) 1, Melee (Bludgeon) 1, Mechanic 1, Persuade 1"
		case 5:
			npc.skills = "Drive (Tracked) 1, Explosives 1, Gun Combat (Slug) 1, Heavy Weapons (Vehicle Mounted) 1, Melee (Bludgeon) 1, Tactics (Military) 1"
		case 6:
			npc.skills = "Carouse 2, Deception 2, Interrogation (Questioning) 2, Investigate 2, Advocate (Legal) 1, Deception (Disguise) 1, Persuade 1, Recon 1, Stealth 1"
		case 7:
			npc.skills = "Carouse 2, Persuade 2, Animals (Riding) 1, Athletics (Endurance) 1, Gun Combat (Slug) 1, Melee (Bludgeon) 1, Recon 1, Streetwise 1"
		case 8:
			npc.skills = "Carouse 2, Flyer (Grav) 2, Persuade 2, Gun Combat (Slug) 1, Melee (Bludgeon) 1, Mechanic 1, Streetwise 1"
		case 9:
			npc.skills = "Explosives 2, Gun Combat (Slug) 2, Melee (Bludgeon) 2, Tactics (Military) 2, Drive (Tracked) 1, Heavy Weapons (Vehicle Mounted) 1, Persuade 1, Streetwise 1"
		case 10:
			npc.skills = "Carouse 3, Persuade 3, Interrogation (Questioning) 3, Advocate (Legal) 2, Deception 2, Investigate 2, Recon 1, Streetwise 1"
		case 11:
			npc.skills = "Carouse 3, Persuade 3, Gun Combat (Slug) 2, Melee (Bludgeon) 2, Recon 2, Streetwise 2, Animals (Riding) 1, Athletics (Endurance) 1"
		case 12:
			npc.skills = "Carouse 3, Persuade 3, Flyer (Grav) 2, Gun Combat (Slug) 2, Melee (Bludgeon) 2, Streetwise 1, Diplomat 1, Mechanic 1"
		}
	case occPolitican:
		switch npc.skillList {
		case 2:
			npc.skills = "Admin 1, Advocate (Oratory) 1, Advocate (Politics) 1, Carouse 1, Deception 1, Persuade 1, Electronics 0"
		case 3:
			npc.skills = "Advocate (Legal) 1, Advocate (Oratory) 1, Advocate (Politics) 1, Deception 1, Persuade 1, Leadership 1"
		case 4:
			npc.skills = "Admin 1, Advocate (Politics) 1, Art (Writing) 1, Broker 1, Electronics (Comms) 1, Leadership 1"
		case 5:
			npc.skills = "Advocate (Politics) 1, Broker 1, Carouse 1, Deception 1, Diplomat 1, Streetwise 1, Steward 1"
		case 6:
			npc.skills = "Advocate (Oratory) 2, Advocate (Politics) 2, Carouse 2, Persuade 2, Diplomat 1, Deception 1, Advocate (Legal) 1, Leadership 1"
		case 7:
			npc.skills = "Advocate (Legal) 2, Advocate (Oratory) 2, Advocate (Politics) 2, Deception 1, Diplomat 1, Persuade 1, Leadership 1, Science (Psychology) 1"
		case 8:
			npc.skills = "Admin 2, Advocate (Politics) 2, Art (Writing) 2, Broker 2, Carouse 1, Diplomat 1, Electronics (Comms) 1, Leadership 1, Persuade 1, Streetwise 1"
		case 9:
			npc.skills = "Advocate (Politics) 2, Broker 2, Carouse 2, Deception 2, Diplomat 2, Advocate (Oratory) 1, Persuade 1, Streetwise 1, Steward 1"
		case 10:
			npc.skills = "Advocate (Oratory) 3, Advocate (Politics) 3, Carouse 3, Persuade 3, Diplomat 2, Deception 2, Advocate (Legal) 2, Leadership 1"
		case 11:
			npc.skills = "Advocate (Legal) 3, Advocate (Oratory) 3, Advocate (Politics) 3, Deception 2, Diplomat 2, Persuade 2, Language (Any) 1, Leadership 1, Science (Psychology) 1"
		case 12:
			npc.skills = "Admin 3, Advocate (Politics) 3, Art (Writing) 3, Broker 3, Carouse 2, Diplomat 2, Electronics (Comms) 2, Persuade 2, Leadership 1, Streetwise 1"
		}
	case occPortAuthority:
		switch npc.skillList {
		case 2:
			npc.skills = "Advocate (Legal) 1, Electronics (Sensors) 1, Persuade 1, Recon 1, Streetwise 1"
		case 3:
			npc.skills = "Admin 1, Astrogation 1, Electronics (Computers) 1, Electronics (Sensors) 1, Pilot (Spacecraft) 1, Persuade 1"
		case 4:
			npc.skills = "Advocate (Legal) 1, Electronics (Sensors) 1, Investigate 1, Pilot (Small Craft) 1, Recon 1, Streetwise 1, Vacc Suit 0, Zero-G 0"
		case 5:
			npc.skills = "Carouse 1, Electronics (Sensors) 1, Flyer (Grav) 1, Mechanic 1, Streetwise 1, Vacc Suit 0"
		case 6:
			npc.skills = "Electronics (Sensors) 2, Persuade 2, Recon 2, Carouse 1, Advocate (Legal) 1, Streetwise 1"
		case 7:
			npc.skills = "Astrogation 2, Electronics (Computers) 2, Electronics (Sensors) 2, Admin 1, Carouse 1, Diplomat 1, Pilot (Spacecraft) 1, Persuade 1"
		case 8:
			npc.skills = "Advocate (Legal) 2, Electronics (Sensors) 2, Investigate 2, Recon 2, Streetwise 2, Broker 1, Carouse 1, Persuade 1, Pilot (Small Craft) 1, Vacc Suit 0, Zero-G 0"
		case 9:
			npc.skills = "Carouse 2, Streetwise 2, Admin 1, Broker 1, Electronics (Sensors) 1, Flyer (Grav) 1, Mechanic 1, Persuade 1, Vacc Suit 0"
		case 10:
			npc.skills = "Admin 2, Astrogation 2, Carouse 2, Diplomat 2, Electronics (Computers) 2, Electronics (Sensors) 2, Persuade 2, Broker 1, Pilot (Spacecraft) 1, Streetwise 1"
		case 11:
			npc.skills = "Investigate 3, Recon 3, Streetwise 3, Advocate (Legal) 2, Electronics (Sensors) 2, Broker 2, Carouse 2, Persuade 2, Diplomat 1, Pilot (Small Craft) 1, Vacc Suit 0, Zero-G 0"
		case 12:
			npc.skills = "Carouse 3, Streetwise 3, Broker 2, Electronics (Sensors) 2, Persuade 2, Admin 1, Diplomat 1, Jack of All Trades 1, Flyer (Grav) 1, Mechanic 1, Vacc Suit 0"
		}
	case occProstitute:
		switch npc.skillList {
		case 2:
			npc.skills = "Broker 1, Carouse 1, Deception 1, Language (Any) 1, Persuade 1, Recon 1, Streetwise 1 "
		case 3:
			npc.skills = "Admin 1, Advocate (Legal) 1, Art (Dance) 1, Broker 1, Carouse 1, Persuade 1, Streetwise 1"
		case 4:
			npc.skills = "Art (Dance) 1, Art (Painting) 1, Broker 1, Carouse 1, Diplomat 1, Science (Psychology) 1, Steward 1"
			npc.changeStat(6, 2)
		case 5:
			npc.skills = "Broker 1, Carouse 1, Gun Combat (Slug) 1, Melee (Unarmed) 1, Persuade 1, Streetwise 1, Stealth 0"
			npc.changeStat(6, -2)
		case 6:
			npc.skills = "Broker 2, Carouse 2, Deception 2, Persuade 2, Recon 2, Streetwise 2, Diplomat 1, Language (Any) 1, Language (Any) 1"
		case 7:
			npc.skills = "Admin 2, Art (Dance) 2, Broker 2, Carouse 2, Persuade 2, Streetwise 2, Advocate (Legal) 1, Language (Any) 1, Diplomat 1"
		case 8:
			npc.skills = "Broker 2, Carouse 2, Persuade 2, Streetwise 2, Gun Combat (Slug) 1, Melee (Unarmed) 1, Stealth 0"
			npc.changeStat(6, -2)
		case 9:
			npc.skills = "Art (Dance) 2, Art (Painting) 2, Broker 2, Carouse 2, Diplomat 2, Steward 2, Jack of All Trades 1, Science (Psychology) 1"
			npc.changeStat(6, 2)
		case 10:
			npc.skills = "Admin 3, Art (Dance) 3, Broker 3, Carouse 3, Persuade 3, Streetwise 3, Diplomat 2, Advocate (Legal) 1, Language (Any) 1, Jack of All Trades 1"
		case 11:
			npc.skills = "Broker 3, Carouse 3, Persuade 3, Streetwise 3, Gun Combat (Slug) 2, Diplomat 1, Jack of All Trades 1, Melee (Unarmed) 1, Stealth 0"
			npc.changeStat(6, -2)
		case 12:
			npc.skills = "Art (Dance) 3, Art (Painting) 3, Broker 3, Carouse 3, Diplomat 2, Steward 2, Jack of All Trades 1, Science (Psychology) 1"
			npc.changeStat(6, 3)
		}
	case occScavenger:
		switch npc.skillList {
		case 2:
			npc.skills = "Broker 1, Carouse 1, Language (Any) 1, Mechanic 1, Persuade 1"
		case 3:
			npc.skills = "Admin 1, Broker 1, Deception (Intrusion) 1, Mechanic 1, Persuade 1, Streetwise 1"
		case 4:
			npc.skills = "Admin 1, Advocate (Legal) 1, Broker 1, Carouse 1, Deception 1, Persuade 1"
		case 5:
			npc.skills = "Admin 1, Art (Painting) 1, Broker 1, Electronics (Computers) 1, Mechanic 1, Science (History) 1"
		case 6:
			npc.skills = "Broker 2, Carouse 2, Mechanic 2, Persuade 2, Investigate 1, Language (Any) 1, Recon 1, Streetwise 1"
		case 7:
			npc.skills = "Admin 2, Broker 2, Mechanic 2, Persuade 2, Carouse 1, Deception (Intrusion) 1, Streetwise 1"
		case 8:
			npc.skills = "Broker 2, Carouse 2, Deception 2, Persuade 2, Admin 1, Advocate (Legal) 1, Science (History) 1"
		case 9:
			npc.skills = "Art (Painting) 2, Broker 2, Electronics (Computers) 2, Mechanic 2, Admin 1, Science (History) 1"
		case 10:
			npc.skills = "Broker 3, Carouse 3, Mechanic 3, Persuade 3, Investigate 2, Language (Any) 1, Recon 1, Streetwise 1, Jack of All Trades 1"
		case 11:
			npc.skills = "Admin 3, Broker 3, Mechanic 3, Carouse 2, Persuade 2, Diplomat 1, Deception (Intrusion) 1, Jack of All Trades 1, Streetwise 1"
		case 12:
			npc.skills = "Broker 3, Carouse 3, Persuade 3, Deception 2, Admin 1, Advocate (Legal) 1, Diplomat 1, Jack of All Trades 1, Science (History) 1"
		}
	case occSports:
		switch npc.skillList {
		case 2:
			npc.skills = "Athletics (Any) 1, Carouse 1, Melee (Unarmed) 1, Persuade 1, Tactics (Sport) 1, Leadership 0"
		case 3:
			npc.skills = "Athletics (Coordination) 1, Athletics (Endurance) 1, Athletics (Strength) 1, Carouse 1, Tactics (Sport) 1, Leadership 0"
		case 4:
			npc.skills = "Admin 1, Advocate (Oratory) 1, Athletics (Endurance) 1, Broker 1, Diplomat 1, Leadership 1, Persuade 1, Tactics (Sport) 1"
		case 5:
			npc.skills = "Athletics (Endurance) 2, Carouse 2, Tactics (Sport) 2, Athletics (Strength) 1, Broker 1, Persuade 1"
		case 6:
			npc.skills = "Athletics (Coordination) 2, Athletics (Strength) 2, Broker 1, Carouse 1, Diplomat 1, Tactics (Sport) 1, Leadership 0"
		case 7:
			npc.skills = "Athletics (Coordination) 2, Athletics (Endurance) 2, Athletics (Strength) 2, Carouse 2, Tactics (Sport) 2, Broker 1, Gambler 1, Leadership 1, Streetwise 0"
		case 8:
			npc.skills = "Admin 2, Advocate (Oratory) 2, Broker 2, Diplomat 2, Leadership 2, Persuade 2, Tactics (Sport) 2, Athletics (Endurance) 1, Science (Psychology) 1"
		case 9:
			npc.skills = "Athletics (Endurance) 2, Carouse 2, Tactics (Sport) 2, Athletics (Strength) 2, Broker 1, Diplomat 1, Art (Writing) 1, Persuade 1"
		case 10:
			npc.skills = "Athletics (Coordination) 2, Athletics (Endurance) 2, Athletics (Strength) 2, Carouse 2, Tactics (Sport) 2, Diplomat 1, Leadership 1, Persuade 1"
		case 11:
			npc.skills = "Athletics (Endurance) 3, Carouse 3, Tactics (Sport) 2, Athletics (Strength) 2, Broker 2, Diplomat 1, Leadership 1, Persuade 1"
		case 12:
			npc.skills = "Athletics (Coordination) 3, Athletics (Endurance) 3, Athletics (Strength) 3, Broker 2, Carouse 2, Tactics (Sport) 2, Diplomat 1, Leadership 1, Persuade 1"
		}
	case occSpy:
		switch npc.skillList {
		case 2:
			npc.skills = "Deception 1, Electronics (Computers) 1, Gun Combat (Slug) 1, Investigate 1, Language (Any) 1, Language (Any) 1, Melee (Unarmed Combat) 1"
		case 3:
			npc.skills = "Deception (Disguise) 1, Flyer (Any) 1, Gun Combat (Slug) 1, Investigate 1, Language (Any) 1, Streetwise 1, Stealth 1"
		case 4:
			npc.skills = "Carouse 1, Diplomat 1, Electronics (Sensors) 1, Investigate 1, Language (Any) 1, Persuade 1, Recon 1, Stealth 1, Survival (Any) 1"
		case 5:
			npc.skills = "Broker 1, Deception (Disguise) 1, Deception 1, Investigate 1, Jack of All Trades 1, Trade (Any) 1, Gun Combat 0, Melee 0"
		case 6:
			npc.skills = "Deception (Disguise) 2, Investigate 2, Flyer (Any) 1, Gun Combat (Slug) 1, Language (Any) 1, Streetwise 1, Stealth 1"
		case 7:
			npc.skills = "Deception 2, Electronics (Computers) 2, Gun Combat (Slug) 2, Investigate 2, Language (Any) 2, Language (Any) 1, Language (Any) 1, Melee (Unarmed Combat) 1, Persuade 1"
		case 8:
			npc.skills = "Carouse 2, Diplomat 2, Electronics (Sensors) 2, Investigate 2, Language (Any) 1, Persuade 1, Recon 1, Stealth 1, Survival (Any) 1"
		case 9:
			npc.skills = "Broker 2, Deception (Disguise) 2, Deception 2, Investigate 2, Jack of All Trades 1, Trade (Any) 1, Gun Combat 0, Melee 0"
		case 10:
			npc.skills = "Deception (Disguise) 3, Investigate 3, Flyer (Any) 2, Gun Combat (Slug) 2, Language (Any)  2, Investigate 1, Recon 1, Streetwise 1, Stealth 1"
		case 11:
			npc.skills = "Deception 3, Electronics (Computers) 3, Investigate 3, Language (Any) 3, Gun Combat (Slug) 2, Language (Any) 2, Language (Any) 2, Language (Any) 1, Melee (Unarmed Combat) 1, Persuade 1, Jack of All Trades 1"
		case 12:
			npc.skills = "Carouse 3, Deception (Disguise) 3, Deception 3, Investigate 3, Broker 2, Diplomat 2, Jack of All Trades 1, Trade (Any) 1, Gun Combat 0, Melee 0"
		}
	case occThief:
		switch npc.skillList {
		case 2:
			npc.skills = "Deception (Intrusion) 1, Electronics (Computers) 1, Melee (Unarmed Combat) 1, Persuade 1, 	Streetwise 1, Gun Combat 0"
		case 3:
			npc.skills = "Broker 1, Carouse 1, Deception (Disguise) 1, Gambler 1, Persuade 1, Steward 1"
			npc.changeStat(6, 3)
		case 4:
			npc.skills = "Electronics (Computers) 2, Broker 1, Carouse 1, Mechanic 1, Persuade 1, Advocate 0"
		case 5:
			npc.skills = "Deception (Intrusion) 1, Electronics (Sensors) 1, Gun Combat (Slug) 1, Recon 1, Stealth 1"
		case 6:
			npc.skills = "Deception (Intrusion) 2, Electronics (Computers) 2, Streetwise 2, Melee (Unarmed Combat) 1, Persuade 1, Stealth 1, Gun Combat 0, Investigate 0"
		case 7:
			npc.skills = "Deception (Intrusion) 2, Electronics (Sensors) 2, Recon 2, Stealth 2, Gun Combat (Slug) 1, Investigate 1, Streetwise 1"
		case 8:
			npc.skills = "Broker 2, Carouse 2, Deception (Disguise) 2, Investigate 2, Persuade 2, Deception (Intrusion) 1, Gambler 1, Steward 1"
			npc.changeStat(6, 3)
		case 9:
			npc.skills = "Electronics (Computers) 2, Broker 2, Carouse 2, Advocate (Politics) 1, Mechanic 1, Persuade 1"
		case 10:
			npc.skills = "Deception (Intrusion) 3, Electronics (Sensors) 3, Recon 2, Investigate 2, Streetwise 2, Stealth 2, Gun Combat (Slug) 1"
		case 11:
			npc.skills = "Electronics (Computers) 3, Broker 3, Carouse 2, Advocate (Politics) 1, Mechanic 1, Persuade 1, Streetwise 1"
		case 12:
			npc.skills = "Carouse 3, Persuade 3, Art (Acting) 2, Broker 2, Diplomat 2, Deception (Disguise) 2, Investigate 2, Deception (Intrusion) 1, Gambler 1, Steward 1"
			npc.changeStat(6, 3)
		}
	case occVagabond:
		switch npc.skillList {
		case 2:
			npc.skills = "Melee (Unarmed Combat) 1, Persuade 1, Streetwise 1, Survival (Any) 1, Stealth 0"
		case 3:
			npc.skills = "Navigation 1, Persuade 1, Streetwise 1, Survival (Any) 1, Gun Combat (Slug) 1, Deception 0, Stealth 0"
		case 4:
			npc.skills = "Carouse 1, Deception 1, Melee (Unarmed Combat) 1, Mechanic 1, Survival (Any) 1, Vacc Suit 0"
		case 5:
			npc.skills = "Carouse 1, Language (Any) 1, Melee (Unarmed Combat) 1, Streetwise 1, Survival (Any) 1"
		case 6:
			npc.skills = "Persuade 2, Streetwise 2, Survival (Any) 2, Melee (Unarmed Combat) 1, Stealth 1, Deception 0"
		case 7:
			npc.skills = "Persuade 2, Streetwise 2, Navigation 1, Survival (Any) 1, Gun Combat (Slug) 1, Deception 0, Stealth 0"
		case 8:
			npc.skills = "Carouse 2, Deception 2, Survival (Any) 2, Melee (Unarmed Combat) 1, Mechanic 1, Survival (Any) 1, Vacc Suit 0"
		case 9:
			npc.skills = "Carouse 2, Streetwise 2, Survival (Any) 2, Language (Any) 1, Melee (Unarmed Combat) 1, Persuade 1"
		case 10:
			npc.skills = "Persuade 2, Streetwise 2, Survival (Any) 2, Jack of All Trades 1, Melee (Unarmed Combat) 1, Stealth 1, Deception 0"
		case 11:
			npc.skills = "Persuade 2, Streetwise 2, Navigation 2, Survival (Any) 2, Jack of All Trades 1, Gun Combat (Slug) 1, Deception 0, Stealth 0"
		case 12:
			npc.skills = "Carouse 3, Deception 2, Survival (Any) 2, Jack of All Trades 1, Melee (Unarmed Combat) 1, Mechanic 1, Survival (Any) 1, Vacc Suit 0"
		}
	}
}

func firstName() string {
	names := []string{
		"Aamir",
		"Ayub",
		"Binyamin",
		"Efraim",
		"Ibrahim",
		"Ilyas",
		"Ismail",
		"Jibril",
		"Jumanah",
		"Kazi",
		"Lut",
		"Matta",
		"Mohammed",
		"Mubarak",
		"Mustafa",
		"Nazir",
		"Rahim",
		"Reza",
		"Sharif",
		"Taimur",
		"Usman",
		"Yakub",
		"Yusuf",
		"Zakariya",
		"Zubair",
		"Aisha",
		"Alimah",
		"Badia",
		"Bisharah",
		"Chanda",
		"Daliya",
		"Fatimah",
		"Ghania",
		"Halah",
		"Kaylah",
		"Khayrah",
		"Layla",
		"Mina",
		"Munisa",
		"Mysha",
		"Naimah",
		"Nissa",
		"Nura",
		"Parveen",
		"Rana",
		"Shalha",
		"Suhira",
		"Tahirah",
		"Yasmin",
		"Zulehka",
		"Adan",
		"Ahsa",
		"Andalus",
		"Asmara",
		"Asqlan",
		"Baqubah",
		"Basit",
		"Baysan",
		"Baytlahm",
		"Bursaid",
		"Dahilah",
		"Darasalam",
		"Dawhah",
		"Ganin",
		"Gebal",
		"Gibuti",
		"Giddah",
		"Harmah",
		"Hartum",
		"Hibah",
		"Hims",
		"Hubar",
		"Karbala",
		"Kut",
		"Lacant",
		"Magrit",
		"Masqat",
		"Misr",
		"Muruni",
		"Qabis",
		"Qina",
		"Rabat",
		"Ramlah",
		"Riyadh",
		"Sabtah",
		"Salalah",
		"Sana",
		"Sinqit",
		"Suqutrah",
		"Sur",
		"Tabuk",
		"Tangah",
		"Tarifah",
		"Tarrakunah",
		"Tisit",
		"Uman",
		"Urdunn",
		"Wasqah",
		"Yaburah",
		"Yaman",
		"Aiguo",
		"Bohai",
		"Chao",
		"Dai",
		"Dawei",
		"Duyi",
		"Fa",
		"Fu",
		"Gui",
		"Hong",
		"Jianyu",
		"Kang",
		"Li",
		"Niu",
		"Peng",
		"Quan",
		"Ru",
		"Shen",
		"Shi",
		"Song",
		"Tao",
		"Xue",
		"Yi",
		"Yuan",
		"Zian",
		"Biyu",
		"Changying",
		"Daiyu",
		"Huidai",
		"Huiliang",
		"Jia",
		"Jingfei",
		"Lan",
		"Liling",
		"Liu",
		"Meili",
		"Niu",
		"Peizhi",
		"Qiao",
		"Qing",
		"Ruolan",
		"Shu",
		"Suyin",
		"Ting",
		"Xia",
		"Xiaowen",
		"Xiulan",
		"Ya",
		"Ying",
		"Zhilan",
		"Andong",
		"Anqing",
		"Anshan",
		"Chaoyang",
		"Chaozhou",
		"Chifeng",
		"Dalian",
		"Dunhuang",
		"Fengjia",
		"Fengtian",
		"Fuliang",
		"Fushun",
		"Gansu",
		"Ganzhou",
		"Guizhou",
		"Hotan",
		"Hunan",
		"Jinan",
		"Jingdezhen",
		"Jinxi",
		"Jinzhou",
		"Kunming",
		"Liaoning",
		"Linyi",
		"Lushun",
		"Luzhou",
		"Ningxia",
		"Pingxiang",
		"Pizhou",
		"Qidong",
		"Qingdao",
		"Qinghai",
		"Rehe",
		"Shanxi",
		"Taiyuan",
		"Tengzhou",
		"Urumqi",
		"Weifang",
		"Wugang",
		"Wuxi",
		"Xiamen",
		"Xian",
		"Xikang",
		"Xining",
		"Xinjiang",
		"Yidu",
		"Yingkou",
		"Yuxi",
		"Zigong",
		"Zoige",
		"Adam",
		"Albert",
		"Alfred",
		"Allan",
		"Archibald",
		"Arthur",
		"Basil",
		"Charles",
		"Colin",
		"Donald",
		"Douglas",
		"Edgar",
		"Edmund",
		"Edward",
		"George",
		"Harold",
		"Henry",
		"Ian",
		"James",
		"John",
		"Lewis",
		"Oliver",
		"Philip",
		"Richard",
		"William",
		"Abigail",
		"Anne",
		"Beatrice",
		"Blanche",
		"Catherine",
		"Charlotte",
		"Claire",
		"Eleanor",
		"Elizabeth",
		"Emily",
		"Emma",
		"Georgia",
		"Harriet",
		"Joan",
		"Judy",
		"Julia",
		"Lucy",
		"Lydia",
		"Margaret",
		"Mary",
		"Molly",
		"Nora",
		"Rosie",
		"Sarah",
		"Victoria",
		"Aldington",
		"Appleton",
		"Ashdon",
		"Berwick",
		"Bramford",
		"Brimstage",
		"Carden",
		"Churchill",
		"Clifton",
		"Colby",
		"Copford",
		"Cromer",
		"Davenham",
		"Dersingham",
		"Doverdale",
		"Elsted",
		"Ferring",
		"Gissing",
		"Heydon",
		"Holt",
		"Hunston",
		"Hutton",
		"Inkberrow",
		"Inworth",
		"Isfield",
		"Kedington",
		"Latchford",
		"Leigh",
		"Leighton",
		"Maresfield",
		"Markshall",
		"Netherpool",
		"Newton",
		"Oxton",
		"Preston",
		"Ridley",
		"Rochford",
		"Seaford",
		"Selsey",
		"Stanton",
		"Stockham",
		"Stoke",
		"Sutton",
		"Thakeham",
		"Thetford",
		"Thorndon",
		"Ulting",
		"Upton",
		"Westhorpe",
		"Worcester",
		"Alexander",
		"Alexius",
		"Anastasius",
		"Christodoulos",
		"Christos",
		"Damian",
		"Dimitris",
		"Dysmas",
		"Elias",
		"Giorgos",
		"Ioannis",
		"Konstantinos",
		"Lambros",
		"Leonidas",
		"Marcos",
		"Miltiades",
		"Nestor",
		"Nikos",
		"Orestes",
		"Petros",
		"Simon",
		"Stavros",
		"Theodore",
		"Vassilios",
		"Yannis",
		"Alexandra",
		"Amalia",
		"Callisto",
		"Charis",
		"Chloe",
		"Dorothea",
		"Elena",
		"Eudoxia",
		"Giada",
		"Helena",
		"Ioanna",
		"Lydia",
		"Melania",
		"Melissa",
		"Nika",
		"Nikolina",
		"Olympias",
		"Philippa",
		"Phoebe",
		"Sophia",
		"Theodora",
		"Valentina",
		"Valeria",
		"Yianna",
		"Zoe",
		"Adramyttion",
		"Ainos",
		"Alikarnassos",
		"Avydos",
		"Dakia",
		"Dardanos",
		"Dekapoli",
		"Dodoni",
		"Efesos",
		"Efstratios",
		"Elefsina",
		"Ellada",
		"Epidavros",
		"Erymanthos",
		"Evripos",
		"Gavdos",
		"Gytheio",
		"Ikaria",
		"Ilios",
		"Illyria",
		"Iraia",
		"Irakleio",
		"Isminos",
		"Ithaki",
		"Kadmeia",
		"Kallisto",
		"Katerini",
		"Kithairon",
		"Kydonia",
		"Lakonia",
		"Leros",
		"Lesvos",
		"Limnos",
		"Lykia",
		"Megara",
		"Messene",
		"Milos",
		"Nikaia",
		"Orontis",
		"Parnasos",
		"Petro",
		"Samos",
		"Syros",
		"Thapsos",
		"Thessalia",
		"Thira",
		"Thiva",
		"Varvara",
		"Voiotia",
		"Vyvlos",
		"Amrit",
		"Ashok",
		"Chand",
		"Dinesh",
		"Gobind",
		"Harinder",
		"Jagdish",
		"Johar",
		"Kurien",
		"Lakshman",
		"Madhav",
		"Mahinder",
		"Mohal",
		"Narinder",
		"Nikhil",
		"Omrao",
		"Prasad",
		"Pratap",
		"Ranjit",
		"Sanjay",
		"Shankar",
		"Thakur",
		"Vijay",
		"Vipul",
		"Yash",
		"Amala",
		"Asha",
		"Chandra",
		"Devika",
		"Esha",
		"Gita",
		"Indira",
		"Indrani",
		"Jaya",
		"Jayanti",
		"Kiri",
		"Lalita",
		"Malati",
		"Mira",
		"Mohana",
		"Neela",
		"Nita",
		"Rajani",
		"Sarala",
		"Sarika",
		"Sheela",
		"Sunita",
		"Trishna",
		"Usha",
		"Vasanta",
		"Ahmedabad",
		"Alipurduar",
		"Alubari",
		"Anjanadri",
		"Ankleshwar",
		"Balarika",
		"Bhanuja",
		"Bhilwada",
		"Brahmaghosa",
		"Bulandshahar",
		"Candrama",
		"Chalisgaon",
		"Chandragiri",
		"Charbagh",
		"Chayanka",
		"Chittorgarh",
		"Dayabasti",
		"Dikpala",
		"Ekanga",
		"Gandhidham",
		"Gollaprolu",
		"Grahisa",
		"Guwahati",
		"Haridasva",
		"Indraprastha",
		"Jaisalmer",
		"Jharonda",
		"Kadambur",
		"Kalasipalyam",
		"Karnataka",
		"Kutchuhery",
		"Lalgola",
		"Mainaguri",
		"Nainital",
		"Nandidurg",
		"Narayanadri",
		"Panipat",
		"Panjagutta",
		"Pathankot",
		"Pathardih",
		"Porbandar",
		"Rajasthan",
		"Renigunta",
		"Sewagram",
		"Shakurbasti",
		"Siliguri",
		"Sonepat",
		"Teliwara",
		"Tinpahar",
		"Villivakkam",
		"Akira",
		"Daisuke",
		"Fukashi",
		"Goro",
		"Hiro",
		"Hiroya",
		"Hotaka",
		"Katsu",
		"Katsuto",
		"Keishuu",
		"Kyuuto",
		"Mikiya",
		"Mitsunobu",
		"Mitsuru",
		"Naruhiko",
		"Nobu",
		"Shigeo",
		"Shigeto",
		"Shou",
		"Shuji",
		"Takaharu",
		"Teruaki",
		"Tetsushi",
		"Tsukasa",
		"Yasuharu",
		"Aemi",
		"Airi",
		"Ako",
		"Ayu",
		"Chikaze",
		"Eriko",
		"Hina",
		"Kaori",
		"Keiko",
		"Kyouka",
		"Mayumi",
		"Miho",
		"Namiko",
		"Natsu",
		"Nobuko",
		"Rei",
		"Ririsa",
		"Sakimi",
		"Shihoko",
		"Shika",
		"Tsukiko",
		"Tsuzune",
		"Yoriko",
		"Yorimi",
		"Yoshiko",
		"Agrippa",
		"Appius",
		"Aulus",
		"Caeso",
		"Decimus",
		"Faustus",
		"Gaius",
		"Gnaeus",
		"Hostus",
		"Lucius",
		"Mamercus",
		"Manius",
		"Marcus",
		"Mettius",
		"Nonus",
		"Numerius",
		"Opiter",
		"Paulus",
		"Proculus",
		"Publius",
		"Quintus",
		"Servius",
		"Tiberius",
		"Titus",
		"Volescus",
		"Appia",
		"Aula",
		"Caesula",
		"Decima",
		"Fausta",
		"Gaia",
		"Gnaea",
		"Hosta",
		"Lucia",
		"Maio",
		"Marcia",
		"Maxima",
		"Mettia",
		"Nona",
		"Numeria",
		"Octavia",
		"Postuma",
		"Prima",
		"Procula",
		"Septima",
		"Servia",
		"Tertia",
		"Tiberia",
		"Titia",
		"Vibia",
		"Adesegun",
		"Akintola",
		"Amabere",
		"Arikawe",
		"Asagwara",
		"Chidubem",
		"Chinedu",
		"Chiwetei",
		"Damilola",
		"Esangbedo",
		"Ezenwoye",
		"Folarin",
		"Genechi",
		"Idowu",
		"Kelechi",
		"Ketanndu",
		"Melubari",
		"Nkanta",
		"Obafemi",
		"Olatunde",
		"Olumide",
		"Tombari",
		"Udofia",
		"Uyoata",
		"Uzochi",
		"Abike",
		"Adesuwa",
		"Adunola",
		"Anguli",
		"Arewa",
		"Asari",
		"Bisola",
		"Chioma",
		"Eduwa",
		"Emilohi",
		"Fehintola",
		"Folasade",
		"Mahparah",
		"Minika",
		"Nkolika",
		"Nkoyo",
		"Nuanae",
		"Obioma",
		"Olafemi",
		"Shanumi",
		"Sominabo",
		"Suliat",
		"Tariere",
		"Temedire",
		"Yemisi",
		"Aleksandr",
		"Andrei",
		"Arkady",
		"Boris",
		"Dmitri",
		"Dominik",
		"Grigory",
		"Igor",
		"Ilya",
		"Ivan",
		"Kiril",
		"Konstantin",
		"Leonid",
		"Nikolai",
		"Oleg",
		"Pavel",
		"Petr",
		"Sergei",
		"Stepan",
		"Valentin",
		"Vasily",
		"Viktor",
		"Yakov",
		"Yegor",
		"Yuri",
		"Aleksandra",
		"Anastasia",
		"Anja",
		"Catarina",
		"Devora",
		"Dima",
		"Ekaterina",
		"Eva",
		"Irina",
		"Karolina",
		"Katlina",
		"Kira",
		"Ludmilla",
		"Mara",
		"Nadezdha",
		"Nastassia",
		"Natalya",
		"Oksana",
		"Olena",
		"Olga",
		"Sofia",
		"Svetlana",
		"Tatyana",
		"Vilma",
		"Yelena",
		"Alejandro",
		"Alonso",
		"Amelio",
		"Armando",
		"Bernardo",
		"Carlos",
		"Cesar",
		"Diego",
		"Emilio",
		"Estevan",
		"Felipe",
		"Francisco",
		"Guillermo",
		"Javier",
		"Jose",
		"Juan",
		"Julio",
		"Luis",
		"Pedro",
		"Raul",
		"Ricardo",
		"Salvador",
		"Santiago",
		"Valeriano",
		"Vicente",
		"Adalina",
		"Aleta",
		"Ana",
		"Ascencion",
		"Beatriz",
		"Carmela",
		"Celia",
		"Dolores",
		"Elena",
		"Emelina",
		"Felipa",
		"Inez",
		"Isabel",
		"Jacinta",
		"Lucia",
		"Lupe",
		"Maria",
		"Marta",
		"Nina",
		"Paloma",
		"Rafaela",
		"Soledad",
		"Teresa",
		"Valencia",
		"Zenaida",
	}
	die := strconv.Itoa(len(names))
	return names[dice.Roll("1d"+die).Sum()-1]
}

func familyName() string {
	names := []string{
		"Arellano",
		"Arispana",
		"Borrego",
		"Carderas",
		"Carranzo",
		"Cordova",
		"Enciso",
		"Espejo",
		"Gavilan",
		"Guerra",
		"Guillen",
		"Huertas",
		"Illan",
		"Jurado",
		"Moretta",
		"Motolinia",
		"Pancorbo",
		"Paredes",
		"Quesada",
		"Roma",
		"Rubiera",
		"Santoro",
		"Torrillas",
		"Vera",
		"Vivero",
		"Aguascebas",
		"Alcazar",
		"Barranquete",
		"Bravatas",
		"Cabezudos",
		"Calderon",
		"Cantera",
		"Castillo",
		"Delgadas",
		"Donablanca",
		"Encinetas",
		"Estrella",
		"Faustino",
		"Fuentebravia",
		"Gafarillos",
		"Gironda",
		"Higueros",
		"Huelago",
		"Humilladero",
		"Illora",
		"Isabela",
		"Izbor",
		"Jandilla",
		"Jinetes",
		"Limones",
		"Loreto",
		"Lujar",
		"Marbela",
		"Matagorda",
		"Nacimiento",
		"Niguelas",
		"Ogijares",
		"Ortegicar",
		"Pampanico",
		"Pelado",
		"Quesada",
		"Quintera",
		"Riguelo",
		"Ruescas",
		"Salteras",
		"Santopitar",
		"Taberno",
		"Torres",
		"Umbrete",
		"Valdecazorla",
		"Velez",
		"Vistahermosa",
		"Yeguas",
		"Zahora",
		"Zumeta",
		"Abdel",
		"Awad",
		"Dahhak",
		"Essa",
		"Hanna",
		"Harbi",
		"Hassan",
		"Isa",
		"Kasim",
		"Katib",
		"Khalil",
		"Malik",
		"Mansoor",
		"Mazin",
		"Musa",
		"Najeeb",
		"Namari",
		"Naser",
		"Rahman",
		"Rasheed",
		"Saleh",
		"Salim",
		"Shadi",
		"Sulaiman",
		"Tabari",
		"Bai",
		"Cao",
		"Chen",
		"Cui",
		"Ding",
		"Du",
		"Fang",
		"Fu",
		"Guo",
		"Han",
		"Hao",
		"Huang",
		"Lei",
		"Li",
		"Liang",
		"Liu",
		"Long",
		"Song",
		"Tan",
		"Tang",
		"Wang",
		"Wu",
		"Xing",
		"Yang",
		"Zhang",
		"Barker",
		"Brown",
		"Butler",
		"Carter",
		"Chapman",
		"Collins",
		"Cook",
		"Davies",
		"Gray",
		"Green",
		"Harris",
		"Jackson",
		"Jones",
		"Lloyd",
		"Miller",
		"Roberts",
		"Smith",
		"Taylor",
		"Thomas",
		"Turner",
		"Watson",
		"White",
		"Williams",
		"Wood",
		"Young",
		"Andreas",
		"Argyros",
		"Dimitriou",
		"Floros",
		"Gavras",
		"Ioannidis",
		"Katsaros",
		"Kyrkos",
		"Leventis",
		"Makris",
		"Metaxas",
		"Nikolaidis",
		"Pallis",
		"Pappas",
		"Petrou",
		"Raptis",
		"Simonides",
		"Spiros",
		"Stavros",
		"Stephanidis",
		"Stratigos",
		"Terzis",
		"Theodorou",
		"Vasiliadis",
		"Yannakakis",
		"Achari",
		"Banerjee",
		"Bhatnagar",
		"Bose",
		"Chauhan",
		"Chopra",
		"Das",
		"Dutta",
		"Gupta",
		"Johar",
		"Kapoor",
		"Mahajan",
		"Malhotra",
		"Mehra",
		"Nehru",
		"Patil",
		"Rao",
		"Saxena",
		"Shah",
		"Sharma",
		"Singh",
		"Trivedi",
		"Venkatesan",
		"Verma",
		"Yadav",
		"Abe",
		"Arakaki",
		"Endo",
		"Fujiwara",
		"Goto",
		"Ito",
		"Kikuchi",
		"Kinjo",
		"Kobayashi",
		"Koga",
		"Komatsu",
		"Maeda",
		"Nakamura",
		"Narita",
		"Ochi",
		"Oshiro",
		"Saito",
		"Sakamoto",
		"Sato",
		"Suzuki",
		"Takahashi",
		"Tanaka",
		"Watanabe",
		"Yamamoto",
		"Yamasaki",
		"Bando",
		"Chikuma",
		"Chikusei",
		"Chino",
		"Hitachi",
		"Hitachinaka",
		"Hitachiomiya",
		"Hitachiota",
		"Iida",
		"Iiyama",
		"Ina",
		"Inashiki",
		"Ishioka",
		"Itako",
		"Kamisu",
		"Kasama",
		"Kashima",
		"Kasumigaura",
		"Kitaibaraki",
		"Kiyose",
		"Koga",
		"Komagane",
		"Komoro",
		"Matsumoto",
		"Mito",
		"Mitsukaido",
		"Moriya",
		"Nagano",
		"Naka",
		"Nakano",
		"Ogi",
		"Okaya",
		"Omachi",
		"Ryugasaki",
		"Saku",
		"Settsu",
		"Shimotsuma",
		"Shiojiri",
		"Suwa",
		"Suzaka",
		"Takahagi",
		"Takeo",
		"Tomi",
		"Toride",
		"Tsuchiura",
		"Tsukuba",
		"Ueda",
		"Ushiku",
		"Yoshikawa",
		"Yuki",
		"Antius",
		"Aurius",
		"Barbatius",
		"Calidius",
		"Cornelius",
		"Decius",
		"Fabius",
		"Flavius",
		"Galerius",
		"Horatius",
		"Julius",
		"Juventius",
		"Licinius",
		"Marius",
		"Minicius",
		"Nerius",
		"Octavius",
		"Pompeius",
		"Quinctius",
		"Rutilius",
		"Sextius",
		"Titius",
		"Ulpius",
		"Valerius",
		"Vitellius",
		"Adegboye",
		"Adeniyi",
		"Adeyeku",
		"Adunola",
		"Agbaje",
		"Akpan",
		"Akpehi",
		"Aliki",
		"Asuni",
		"Babangida",
		"Ekim",
		"Ezeiruaku",
		"Fabiola",
		"Fasola",
		"Nwokolo",
		"Nzeocha",
		"Ojo",
		"Okonkwo",
		"Okoye",
		"Olaniyan",
		"Olawale",
		"Olumese",
		"Onajobi",
		"Soyinka",
		"Yamusa",
		"Abadan",
		"Ador",
		"Agatu",
		"Akamkpa",
		"Akpabuyo",
		"Ala",
		"Askira",
		"Bakassi",
		"Bama",
		"Bayo",
		"Bekwara",
		"Biase",
		"Boki",
		"Buruku",
		"Calabar",
		"Chibok",
		"Damboa",
		"Dikwa",
		"Etung",
		"Gboko",
		"Gubio",
		"Guzamala",
		"Gwoza",
		"Hawul",
		"Ikom",
		"Jere",
		"Kalabalge",
		"Katsina",
		"Knoduga",
		"Konshishatse",
		"Kukawa",
		"Kwande",
		"Kwayakusar",
		"Logo",
		"Mafa",
		"Makurdi",
		"Nganzai",
		"Obanliku",
		"Obi",
		"Obubra",
		"Obudu",
		"Odukpani",
		"Ogbadibo",
		"Ohimini",
		"Okpokwu",
		"Otukpo",
		"Shani",
		"Ugep",
		"Vandeikya",
		"Yala",
		"Abelev",
		"Bobrikov",
		"Chemerkin",
		"Gogunov",
		"Gurov",
		"Iltchenko",
		"Kavelin",
		"Komarov",
		"Korovin",
		"Kurnikov",
		"Lebedev",
		"Litvak",
		"Mekhdiev",
		"Muraviov",
		"Nikitin",
		"Ortov",
		"Peshkov",
		"Romasko",
		"Shvedov",
		"Sikorski",
		"Stolypin",
		"Turov",
		"Volokh",
		"Zaitsev",
		"Zhukov",
	}
	die := strconv.Itoa(len(names))
	return names[dice.Roll("1d"+die).Sum()-1]
}

func randomName() string {
	return firstName() + " " + familyName()
}

func describeReaction(r int) string {
	ds := ""
	switch r {
	default:
		if r < 3 {
			ds = "Hostile. "
			if dice.Roll("2d6").Sum() >= 8 {
				ds += "Attacks first if situation allows. "
			} else {
				ds += "Attacks if not Persuaded. "
			}
			ds += "Interaction DM: " + strconv.Itoa(r-6)
		}
		if r > 11 {
			ds = "Enthusiastic. Interaction DM: +" + strconv.Itoa(r-9)
		}
	case 3:
		ds = "Hostile. "
		if dice.Roll("2d6").Sum() >= 5 {
			ds += "Provoking fight if situation allows. "
		}
		ds += "Interaction DM: -3"
	case 4:
		ds = "Hostile. "
		if dice.Roll("2d6").Sum() >= 8 {
			ds += "Provoking fight if situation allows. "
		}
		ds += "Interaction DM: -2"
	case 5:
		ds = "Hostile. "
		if dice.Roll("2d6").Sum() >= 11 {
			ds += "Provoking fight if situation allows. "
		}
		ds += "Interaction DM: -1"
	case 6:
		ds = "Unreseptive. (Impolite)"
	case 7:
		ds = "Uninterested"
	case 8:
		ds = "Noncommitial (Curt)"
	case 9:
		ds = "Noncommitial (Polite)"
	case 10:
		ds = "Interested. Interactions DM: +1"
	case 11:
		ds = "Responsive. Interactions DM: +2"
	}

	return ds
}
