package npcmakerv2

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/entity"
)

func Test() {
	fmt.Println("Start test")
	trv := Traveller{
		make(map[string]TrvCore.Ehex),
		entity.NewSkillMap(),
	}
	trv.skills.Train("Admin")
	//fmt.Println(trv.skills["Admin"])
	fmt.Println("End test")
}

type Traveller struct {
	attributes map[string]TrvCore.Ehex
	skills     entity.Skill
}

/*
НПС это сущность
Интерфейсы:
-делать проверки навыка (SkillTester)
-делать проверки характеристики (AttributeTester)
-рассказывать о своих навыках (SkillGiver)
-рассказывать о своих характеристиках (AtribbuteGiver)
-рассказывать о себе (Describer)

*/
