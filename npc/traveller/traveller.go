package traveller

import "github.com/Galdoba/TR_Dynasty/npc/traveller/characteristic"

type Traveller struct {
	gmmod               bool
	CharacteristicsCore []*characteristic.Characteristic
}

func New(gmmod bool, name, homeworldUWP, race string) *Traveller {
	//0
	t := Traveller{}
	t.gmmod = gmmod

	return &t
}

type PesonalData struct {
	Name      string
	Age       int
	Specie    string
	Traits    []string //temp
	Homeworld string
	Rads      int
}

/*
Generator steps:
0. Construct Object
1. Characteristics
2. Background
3. Aperiance
4. Pre-Career
5. Careers




*/
