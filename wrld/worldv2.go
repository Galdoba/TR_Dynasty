package wrld

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Galdoba/TR_Dynasty/dice"
	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/profile"
	"github.com/Galdoba/devtools/cli/user"
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

//CodePops - return UWP data
func (w *World) CodePops() string {
	if val, ok := w.data[worldUWP]; ok {
		uwp := profile.NewUWP(val)
		// if err != nil {
		// 	return err.Error()
		// }
		return uwp.Pops().String()
	}
	return "--NO DATA--"
}

//CodeTL - return UWP data
func (w *World) CodeTL() string {
	if val, ok := w.data[worldUWP]; ok {
		uwp := profile.NewUWP(val)
		// if err != nil {
		// 	return err.Error()
		// }
		return uwp.TL().String()
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

//ImportanceVal - return Importance As Int
func (w *World) ImportanceVal() int {
	if val, ok := w.data[worldImportanceEx]; ok {
		val = strings.TrimPrefix(val, "{")
		val = strings.TrimPrefix(val, " ")
		val = strings.TrimSuffix(val, "}")
		val = strings.TrimSuffix(val, " ")
		im, _ := strconv.Atoi(val) //TODO: Err
		return im
	}
	return -999
}

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
	w.checkLtHtTradeCodes()
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
	w.dicepool = *dice.New().SetSeed(w.data[worldNAME] + w.data[worldNAME])
	return w, nil
}

func slToStr(sl []string) string {
	str := ""
	for _, val := range sl {
		if strings.Contains(val, "O:") {
			continue //Пока исключаем по эстетическим причинам
			//потом решить что делать с ремарками типа "O:2324"
		}
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

func (w *World) checkLtHtTradeCodes() {
	uwp := profile.NewUWP(w.UWP())
	if uwp.TL().Value() <= 5 {
		w.data[worldTradeClassifications] += " Lt "
	}
	if uwp.TL().Value() >= 12 {
		w.data[worldTradeClassifications] += " Ht "
	}

	// if TrvCore.EhexToDigit(w.CodeTL()) <= 5 {
	// 	w.data[worldTradeClassifications] += " Lt "
	// }
	// if TrvCore.EhexToDigit(w.CodeTL()) >= 12 {
	// 	w.data[worldTradeClassifications] += " Ht "
	// }

	// if TrvCore.EhexToDigit(w.data[constant.PrTL]) <= TrvCore.EhexToDigit("5") {
	// 	fmt.Println(TrvCore.EhexToDigit(w.data[constant.PrTL]))
	// 	fmt.Println(w.data[constant.PrTL])
	// 	panic(1)
	// 	w.data[worldTradeClassifications] += " Lt "
	// }
	// if TrvCore.EhexToDigit(w.data[constant.PrTL]) >= TrvCore.EhexToDigit("C") {
	// 	w.data[worldTradeClassifications] += " Ht "
	// }
	w.data[worldTradeClassifications] = strings.TrimSuffix(w.data[worldTradeClassifications], " ")
}

func PickWorld() World {
	dataFound := false
	for !dataFound {
		fmt.Print("Enter world's Name, Hex or UWP: ")
		input, err := user.InputStr()
		if err != nil {

		}
		data, err := otu.GetDataOn(input)
		if err != nil {
			fmt.Print("WARNING: " + err.Error() + "\n")
			continue
		}
		w, err := FromOTUdata(data)
		if err != nil {
			fmt.Print(err.Error() + "\n")
			continue
		}
		return w
	}
	fmt.Println("This must not happen!")
	return World{}
}
