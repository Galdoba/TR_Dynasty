package law

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/universe/survey"
)

func TestLaw(t *testing.T) {
	ssd, _ := survey.Search("Paal")
	world := Parse(ssd[0])
	ls, err := NewSecurityStatus(world)
	fmt.Println(ls)
	fmt.Println("Long String:", ls.LawsString())
	fmt.Println(err)
	ls.SecurityStatusCard()
}
