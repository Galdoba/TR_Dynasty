package starsystem

import (
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/Astrogation"
)

type bodyDetails struct {
	nomena           string
	name             string
	uwp              string
	tags             string
	bodyType         string
	orbitDistance    float64 //радиус орбиты от звезды в AU
	orbitDistanceSat float64 //радиус орбиты от планеты в ММ
	jumpPointToBody  float64 //расстояние до точки прыжка с учетом (пока нет) тени звезды
	orbitSpeed       int
	position         numCode //
	//SISE RELATED:
	diameter      float64
	planetDensity int
}

const (
	DensityHeavyCore  = 0
	DensityMoltenCore = 1
	DensityRockyBody  = 2
	DensityIcyBody    = 3
)

/*

world.StarsystemDetails() []string
starsystem.Details()
*/

// type SystemDetails struct {
// 	bodyDetail map[string]string
// 	//PositionFromStar__PositionFromPlanet__Name__UWP__Actual Orbit
// 	/*
// 		Primary		Fijari	Ka V
// 		0		Keetle	SGG
// 		5		Ra-La-Lantra	LGG
// 		5	0	Ring System	YR00000-0	3.28
// 		5	8	B'kolior	X621000-0	3.28
// 	*/
// }

// func from(world wrld.World) SystemDetails {
// 	d := SystemDetails{}
// 	d.bodyDetail = make(map[string]string)
// 	return d
// }

func (bd *bodyDetails) ShortInfo() string {
	str := ""
	str += code2string(bd.position) + "	"
	if bd.position.planetCode() != -1 {
		str += "    "
	}
	if bd.position.sateliteCode() != -1 {
		str += "    "
	}

	str += bd.nomena

	if len(bd.nomena) < 8 {
		str += "    "
	}
	nLen := len(strings.Split(bd.nomena, " "))
	switch nLen {
	default:
		str += "	  "
	case 2:
		str += "	" + strconv.FormatFloat(bd.orbitDistance, 'f', 2, 64) + " au"
	case 3:
		str += "	" + strconv.FormatFloat(bd.orbitDistanceSat, 'f', 2, 64) + " Mm"
	}
	str += "	" + bd.bodyType
	str += "	" + bd.uwp
	str += "	" + bd.tags
	if bd.name == "" {
		str += "         "
	}
	str += "	" + bd.name
	return str
}

func (bd *bodyDetails) FullInfo() string {
	str := bd.ShortInfo()
	str += "\nASTROGATION DATA:\n"
	str += "Orbital Distance	: " + strconv.FormatFloat(bd.orbitDistance, 'f', 2, 64) + " au\n"
	str += "Jump point To Orbit	: " + strconv.FormatFloat(bd.jumpPointToBody*1000, 'f', 0, 64) + " km\n"
	for i := 1; i <= 7; i++ {
		str += "	Thrust " + strconv.Itoa(i) + ": " + Astrogation.TravelTimeStr(bd.jumpPointToBody, float64(i)) + "\n"
	}
	str += "Diameter	: " + strconv.FormatFloat(bd.diameter, 'f', 2, 64) + " Mm\n"
	return str
}

//---- -------------------- --------- -------------------------- ------ ------- ------ ----- -- - --- -- ---- --------------
//---- -------------------- --------- ---------------------------------------- ------ ------- ------ ----- -- - --- -- ---- --------------
//Sector	SS	Hex	Name	UWP	Bases	Remarks	Zone	PBG	Allegiance	Stars	{Ix}	(Ex)	[Cx]	Nobility	W	RU

/*
+----------------------------------------------------------------------------------------------
|Alpha 3       | Borite | E655796-4 | Mainworld | Ag Ga Lt Fr |
|


*/
