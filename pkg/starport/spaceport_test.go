package starport

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/otu"
	"github.com/Galdoba/TR_Dynasty/wrld"
)

func TestAsemble(t *testing.T) {
	worldData, err := otu.GetDataOn("Paal")
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

	worldData2, err := otu.GetDataOn("Tktk")
	if err != nil {
		t.Errorf(err.Error())
	}
	world2, err := wrld.FromOTUdata(worldData2)
	if err != nil {
		t.Errorf(err.Error())
	}
	sp2, err := Assemble(&world2)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(sp2)
}
