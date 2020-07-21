package world

import (
	"strconv"

	"github.com/Galdoba/TR_Dynasty/DateManager"
	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/constant"
	"github.com/Galdoba/utils"
)

//TODO: подумать как влить в механику данные от карты планеты
type weather struct {
	wx       int
	climate  int
	season   int
	rotation int
	latitude int
	desc     string
}

func newWeather(w World) weather {
	we := weather{}
	a := TrvCore.EhexToDigit(w.data[constant.PrAtmo])
	h := TrvCore.EhexToDigit(w.data[constant.PrHydr])
	s := TrvCore.EhexToDigit(w.data[constant.PrSize])
	if s == 0 {
		we.wx = 2
		return we
	}
	we.wx = (a * h) / s
	if we.wx < 2 {
		we.wx = 2
	}
	return we
}

func (we weather) setClimate(cl int) {
	we.climate = cl
}

func (we weather) setSeason(se int) {
	we.season = se
}
func (we weather) setRotation(ro int) {
	we.rotation = ro
}
func (we weather) setLatitude(la int) {
	we.latitude = la
}

//Weather - Возвращает описание текущей погоды набрасываемое
//чистым рандомом
func Weather(w World) string {
	we := newWeather(w)
	r := utils.RollDiceRandom("d6", we.wx)
	test := utils.RollDiceRandom("2d6")
	report := "Weather Score: " + strconv.Itoa(we.wx) + " roll:" + strconv.Itoa(test) + "\n"
	if test > we.wx {
		return report + "\nNo Activity until end of the day"
	}
	for i := r; i >= 0; i-- {
		report = report + we.forcast(i) + "\n"
	}
	return report
}

func (we weather) forcast(r int) string {
	dur := 100
	cond := ""
	switch r {
	default:
		if r < 1 {
			return "No Activity until end of the day"
		}
		if r > 6 {
			cond = "Major Storm, Regional Cyclonic"
			dur = 400
		}
	case 1:
		dur = utils.RollDiceRandom("d3") * 100
		cond = "Overcast"
	case 2:
		dur = utils.RollDiceRandom("d6") * 100
		cond = "Overcast"
	case 3:
		cond = "Minor Storm, Local"
	case 4:
		cond = "Minor Storm, Regional"
	case 5:
		cond = "Major Storm, Local"
		dur = 50
	case 6:
		cond = "Violent Storm, Local"
		dur = 25

	}
	t := ((utils.RollDiceRandom("6d6") * 10) * dur) / 100
	return cond + " for next " + DateManager.TimeToHuman(float64(t)/60)
}
