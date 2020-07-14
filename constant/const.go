package constant

const (
	DIVIDER       = "-"
	UNKNOWN       = "_"
	PrStarport    = "Starport"
	PrSize        = "Size"
	PrAtmo        = "Atmo"
	PrHydr        = "Hydr"
	PrPops        = "Pops"
	PrGovr        = "Govr"
	PrLaws        = "Laws"
	PrTL          = "TL"
	WTpHospitable = "Hospitable"
	WTpPlanetoid  = "Planetoid"
	WTpIceWorld   = "IceWorld"
	WTpRadWorld   = "RadWorld"
	WTpInferno    = "Inferno"
	WTpBigWorld   = "BigWorld"
	WTpWorldlet   = "Worldlet"
	WTpInnerWorld = "InnerWorld"
	WTpStormWorld = "StormWorld"
	WTpGG         = "GG"
)

func WorldTypeValid(wType string) bool {
	allTypes := []string{
		WTpHospitable,
		WTpPlanetoid,
		WTpIceWorld,
		WTpRadWorld,
		WTpInferno,
		WTpBigWorld,
		WTpWorldlet,
		WTpInnerWorld,
		WTpStormWorld,
		WTpGG,
	}
	for i := range allTypes {
		if wType == allTypes[i] {
			return true
		}
	}
	return false
}

func DEBUGallWorldTypes() []string {
	return []string{
		WTpHospitable,
		WTpPlanetoid,
		WTpIceWorld,
		WTpRadWorld,
		WTpInferno,
		WTpBigWorld,
		WTpWorldlet,
		WTpInnerWorld,
		WTpStormWorld,
		WTpGG,
	}
}
