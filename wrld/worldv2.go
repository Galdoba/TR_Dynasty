package wrld

import (
	"strings"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/utils"
)

const (
	worldHEX                  = "Hex"
	worldNAME                 = "Name"
	worldUWP                  = "UWP"
	worldTradeClassifications = "TC"
	worldImportanceEx         = "Ix"
	worldEconomicEx           = "Ex"
	worldCulturalEx           = "Cx"
	worldNobility             = "N"
	worldBases                = "B"
	worldTravelZone           = "Z"
	worldPBG                  = "PBG"
	worldNumOfWorlds          = "W"
	worldAllegiance           = "A"
	worldStellar              = "Stellar"
)

/*
нужен интерфейс который будет читать и редактировать информацию в мире
Profiler:
должен уметь Get() и Set()
*/

//World - упрощеная структура мира, представляет собой коллекцию профайлов
type World struct {
	data     map[string]string
	dicepool dice.Dicepool
}

//Hex - return Hex data
func (w *World) Hex() string {
	if val, ok := w.data[worldHEX]; ok {
		return val
	}
	return "--NO DATA--"
}

//Name - return Hex data
func (w *World) Name() string {
	if val, ok := w.data[worldNAME]; ok {
		return val
	}
	return "--NO DATA--"
}

//UWP - return UWP data
func (w *World) UWP() string {
	if val, ok := w.data[worldUWP]; ok {
		return val
	}
	return "--NO DATA--"
}

//SetUWP - Set UWP data
func (w *World) SetUWP(uwp string) {
	w.data[worldUWP] = uwp
}

//TradeClassifications - return TradeClassifications data
func (w *World) TradeClassifications() string {
	if val, ok := w.data[worldTradeClassifications]; ok {
		return val
	}
	return "--NO DATA--"
}

//ImportanceEx - return Importance data
func (w *World) ImportanceEx() string {
	if val, ok := w.data[worldImportanceEx]; ok {
		return val
	}
	return "--NO DATA--"
}

// //Importance - return Importance data
// func (w *World) ImportanceVal() int {
// 	if val, ok := w.data[worldImportanceEx]; ok {
// 		val = strings.TrimPrefix(val, "{")
// 		val = strings.TrimPrefix(val, " ")
// 		val = strings.TrimSuffix(val, "}")
// 		val = strings.TrimSuffix(val, " ")
// 		im, _ := strconv.Atoi(val) //TODO: Err
// 		return im
// 	}
// 	return -999
// }

//EconomicEx - return EconomicEx data
func (w *World) EconomicEx() string {
	if val, ok := w.data[worldEconomicEx]; ok {
		return val
	}
	return "--NO DATA--"
}

//CulturalEx - return CulturalEx data
func (w *World) CulturalEx() string {
	if val, ok := w.data[worldCulturalEx]; ok {
		return val
	}
	return "--NO DATA--"
}

//Nobility - return Nobility data
func (w *World) Nobility() string {
	if val, ok := w.data[worldNobility]; ok {
		return val
	}
	return "--NO DATA--"
}

//Bases - return Bases data
func (w *World) Bases() string {
	if val, ok := w.data[worldBases]; ok {
		return val
	}
	return "--NO DATA--"
}

//TravelZone - return TravelZone data
func (w *World) TravelZone() string {
	if val, ok := w.data[worldTravelZone]; ok {
		return val
	}
	return "--NO DATA--"
}

//PBG - return PBG data
func (w *World) PBG() string {
	if val, ok := w.data[worldPBG]; ok {
		return val
	}
	return "--NO DATA--"
}

//NumOfWorlds - return NumOfWorlds data
func (w *World) NumOfWorlds() string {
	if val, ok := w.data[worldNumOfWorlds]; ok {
		return val
	}
	return "--NO DATA--"
}

//Allegiance - return Allegiance data
func (w *World) Allegiance() string {
	if val, ok := w.data[worldAllegiance]; ok {
		return val
	}
	return "--NO DATA--"
}

//Stellar - return Stellar data
func (w *World) Stellar() string {
	if val, ok := w.data[worldStellar]; ok {
		return val
	}
	return "--NO DATA--"
}

//FromOTUdata -
func FromOTUdata(otuData otu.Info) (World, error) {
	w := World{}
	w.data = make(map[string]string)

	// data := strings.Split(otuData, "	")
	// //fmt.Println(otu.Info{otuData})
	// if len(data) != 17 {
	// 	return w, errors.New("OTU data unparseble: (Len != 17)")
	// }
	// info, err := otu.GetDataOn(otuData)
	// if err != nil {
	// 	return w, err
	// }
	w.data["Sector"] = otuData.Sector()
	w.data["SS"] = otuData.SubSector()
	w.data[worldHEX] = otuData.Hex()
	w.data[worldNAME] = otuData.Name()
	w.data[worldUWP] = otuData.UWP()
	w.data[worldBases] = slToStr(otuData.Bases())
	w.data[worldTradeClassifications] = slToStr(otuData.Remarks())
	w.data[worldTravelZone] = otuData.Zone()
	w.data[worldPBG] = otuData.PBG()
	w.data[worldAllegiance] = otuData.Allegiance()
	w.data[worldStellar] = otuData.Stars()
	w.data[worldImportanceEx] = otuData.Iextention()
	w.data[worldEconomicEx] = otuData.Eextention()
	w.data[worldCulturalEx] = otuData.Cextention()
	w.data[worldNobility] = otuData.Nobility()
	w.data[worldNumOfWorlds] = otuData.Worlds()
	w.data["RU"] = otuData.RU()
	w.dicepool = *dice.New(utils.SeedFromString(w.data[worldNAME]))
	return w, nil
}

func slToStr(sl []string) string {
	str := ""
	for _, val := range sl {
		str += val + " "
	}
	str = strings.TrimSuffix(str, " ")
	return str
}

func (w World) SecondSurvey() []string {
	var survey []string
	survey = append(survey, w.Hex())
	survey = append(survey, w.Name())
	survey = append(survey, w.UWP())
	survey = append(survey, slToStr(w.TradeCodes()))

	survey = append(survey, w.ImportanceEx())
	survey = append(survey, w.EconomicEx())
	survey = append(survey, w.CulturalEx())
	survey = append(survey, w.Nobility())
	survey = append(survey, w.Bases())

	survey = append(survey, w.TravelZone())
	survey = append(survey, w.PBG())
	survey = append(survey, w.NumOfWorlds())
	survey = append(survey, w.Stellar())
	return survey
}
