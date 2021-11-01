package orbit

import (
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/universe/survey"
)

type PlacementData struct {
	pDice        *dice.Dicepool
	placementMap map[orbitCode]bodyData
}

type orbitCode struct {
	code string
}

type bodyData struct {
	bType  string
	parent *bodyData
	child  *bodyData
}

func PlaceWorlds(ss survey.SecondSurveyData) *PlacementData {
	//pd := PlacementData{}
	return nil
	//return &pd
}
