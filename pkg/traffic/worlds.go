package traffic

import (
	"github.com/Galdoba/TR_Dynasty/Astrogation"
	"github.com/Galdoba/TR_Dynasty/pkg/universe/survey"
)

const (
	SOURCE = "Source World"
	TARGET = "Target World"
)

type world struct {
	wType      string //Source/Target
	gCoordX    int    //globalX
	gCoordY    int    //globalY
	name       string
	uwp        string
	travelZone string
}

func Parse(input string) *world {
	ssdA, _ := survey.Search(input)
	ssd := ssdA[0]
	w := world{}
	w.gCoordX = ssd.CoordX
	w.gCoordY = ssd.CoordY
	w.name = ssd.MW_Name
	w.uwp = ssd.MW_UWP
	w.travelZone = ssd.TravelZone
	return &w
}

func (w *world) SetAs(wType string) {
	switch wType {
	case SOURCE, TARGET:
		w.wType = wType
	default:
		w.wType = "Trade Type Unknown"
	}
}

func Distance(sw, tw *world) int {
	return Astrogation.Distance(Astrogation.NewCoordinates(sw.gCoordX, sw.gCoordY), Astrogation.NewCoordinates(tw.gCoordX, tw.gCoordY))
}

func coordinates(w *world) Astrogation.Coordinates {
	return Astrogation.NewCoordinates(w.gCoordX, w.gCoordY)
}
