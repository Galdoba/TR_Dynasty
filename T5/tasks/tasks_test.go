package tasks

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/T5/assets"
)

func TestTasks(t *testing.T) {

	ts := Create()
	ts.SetDifficulty(2)
	testChar := assets.NewCharacteristic(assets.Education, 1)
	testSkill := assets.NewSkill(assets.ART_Author)
	testSkill.Train()
	testSkill.Train()
	testSkill.Train()
	testMod := NewMod("Mod1", 3)
	ts.AddAsset(testChar)
	ts.AddAsset(testMod)
	ts.AddAsset(testSkill)

	//ts.SetDifficulty(-1)
	ts.SetPurpose("apply to university")
	fmt.Print(ts.TaskPhrase())
	ts.Resolve()
	fmt.Println(ts.Outcome())

	ts2 := Create()
	ts2.SetupEnviroment("test setupfunctions", 2, 2, TASK_COMMENT_Cautious)
	tChar := assets.NewCharacteristic(assets.Charisma, 5)
	tSkil := assets.NewSkill(assets.SKILL_Broker)
	tMods := NewMod("War", 2)
	ts2.SetupAssets(tChar, tSkil, tMods)
	ts2.Resolve()
	fmt.Println(ts2.TaskPhrase())
	fmt.Println(ts2.Outcome())

}
