package traveller

import (
	"testing"
)

func TestTraveller(t *testing.T) {
	trv := NewTravellerT5()
	if trv.err != nil {
		t.Errorf("creation Error: %v", trv.err.Error())
	}
}

func TestTravellerCard(t *testing.T) {
	//cc := NewCard(NewTravellerT5())
	// fmt.Println("===TEST CARD============")
	// cc.PrintCard()
	// fmt.Println("========================")
}
