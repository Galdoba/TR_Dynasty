package orbit

import (
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/universe/survey"
)

type OrbitFiller interface {
}

type orbitfiller struct {
	dp *dice.Dicepool
}

type PlacementData struct {
	pDice        *dice.Dicepool
	placementMap map[orbitCode]bodyData
}

type orbitCode struct {
	code string
}

/*
нужно для распределения:
Stellar
PBG
Total Worlds
Remarks
FullName (seed)



Define Available Orbits
Place Mainworld - need Remarks, Stellar
	If Satellite, place GG in MW Orbit.
	If Satellite and No Giants, place a BigWorld in MW Orbit.
	If Asteroid Belt, place as Belt without regard to HZ.
Place Gas Giants Rotate Placement Per Star.
Place Planetoid Belts Rotate Placement Per Star.
Place Other Worlds Rotate Placement Per Star, place worlds using P2 World1 Column.
	Last World, place using P2 World2 Column.


*/

type bodyData struct {
	orbit           int
	decimalOrbit    float64
	orbitalDistance float64 //AU если планета, km если спутник
	name            string
	contentType     string //может быть int (код)
	uwp             string
	remarks         string
	averageTemp     int
}

func PlaceWorlds(ss survey.SecondSurveyData) *PlacementData {
	//pd := PlacementData{}
	return nil
	//return &pd
}
