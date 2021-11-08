package law

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/T5/ehex"
	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/profile"
	"github.com/Galdoba/TR_Dynasty/pkg/universe/survey"
)

type worldData struct {
	name string
	uwp  string
}

type SecurityStatus struct {
	name string
	uwp  profile.UWP
	laws []int //Overall/separator/Weapon/Drugs/Info/Tech/Trav
}

type WorldSecurityData interface {
}

func Parse(ssd *survey.SecondSurveyData) worldData {
	w := worldData{}
	w.name = ssd.MW_Name
	w.uwp = ssd.MW_UWP
	return w
}

func NewSecurityStatus(world worldData) (*SecurityStatus, error) {
	wsd := SecurityStatus{}
	wsd.uwp = profile.NewUWP(world.uwp)
	law := wsd.uwp.LawLevel().Value()
	dp := dice.New().SetSeed(world.name + world.uwp)
	for i := 0; i < 7; i++ {
		switch i {
		case 0:
			wsd.laws = append(wsd.laws, law)
		case 1:
			wsd.laws = append(wsd.laws, ehex.New().Set("-").Value())
		default:
			l := dp.FluxNext() + law
			if l < 0 {
				l = 0
			}
			wsd.laws = append(wsd.laws, l)
		}
	}
	return &wsd, nil
}

func (ss *SecurityStatus) LawsString() string {
	str := ""
	for i, v := range ss.laws {
		if i == 1 {
			str += "-"
			continue
		}
		str += ehex.New().Set(v).String()
	}
	return str
}

func (ss *SecurityStatus) SecurityStatusCard() {
	fmt.Printf("Planet %v (%v):\n", ss.name, ss.uwp.String())
	fmt.Printf("Universal Law Profile: %v\n", ss.LawsString())
	switch dice.Roll2D() <= ss.laws[0] {
	case false:
		fmt.Printf("Security check avoided\n")
	case true:
		fmt.Printf("Security check initiated\n")
		fmt.Printf("Succesful Admin or Streetwise check must be passed to avoid investigation\n")
	}
}

/*
Расследование:
Sensor - Сканирование
Infestigate - Физический поиск
Streetwise - Допрос
проверки идут со сложностью 8
каждая успешная проверка бросает 1д6 дней - берем максимальный результат
уровень закона и ТЛ на планете используем как модификаторы характеристики - высоко технологичные миры с высоким уровнем закона будут легче проходить проверки
*/
