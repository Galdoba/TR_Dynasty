package npc

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/dice"
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

type NPCensembleCast struct {
	name          string
	race          string
	upp           string
	relationships int
	quirk1        string
	quirk2        string
	occupation    string
	skillList     int
	age           int
}

func RandomNPC() NPCensembleCast {
	npc := NPCensembleCast{}
	npc.name = "[NPC_Name]"
	npc.race = "Human"
	npc.quirk1 = TrvCore.RollD66()
	npc.quirk1 = dice.Roll("2d6").ResultString()
	npc.quirk2 = "00"
	if utils.RollDiceRandom("d2") > 1 {
		npc.quirk2 = TrvCore.RollD66()
	}
	npc.occupation = occupationChart()
	npc.occupation = occClergy
	npc.skillList = utils.RollDiceRandom("2d6")
	npc.age = 18 + utils.RollDiceRandom("6d6")
	fmt.Println(npc)
	npc.rollStats()
	return npc
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
		[]string{"10", "8", "8", "8", "7", "7", "7"},
		[]string{"11", "7", "7", "7", "8", "8", "8"},
		[]string{"12", "8", "8", "8", "8", "8", "8"},
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
	tab.occ[occCorporateShipper] = tabConstructor([]string{"2", "51", "63", "52", "71", "61", "61"})

	return tab
}

func tabConstructor(st []string) [][]string {
	var res [][]string

	for i := range st {
		switch i {
		case 0:
			res = append(res, []string{"2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C"})
		case 1, 2, 3, 4, 5, 6:
			res = append(res, fromCode(st[i]))
		}
	}
	return res
}

func fromCode(s string) []string {
	var res []string
	i := TrvCore.EhexToDigit(s)
	val := i / 10
	it := i - (val * 10)
	for l := 0; l < 11; l++ {
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
		fmt.Println(npc.occupation, r, c, tab.occ[npc.occupation][c][r])
		upp = upp + tab.occ[npc.occupation][c][r]
	}
	npc.upp = upp
}

func occupationChart() string {
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
	return occupation[rand.Intn(len(occupation))]
}
