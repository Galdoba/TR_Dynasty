package starsystem

import (
	"strconv"
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
	str += bd.nomena
	if len(bd.nomena) < 8 {
		str += "    "
	}
	str += "	" + bd.name
	if bd.name == "" {
		str += "         "
	}
	str += "	" + bd.uwp
	str += "	" + bd.bodyType
	str += "	" + bd.tags
	return str
}

func (bd *bodyDetails) FullInfo() string {
	str := bd.ShortInfo()
	str += "\nASTROGATION DATA:\n"
	str += "Orbital Distance	: " + strconv.FormatFloat(bd.orbitDistance, 'f', 2, 64) + " au\n"
	// str += "Jump point To Orbit	: " + strconv.FormatFloat(bd.jumpPointToOrbit*1000, 'f', 0, 64) + " km\n"
	// for i := 1; i <= 7; i++ {
	// 	str += "	Thrust " + strconv.Itoa(i) + ": " + Astrogation.TravelTimeStr(bd.jumpPointToOrbit, float64(i)) + "\n"
	// }

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
