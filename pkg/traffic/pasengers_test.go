package traffic

import (
	"fmt"
	"testing"
)

func TestWorld(t *testing.T) {

	sw := Parse("Paal")
	tw := Parse("Tech-World")
	fmt.Println(sw)
	fmt.Println(tw)
	fmt.Println(Distance(sw, tw))
	pt, err := NewPassengerTrafficMgT2(sw, tw)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(pt)
	PassengerTrafficCard(pt)
}
