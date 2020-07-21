package npc

import (
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/utils"
)

type NPCensembleCast struct {
	name          string
	race          string
	upp           string
	relationships int
	quirk1        string
	quirk2        string
	ocupation     string
	skillList     int
	age           int
}

func RandomNPC() NPCensembleCast {
	npc := NPCensembleCast{}
	npc.name = "[NPC_Name]"
	npc.race = "Human"
	npc.quirk1 = TrvCore.RollD66()
	npc.quirk2 = "00"
	if utils.RollDiceRandom("d2") > 1 {
		npc.quirk2 = TrvCore.RollD66()
	}
	npc.ocupation = TrvCore.RollD66()
	npc.skillList = utils.RollDiceRandom("2d6")
	npc.age = 18 + utils.RollDiceRandom("6d6")
	return npc
}
