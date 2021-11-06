package starport

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

func TestAsemble(t *testing.T) {
	worldData, err := otu.GetDataOn("Acis")
	if err != nil {
		t.Errorf(err.Error())
	}
	world, err := wrld.FromOTUdata(worldData)
	if err != nil {
		t.Errorf(err.Error())
	}
	sp, err := Assemble(&world)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(sp)
}
