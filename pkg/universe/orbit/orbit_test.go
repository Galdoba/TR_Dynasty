package orbit

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/universe/survey"
)

func TestPlaceWorlds(t *testing.T) {
	ss, err := survey.Search("Drinax")
	if err != nil {
		t.Errorf("search failed: %v", err.Error())
	}
	pd := PlaceWorlds(*ss[0])
	if pd == nil {
		t.Errorf("PlaceWorlds() - not implemented")
	}
	fmt.Println(ss[0])
}
