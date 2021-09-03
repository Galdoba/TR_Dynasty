package tasks

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/T5/assets"
	"github.com/Galdoba/TR_Dynasty/character/t5/traveller"
)

func TestTasks(t *testing.T) {

	trv := traveller.NewTravellerT5()
	cc := traveller.NewCard(trv)
	cc.PrintCard()
	fmt.Println("=======")
	ts := Create()
	ts.AddAsset(trv.Characteristic(assets.Intelligence))
	testSkill := assets.NewSkill(assets.ART_Author)
	testSkill.Train()
	testSkill.Train()
	testSkill.Train()
	testMod := NewMod("Enviroment", -5)
	ts.AddAsset(testMod)
	ts.AddAsset(testSkill)

	ts.SetDifficulty(2)
	ts.SetPurpose("apply to university")
	fmt.Println(ts.TaskPhrase())
	ts.Resolve()
	fmt.Println(ts.Outcome())
}
