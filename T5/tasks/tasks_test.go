package tasks

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/T5/assets"
)

func TestTasks(t *testing.T) {

	ts := Create()
	ts.ApplyInstrucuctions(
		ts.SetDifficulty(3),
	)
	testChar := assets.NewCharacteristic(assets.Grace, 1)
	testSkill := assets.NewSkill(assets.ART_Author)
	testSkill.Train()
	testSkill.Train()
	testSkill.Train()
	testMod := NewMod("Mod1", -3)
	ts.AddAsset(testChar)
	ts.AddAsset(testMod)
	ts.AddAsset(testSkill)

	//ts.SetDifficulty(-1)
	ts.SetPurpose("apply to university")
	fmt.Println(ts.TaskPhrase())
	ts.Resolve()
	fmt.Println(ts.Outcome())
}
