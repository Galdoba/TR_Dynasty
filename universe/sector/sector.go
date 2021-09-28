package sector

const ()

type sector struct {
	horizontal   int
	vertical     int
	abbriviation string
	names        map[string]string
	mileu        string
	tags         []string
	coreward     *sector
	rimward      *sector
	spitward     *sector
	trailward    *sector
}

/*
{"Worlds":[{
		"Name":"Reference",
		"Hex":"0140",
		"UWP":"B100727-C",
		"PBG":"302",
		"Zone":"",
		"Bases":"NS",
		"Allegiance":"ImDc",
		"Stellar":"K0 V",
		"SS":"M",
		"Ix":"{ 3 }",
		"CalculatedImportance":3,
		"Ex":"(B6D+3)",
		"Cx":"[7A5C]",
		"Nobility":"BD",
		"Worlds":12,
		"ResourceUnits":2574,
		"Subsector":12,
		"Quadrant":4,
		"WorldX":0,
		"WorldY":0,
		"Remarks":"Na Va Pi RsA Ab",
		"LegacyBaseCode":"A",
		"Sector":"Core",
		"SubsectorName":"Cadion",
		"SectorAbbreviation":"Core",
		"AllegianceName":"Third Imperium, Domain of Sylea"
		"Orbit": "A1a"
		}]
}
*/
