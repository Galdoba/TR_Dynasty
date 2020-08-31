package constant

const (
	//DIVIDER - не несет утверждения. Просто разделаяет части строки
	DIVIDER = "-"
	//UNKNOWN - неизвестный EHEX
	UNKNOWN             = "_"
	PrStarport          = "Starport"
	PrSize              = "Size"
	PrAtmo              = "Atmo"
	PrHydr              = "Hydr"
	PrPops              = "Pops"
	PrGovr              = "Govr"
	PrLaws              = "Laws"
	PrTL                = "TL"
	WTpHospitable       = "Hospitable"
	WTpPlanetoid        = "Planetoid"
	WTpIceWorld         = "IceWorld"
	WTpRadWorld         = "RadWorld"
	WTpInferno          = "Inferno"
	WTpBigWorld         = "BigWorld"
	WTpWorldlet         = "Worldlet"
	WTpInnerWorld       = "InnerWorld"
	WTpStormWorld       = "StormWorld"
	WTpGG               = "GG"
	MerchantTypeCommon  = "Common"
	MerchantTypeTrade   = "Trade"
	MerchantTypeNeutral = "Neutral"
	MerchantTypeIlligal = "Illigal"
	//MgT2 Codes
	TradeCodeAgricultural    = "Ag"
	TradeCodeAsteroid        = "As"
	TradeCodeBarren          = "Ba"
	TradeCodeDesert          = "De"
	TradeCodeFluidOceans     = "Fl"
	TradeCodeGarden          = "Ga"
	TradeCodeHighPopulation  = "Hi"
	TradeCodeHighTech        = "Ht"
	TradeCodeIceCapped       = "Ic"
	TradeCodeIndustrial      = "In"
	TradeCodeLowPopulation   = "Lo"
	TradeCodeLowTech         = "Lt"
	TradeCodeNonAgricultural = "Na"
	TradeCodeNonIndustrial   = "Ni"
	TradeCodePoor            = "Po"
	TradeCodeRich            = "Ri"
	TradeCodeVacuum          = "Va"
	TradeCodeWaterWorld      = "Wa"
	TravelCodeAmber          = "A"
	TravelCodeRed            = "R"
	TradeCodeAsteroidBelt    = "As"
	//T5 Codes
	TradeCodeHellworld        = "He"
	TradeCodeOceanWorld       = "Oc"
	TradeCodeDieback          = "Di"
	TradeCodePreHigh          = "Ph"
	TradeCodePreAgricultural  = "Pa"
	TradeCodePrisonExileCamp  = "Px"
	TradeCodePreIndustrial    = "Pi"
	TradeCodePreRich          = "Pr"
	TradeCodeFrozen           = "Fr"
	TradeCodeHot              = "Ho"
	TradeCodeCold             = "Co"
	TradeCodeLocked           = "Lk"
	TradeCodeTropic           = "Tr"
	TradeCodeTundra           = "Tu"
	TradeCodeTwilightZone     = "Tz"
	TradeCodeFarming          = "Fa"
	TradeCodeMining           = "Mi"
	TradeCodeMilitaryRule     = "Mr"
	TradeCodePenalColony      = "Pe"
	TradeCodeReserve          = "Re"
	TradeCodeSubsectorCapital = "Cp"
	TradeCodeSectorCapital    = "Cs"
	TradeCodeCapital          = "Cx"
	TradeCodeColony           = "Cy"
	TradeCodeSatellite        = "Sa"
	TradeCodeForbidden        = "Fo"
	TradeCodePuzzle           = "Pz"
	TradeCodeDangerous        = "Da"
	TradeCodeDataRepository   = "Ab"
	TradeCodeAncientSite      = "An"
)

func TravelCodesMgT2() []string {
	return []string{
		TradeCodeAgricultural,
		TradeCodeAsteroid,
		TradeCodeBarren,
		TradeCodeDesert,
		TradeCodeFluidOceans,
		TradeCodeGarden,
		TradeCodeHighPopulation,
		TradeCodeHighTech,
		TradeCodeIceCapped,
		TradeCodeIndustrial,
		TradeCodeLowPopulation,
		TradeCodeLowTech,
		TradeCodeNonAgricultural,
		TradeCodeNonIndustrial,
		TradeCodePoor,
		TradeCodeRich,
		TradeCodeVacuum,
		TradeCodeWaterWorld,
		TravelCodeAmber,
		TravelCodeRed,
	}
}

//WorldTypeValid - Проверка правильности написания типа Планеты
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

//DEBUGallWorldTypes - временная функция
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
