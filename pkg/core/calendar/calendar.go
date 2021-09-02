package calendar

import (
	"fmt"
	"strconv"
)

const (
	WeekDay_Holiday = "Holiday"
	WeekDay_Wonday  = "Wonday"
	WeekDay_Tuday   = "Tuday"
	WeekDay_Thirday = "Thirday"
	WeekDay_Forday  = "Forday"
	WeekDay_Fiday   = "Fiday"
	WeekDay_Sixday  = "Sixday"
	WeekDay_Senday  = "Senday"
	GameStartDay    = 403326
	Period_Week     = 7
	Period_Year     = 365
)

type ImperialDate struct {
	day     int
	year    int
	weekDay string
	err     error
}

func NewImperialDate(data ...int) *ImperialDate {
	day := 0
	if len(data) > 0 {
		day = data[0]
	}
	year := 0
	if len(data) > 1 {
		year = data[1]
	}
	d := ImperialDate{}
	d.err = fmt.Errorf("Initial")
	d.day = day
	d.year = year
	d.validate()
	return &d
}

func (d *ImperialDate) validate() {
	for d.day < 1 {
		d.day += 365
		d.year--
		if d.year == 0 {
			d.year--
		}
	}
	for d.day > 365 {
		d.day -= 365
		d.year++
		if d.year == 0 {
			d.year++
		}
	}
	if d.year == 0 {
		d.year++
	}
	d.weekDay = currentWeekDay(d.day)
	d.err = nil
}

func (d *ImperialDate) String() string {
	str := strconv.Itoa(d.day) + "-"
	if d.day < 100 {
		str = "0" + str
	}
	if d.day < 10 {
		str = "0" + str
	}
	switch {
	case d.year < 0:
		str = str + strconv.Itoa(d.year*-1) + " BI"
	case d.year >= 0:
		str = str + strconv.Itoa(d.year)
	}
	return str
}

func currentWeekDay(day int) string {
	if day == 1 {
		return WeekDay_Holiday
	}
	i := (day - 1) % 7
	if i == 0 {
		i = 7
		return weekDays()[i]
	}
	return weekDays()[i]
}

func weekDays() []string {
	return []string{
		WeekDay_Holiday,
		WeekDay_Wonday,
		WeekDay_Tuday,
		WeekDay_Thirday,
		WeekDay_Forday,
		WeekDay_Fiday,
		WeekDay_Sixday,
		WeekDay_Senday,
	}
}

func (d *ImperialDate) MoveDateByDays(days int) {
	d.day = d.day + days
	d.validate()
}

func (d *ImperialDate) MoveDateByYears(years int) {
	d.year = d.year + years
	d.validate()
}

func ProjectDate(date *ImperialDate, period int) *ImperialDate {
	newDate := NewImperialDate(date.day+period, date.year)
	return newDate
}
