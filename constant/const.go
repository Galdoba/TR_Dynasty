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
)

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
