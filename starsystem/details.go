package starsystem

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/Astrogation"
	"github.com/Galdoba/TR_Dynasty/pkg/core/astronomical"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/utils"
)

//BodyDetails -
type BodyDetails struct {
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
	playersInOrbit   bool
	//SISE RELATED:
	diameter      float64
	densityType   int
	planetDensity float64
	parentStar    string
	stringKey     string
}

func (bd *BodyDetails) PositionCode() string {
	return code2string(bd.position)
}

func (bd *BodyDetails) ParentStar() string {
	return bd.parentStar
}

func (bd *BodyDetails) UWPstr() string {
	return bd.uwp
}

func (bd *BodyDetails) BodyType() string {
	return bd.bodyType
}

func (bd *BodyDetails) PlanetDensity() float64 {
	return bd.planetDensity
}

const (
	DensityUndefined  = 0
	DensityHeavyCore  = 1
	DensityMoltenCore = 2
	DensityRockyBody  = 3
	DensityIcyBody    = 4
)

func (bd *BodyDetails) calculatePlanetDensity() {
	if bd.densityType == 0 {
		fmt.Println(bd)
		switch bd.bodyType {
		default:
		case "Star":
			panic(bd.bodyType)
		}

	}
	bd.planetDensity = 5.00
}

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

func (bd *BodyDetails) ShortInfo() string {
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
		str += "	 " // + strconv.FormatFloat(bd.orbitDistanceSat, 'f', 2, 64) + " Mm"
	}
	str += "	" + bd.bodyType
	str += "	" + bd.uwp
	str += "	" + bd.tags
	if bd.name == "" {
		str += "         "
	}
	str += "	" + bd.name

	js := astronomical.ShadowOrbit(bd.parentStar)
	switch bd.position.planetCode() <= js {
	case true:
		str += "	" + "X"
	case false:
		str += "	" + ""
	}
	switch bd.playersInOrbit {
	case true:
		str += "	" + "P"
	case false:
		str += "	" + ""
	}
	return str
}

func (bd *BodyDetails) FullInfo() string {
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
Name       : Borite (Alpha 4 Ey)
Type       : Stormworld (Mainworld: B189456-4)
Starport   : A (600/200 Highport TAS Naval Scout)
Physical   : 4 (7623 km / 0.42 g)
Atmosphere : 7 (N2O2, with Sulfur Compounds / 1.00 Ta)
Hydrosphere: 2 (N2O2, with Sulfur Compounds / 1.00 Ta)


*/

/////////////////////ASTEROID BELTS
func (bd *BodyDetails) AsteroidDetails() string {
	dp := dice.New().SetSeed(bd.nomena)
	r1 := dp.RollNext("2d6").Sum()
	r2 := dp.RollNext("1d6").Sum()
	preDomSize := "undefined"
	switch r1 {
	case 2:
		preDomSize = "1m"
	case 3:
		preDomSize = "5m"
	case 4:
		preDomSize = "10m"
	case 5:
		preDomSize = "25m"
	case 6:
		preDomSize = "50m"
	case 7:
		preDomSize = "100m"
	case 8:
		preDomSize = "300m"
	case 9:
		preDomSize = "1km"
	case 10:
		preDomSize = "5km"
	case 11:
		preDomSize = "50km"
	case 12:
		preDomSize = "500km"
	}
	maxSize := "undefined"
	switch r2 {
	case 1, 2:
		maxSize = preDomSize
	case 3:
		maxSize = "1km"
	case 4:
		maxSize = "10km"
	case 5:
		maxSize = "100km"
	case 6:
		maxSize = "1000km"
	}
	orbZone := astronomical.Zone(bd.position.planetCode(), bd.parentStar)
	zoneDM := 0
	switch orbZone {
	case astronomical.ZoneInner:
		zoneDM = -4
	case astronomical.ZoneOuter:
		zoneDM = 2
	}
	beltZone := ""
	dp.RollNext("2d6").DM(zoneDM)
	switch {
	case dp.ResultIs("4-"):
		beltZone = "N-ZONE"
	case dp.ResultIs("5 8"):
		beltZone = "M-ZONE"
	case dp.ResultIs("9+"):
		beltZone = "C-ZONE"
	}
	composition := rollZonePersentage(beltZone, dp)
	width := ""
	widthDM := 2
	switch bd.position.planetCode() {
	case 0, 1, 2, 3, 4:
		widthDM = -3
	case 5, 6, 7, 8:
		widthDM = -1
	case 9, 10, 11, 12:
		widthDM = 1
	}
	r3 := dp.RollNext("2d6").DM(widthDM).Sum()
	r3 = utils.BoundInt(r3, 2, 12)
	switch r3 {
	case 2:
		width = "0.01 AU"
	case 3:
		width = "0.05 AU"
	case 4:
		width = "0.1 AU"
	case 5:
		width = "0.1 AU"
	case 6:
		width = "0.5 AU"
	case 7:
		width = "0.5 AU"
	case 8:
		width = "1.0 AU"
	case 9:
		width = "1.5 AU"
	case 10:
		width = "2.0 AU"
	case 11:
		width = "5.0 AU"
	case 12:
		width = "10.0 AU"
	}
	return preDomSize + "/" + maxSize + ", " + composition + ", " + width
}

func rollZonePersentage(zoneType string, dp *dice.Dicepool) string {
	r := dp.RollNext("2d6").Sum()
	switch zoneType {
	case "N-ZONE":
		switch r {
		case 2:
			return "n-40 m-30 c-30"
		case 3:
			return "n-40 m-40 c-20"
		case 4:
			return "n-40 m-40 c-20"
		case 5:
			return "n-40 m-40 c-20"
		case 6:
			return "n-40 m-40 c-20"
		case 7:
			return "n-50 m-40 c-10"
		case 8:
			return "n-50 m-40 c-10"
		case 9:
			return "n-50 m-40 c-10"
		case 10:
			return "n-50 m-30 c-20"
		case 11:
			return "n-60 m-50 c-10"
		case 12:
			return "n-60 m-40 c-00"
		}
	case "M-ZONE":
		switch r {
		case 2:
			return "n-20 m-30 c-50"
		case 3:
			return "n-30 m-50 c-20"
		case 4:
			return "n-20 m-60 c-20"
		case 5:
			return "n-20 m-60 c-20"
		case 6:
			return "n-30 m-60 c-10"
		case 7:
			return "n-20 m-70 c-10"
		case 8:
			return "n-10 m-70 c-20"
		case 9:
			return "n-10 m-80 c-10"
		case 10:
			return "n-10 m-80 c-10"
		case 11:
			return "n-00 m-80 c-20"
		case 12:
			return "n-00 m-90 c-10"
		}
	case "C-ZONE":
		switch r {
		case 2:
			return "n-20 m-30 c-50"
		case 3:
			return "n-20 m-30 c-50"
		case 4:
			return "n-20 m-30 c-50"
		case 5:
			return "n-10 m-30 c-60"
		case 6:
			return "n-10 m-30 c-60"
		case 7:
			return "n-10 m-20 c-70"
		case 8:
			return "n-10 m-20 c-70"
		case 9:
			return "n-10 m-10 c-80"
		case 10:
			return "n-00 m-10 c-90"
		case 11:
			return "n-00 m-10 c-90"
		case 12:
			return "n-00 m-20 c-80"
		}
	}
	return "Error"
}
