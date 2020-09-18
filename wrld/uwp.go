package wrld

import (
	core "github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
)

// //UWP - механика контроля данных в профайле планеты
// type UWPer interface {
// 	UWP() string
// 	GetСharacteristic(string) core.Ehex
// }

// type uwp struct { //Плодим сущности?
// 	data string
// }

// func UWPFrom(w World) uwp {
// 	return uwp{w.data[worldUWP]}
// }

//GetСharacteristic - возвращает характеристику мира или nil если х-ка не найдена
//поидее это часть Profiler
//GetData - TODO: подумать над тем чтобы сделать единый Геттер
func (w World) GetСharacteristic(characteristic string) core.Ehex {
	bt := -1
	switch characteristic {
	default:
		return nil
	case constant.PrStarport:
		bt = 0
	case constant.PrSize:
		bt = 1
	case constant.PrAtmo:
		bt = 2
	case constant.PrHydr:
		bt = 3
	case constant.PrPops:
		bt = 4
	case constant.PrGovr:
		bt = 5
	case constant.PrLaws:
		bt = 6
	case constant.PrTL:
		bt = 8
	}
	v := string(w.UWP()[bt])
	gl := core.EhexFromStr(v)
	return &gl
}
