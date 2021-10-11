package starsystem

import (
	"fmt"
	"strings"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
)

//RollStellar - выдает рандобный набор звезд в формате Second Survey
func RollStellar(seed ...string) string {
	d := dice.New()
	if len(seed) > 0 {
		d.SetSeed(seed[0])
	}
	stellar := ""
	prim := make(map[string][]string)
	prim["Type"] = []string{"B", "A", "A", "F", "F", "G", "G", "K", "K", "M", "M", "M", "BD", "BD", "BD"}
	prim["O"] = []string{"Ia", "Ia", "Ib", "II", "III", "III", "III", "V", "V", "V", "IV", "D", "IV", "IV", "IV"}
	prim["B"] = []string{"Ia", "Ia", "Ib", "II", "III", "III", "III", "III", "V", "V", "IV", "D", "IV", "IV", "IV"}
	prim["A"] = []string{"Ia", "Ia", "Ib", "II", "III", "IV", "V", "V", "V", "V", "V", "D", "V", "V", "V"}
	prim["F"] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	prim["G"] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	prim["K"] = []string{"II", "II", "III", "IV", "V", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	prim["M"] = []string{"II", "II", "II", "II", "III", "V", "V", "V", "V", "V", "VI", "D", "VI", "VI", "VI"}
	starType := []string{"P"}
	if d.FluxNext() > 2 {
		starType = append(starType, "Pc")
	}
	if d.FluxNext() > 2 {
		starType = append(starType, "C")
		if d.FluxNext() > 2 {
			starType = append(starType, "Cc")
		}
	}
	if d.FluxNext() > 2 {
		starType = append(starType, "N")
		if d.FluxNext() > 2 {
			starType = append(starType, "Nc")
		}
	}
	if d.FluxNext() > 2 {
		starType = append(starType, "F")
		if d.FluxNext() > 2 {
			starType = append(starType, "Fc")
		}
	}
	for _, val := range starType {
		tpRoll := d.FluxNext()
		if val != "P" {
			tpRoll += d.RollNext("1d6").DM(-1).Sum()
		}
		if tpRoll < -6 {
			tpRoll = -6
		}
		if tpRoll > 8 {
			tpRoll = 8
		}
		szRoll := d.FluxNext()
		if val != "P" {
			szRoll += d.RollNext("1d6").DM(-1).Sum()
		}
		if szRoll < -6 {
			szRoll = -6
		}
		if szRoll > 8 {
			szRoll = 8
		}
		decRoll := d.RollNext("1d10").DM(-1).Sum()
		strType := ""
		strType += prim["Type"][tpRoll+6]
		if strType == "BD" {
			stellar = " " + strType
			continue
		}
		star := ""
		sz := prim[strType][szRoll+6]
		if sz == "D" {
			star = strType + "D"
			stellar = " " + star
			continue
		}
		star = fmt.Sprintf("%v%v %v", strType, decRoll, sz)
		stellar += " " + star
	}
	stellar = strings.TrimPrefix(stellar, " ")
	return stellar

}

/*

stars := starsystem.Stars(worldData SecondSurvey) // []string

formatedData => First Survey
formatedData => Second Survey




*/
