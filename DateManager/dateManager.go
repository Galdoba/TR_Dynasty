package DateManager

import (
	"errors"
	"strconv"
	"strings"
)

const (
	tFrameSeconds     = "Second"
	tFrameRoundCombat = "Combat Round"
	tFrameSeconds10   = "10 Seconds"
	tFrameMinutes     = "Minute"
	tFrameRoundSpace  = "Space Combat Round"
	tFrameMinutes10   = "10 Minutes"
	tFrameHours       = "Hour"
	tFrameHours4      = "4 Hours"
	tFrameHours10     = "10 Hours"
	tFrameDays        = "Day"
	tFrameDays4       = "4 Days"
	tFrameWeeks       = "Week"
	tFrameWeeks2      = "2 Weeks"
	tFrameMonths      = "Month"
	tFrameMonths6     = "6 Months"
	tFrameYears       = "Year"
)

// func main() {
// 	//date1 := NewImperialDate(1105, 54, 14, 3, 21)
// 	date := DateFromString("001--2  23:16:57")
// 	fmt.Println(date.ToString())
// 	for i := 0; i < 15; i++ {
// 		date.PassTime(tFrameYears, utils.RollDice("2d6")-1)
// 		fmt.Println(date.ToString())
// 	}
// }

/*
имперский год состоит из 365 дней по 24 часа.
месяцы и дни недели фиксированы
количество недель идет по схеме 5-4-4-5-4-4-5-4-4-5-4-4 (первый месяц квартала имеет 5 недель)
первое число года не является днем недели и называется "праздник"
исчисление идет с точки зрения Астрономии - см. статью о астрономическом летоисчислении
нулевой год был
*/

type Date struct {
	day int
}

func FormatToDate(d int) string {
	//чтобы не тормозить
	//TODO: выяснить б
	day := strconv.Itoa(d % 365)
	year := strconv.Itoa(d / 365)
	postFix := " AFE"
	if d <= 0 {
		day = strconv.Itoa(365 + (d % 365))
		year = strconv.Itoa(-1 * (d / 365))
		postFix = " BFE"
	}
	for len(day) < 3 {
		day = "0" + day
	}
	for len(year) < 4 {
		year = "0" + year
	}
	return day + "-" + year + postFix
}

func FormatToDay(timestamp string) (int, error) {
	data := strings.Split(timestamp, " ")
	if len(data) != 2 {
		return 0, errors.New("Wrong data format '" + timestamp + "'")
	}
	dy := strings.Split(data[0], "-")
	if len(dy) != 2 {
		return 0, errors.New("Wrong data format '" + timestamp + "'")
	}
	day, errD := strconv.Atoi(dy[0])
	if errD != nil {
		return 0, errors.New("Wrong data format '" + timestamp + "' - " + dy[0])
	}
	year, errY := strconv.Atoi(dy[1])
	if errY != nil {
		return 0, errors.New("Wrong data format '" + timestamp + "' - " + dy[1])
	}
	if data[1] == "BFE" {
		day = day // - 365
		year = (year * -1) - 1
	}
	return day + (year * 365), nil
}

//ImperialDate - дата и время (должно стать частью "События")

// type ImperialDate struct {
// 	timeVal int64
// 	//YYYY:DDD:HH:MM:SS
// 	//DCBA 987 65 43 21
// }

// func NewImperialDate(timeVal int64) *ImperialDate {
// 	iDate := &ImperialDate{}
// 	iDate.timeVal = timeVal
// 	iDate.validate()
// 	return iDate
// }

// func disassembleTimeval(timeval int64) (int64, int64, int64, int64, int64) {
// 	ss := timeval % 100
// 	mm := timeval % 10000 / 100
// 	hh := timeval % 1000000 / 10000
// 	ddd := timeval % 1000000000 / 1000000
// 	yyyy := timeval /*% 10000000000000000*/ / 1000000000
// 	return ss, mm, hh, ddd, yyyy
// }

// func assembleTimeval(ss, mm, hh, ddd, yyyy int64) (timeVal int64) {
// 	return ss + mm*100 + hh*10000 + ddd*1000000 + yyyy*1000000000
// }

// func (iDate *ImperialDate) validate() {
// 	//fmt.Print("\033[1A", iDate.timeVal, " ")
// 	ss, mm, hh, ddd, yyyy := disassembleTimeval(iDate.timeVal)
// 	for ss > 59 {
// 		ss = ss - 60
// 		mm++
// 	}
// 	for mm > 59 {
// 		mm = mm - 60
// 		hh++
// 	}
// 	for hh > 23 {
// 		hh = hh - 24
// 		ddd++
// 	}
// 	// if ddd == 0 {
// 	// 	ddd = 1
// 	// }
// 	for ddd > 365 {
// 		ddd = ddd - 365
// 		yyyy++
// 	}
// 	iDate.timeVal = assembleTimeval(ss, mm, hh, ddd, yyyy)

// }

// func (iDate *ImperialDate) PassTime(timeVal int64) {
// 	iDate.timeVal = iDate.timeVal + timeVal
// 	iDate.validate()
// }

// func (iDate *ImperialDate) Time() string {
// 	ss, mm, hh, _, _ := disassembleTimeval(iDate.timeVal)
// 	sec := convert.ItoS(int(ss))
// 	if ss < 10 {
// 		sec = "0" + sec
// 	}
// 	min := convert.ItoS(int(mm))
// 	if mm < 10 {
// 		min = "0" + min
// 	}
// 	hour := convert.ItoS(int(hh))
// 	if hh < 10 {
// 		hour = "0" + hour
// 	}
// 	return hour + ":" + min + ":" + sec
// }

// func (iDate *ImperialDate) Date() string {
// 	_, _, _, ddd, yyyy := disassembleTimeval(iDate.timeVal)
// 	day := convert.ItoS(int(ddd))
// 	if ddd < 100 {
// 		day = "0" + day
// 	}
// 	if ddd < 10 {
// 		day = "0" + day
// 	}
// 	year := convert.ItoS(int(yyyy))
// 	if yyyy < 1000 {
// 		year = "0" + year
// 	}
// 	if yyyy < 100 {
// 		year = "0" + year
// 	}
// 	if yyyy < 10 {
// 		year = "0" + year
// 	}
// 	if yyyy < 0 {
// 		year = convert.ItoS(int(yyyy))
// 	}
// 	return day + "-" + year
// }

// func (iDate *ImperialDate) TimeStamp() string {
// 	return iDate.Date() + " " + iDate.Time()
// }

// func TimeframeSecond() (ss int64) {
// 	ss = 1
// 	return ss
// }

// func TimeframeMinute() (mm int64) {
// 	mm = 100
// 	return mm
// }

// func TimeframeHour() (hh int64) {
// 	hh = 10000
// 	return hh
// }

// func TimeframeDay() (ddd int64) {
// 	ddd = 1000000
// 	return ddd
// }

// func TimeframeWeek() (ddd int64) {
// 	ddd = 1000000
// 	return ddd * 7
// }

// func TimeframeYear() (yyyy int64) {
// 	yyyy = 1000000000
// 	return yyyy
// }

// //DateManager - управляет датой
// type DateManager interface {
// 	PassTime(int64)
// }

// func (iDate *ImperialDate) TimeValReached(timeVal int64) bool {
// 	if iDate.timeVal >= timeVal {
// 		return true
// 	}
// 	return false
// }

// func SecondsToTimeval(seconds int64) int64 {
// 	ss := seconds % 60
// 	mm := seconds / 60
// 	hh := mm / 60
// 	ddd := hh / 24
// 	yyyy := ddd / 365
// 	timeVal := assembleTimeval(ss, mm, hh, ddd, yyyy)
// 	return timeVal
// }

// func TimeToString(timeVal int64) string {
// 	date := NewImperialDate(timeVal)
// 	time := date.TimeStamp()
// 	return "DEBUG: " + time
// }

// func TimeToHuman(hr float64) string {
// 	hours := int(hr)
// 	minutes := int((hr - float64(hours)) * 60)
// 	days := hours / 24
// 	rep := ""
// 	if days == 1 {
// 		rep = "1 day "
// 	}
// 	if days > 1 {
// 		rep = strconv.Itoa(days) + " days "
// 	}
// 	rep = rep + strconv.Itoa(hours) + " hours " + strconv.Itoa(minutes) + " minutes"
// 	return rep
// }
