package profile

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TR_Dynasty/T5/ehex"
)

const (
	STARPORT  = 0
	SIZE      = 1
	ATMO      = 2
	HYDRO     = 3
	POPS      = 4
	GOVR      = 5
	LAWS      = 6
	separator = 7
	TECH      = 8
)

type uwp struct {
	data map[int]ehex.Ehex
	err  error
}

type UWP interface {
	Starport() ehex.Ehex
	Size() ehex.Ehex
	Atmosphere() ehex.Ehex
	Hydrosphere() ehex.Ehex
	Population() ehex.Ehex
	Government() ehex.Ehex
	LawLevel() ehex.Ehex
	TechLevel() ehex.Ehex
	String() string
	Err() error
}

func NewUWP(data string) *uwp {
	u := uwp{}
	u.data = make(map[int]ehex.Ehex)
	data = strings.ReplaceAll(data, " ", "")
	if len(data) < 9 {
		if len(data) == 8 {
			sepData := strings.Split(data, "")
			mergedData := ""
			for i := 0; i < 7; i++ {
				mergedData += sepData[i]
			}
			mergedData += "-" + sepData[7]
		}
		u.err = fmt.Errorf("not enough data in input")
		return &u
	}
	if len(data) > 9 {
		u.err = fmt.Errorf("too many data in input")
		return &u
	}
	sepData := strings.Split(data, "")
	for i := range []int{STARPORT, SIZE, ATMO, HYDRO, POPS, GOVR, LAWS, separator, TECH} {
		switch i {
		default:
			u.err = fmt.Errorf("unknown data input")
		case STARPORT, SIZE, ATMO, HYDRO, POPS, GOVR, LAWS, TECH:
			u.data[i] = ehex.New().Set(sepData[i])
		}
	}
	return &u
}

//Starport - returns data from uwp profile
func (u *uwp) Starport() ehex.Ehex {
	return u.data[STARPORT]
}

//Size - returns data from uwp profile
func (u *uwp) Size() ehex.Ehex {
	return u.data[SIZE]
}

//Atmosphere - returns data from uwp profile
func (u *uwp) Atmosphere() ehex.Ehex {
	return u.data[ATMO]
}

//Hydrosphere - returns data from uwp profile
func (u *uwp) Hydrosphere() ehex.Ehex {
	return u.data[HYDRO]
}

//Population - returns data from uwp profile
func (u *uwp) Population() ehex.Ehex {
	return u.data[POPS]
}

//Government - returns data from uwp profile
func (u *uwp) Government() ehex.Ehex {
	return u.data[GOVR]
}

//LawLevel - returns data from uwp profile
func (u *uwp) LawLevel() ehex.Ehex {
	return u.data[LAWS]
}

//TechLevel - returns data from uwp profile
func (u *uwp) TechLevel() ehex.Ehex {
	return u.data[TECH]
}

func (u *uwp) String() string {
	return fmt.Sprintf("%v%v%v%v%v%v%v-%v",
		u.data[STARPORT],
		u.data[SIZE],
		u.data[ATMO],
		u.data[HYDRO],
		u.data[POPS],
		u.data[GOVR],
		u.data[LAWS],
		u.data[TECH],
	)
}

//Err - return uwp creation error
func (u *uwp) Err() error {
	return u.err
}
