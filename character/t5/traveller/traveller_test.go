package traveller

import (
	"testing"

	"github.com/Galdoba/TR_Dynasty/T5/assets"
)

func TestTraveller(t *testing.T) {
	trv := NewTravellerT5()
	if trv.err != nil {
		t.Errorf("creation Error: %v", trv.err.Error())
	}
	testSkill := assets.NewSkill(assets.SHIP_Engineer)
	trv.skills[assets.SHIP_Engineer] = testSkill
	trv.skills[assets.SHIP_Engineer].Train()
	trv.skills[assets.SHIP_Engineer].Train()
	testKnow := assets.NewKnowledge(assets.KNOWLEDGE_Chemistry)
	trv.knowledges[assets.KNOWLEDGE_Chemistry] = testKnow
	trv.knowledges[assets.KNOWLEDGE_Chemistry].Train()
	trv.knowledges[assets.KNOWLEDGE_Chemistry].Train()

	// cc := NewCard(trv)
	// fmt.Println("===TEST CARD============")
	// cc.PrintCard()
	// fmt.Println("========================")
}

func TestTravellerCard(t *testing.T) {
	//cc := NewCard(NewTravellerT5())
	// fmt.Println("===TEST CARD============")
	// cc.PrintCard()
	// fmt.Println("========================")
}
