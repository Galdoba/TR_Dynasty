package calendar

import (
	"fmt"
	"testing"

	"github.com/Galdoba/utils"
)

func inputFeed() [][]int {
	return [][]int{
		{4, 4},
		{-4, -4},
		{4, -4},
		{-4, 4},
		{0, 0},
		{-4, 0},
		{0, -4},
		{0, 4},
		{400, 400},
		{-400, -400},
		{400, -400},
		{-400, 400},
		{18651, 400},
		{200000, 0},
		{403326, 0},
		{403326, 0, 0},
		{},
		{-3},
		{-3, -1},
		{GameStartDay},
	}
}

func TestCalendar(t *testing.T) {
	for i, input := range inputFeed() {
		d := NewImperialDate(input...)
		switch {
		case d.err != nil:
			fmt.Println("Input", i+1, "|", input)
			t.Errorf("creation Error: %v | Input:= %v", d.err.Error(), input)
		}
		if !utils.ListContains(weekDays(), d.weekDay) {
			t.Errorf("date weekday not valid: '%v'", d.weekDay)
		}
		if d.day < 1 {
			t.Errorf("day cannot be less than 1: day='%v'", d.day)
		}
		if d.day > 365 {
			t.Errorf("day cannot be more than 365: day='%v'", d.day)
		}
		if d.year == 0 {
			t.Errorf("yar cannot be equal 0: year='%v'", d.year)
		}
	}
}

func TestMoveDateBy(t *testing.T) {
	d := NewImperialDate(0 + 350)
	for i := 0; i < 20; i++ {
		d.MoveDateByDays(1)
		d.MoveDateByYears(1)
		secondDate := ProjectDate(d, -10*Period_Year)
		switch {
		case d.err != nil:
			t.Errorf("creation Error: %v | Input:= %v", d.err.Error(), d)
		}
		if !utils.ListContains(weekDays(), d.weekDay) {
			t.Errorf("date weekday not valid: '%v'", d.weekDay)
		}
		if d.day < 1 {
			t.Errorf("day cannot be less than 1: day='%v'", d.day)
		}
		if d.day > 365 {
			t.Errorf("day cannot be more than 365: day='%v'", d.day)
		}
		if d.year == 0 {
			t.Errorf("yar cannot be equal 0: year='%v'", d.year)
		}
		fmt.Println(d, d.weekDay, "|", secondDate)
	}
}
