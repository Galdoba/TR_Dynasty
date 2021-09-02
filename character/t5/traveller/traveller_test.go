package traveller

import "testing"

func TestTraveller(t *testing.T) {
	trv := NewTravellerT5()
	if trv.err != nil {
		t.Errorf("creation Error: %v", trv.err.Error())
	}
}

func TestTravellerCard(t *testing.T) {
	trv := NewTravellerT5()
	trv.PrintCard()
}
