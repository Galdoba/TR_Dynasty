package mission

import (
	"fmt"
	"testing"

	"github.com/Galdoba/TR_Dynasty/pkg/cei"
)

func Test_Mission(t *testing.T) {
	crew := cei.New(7)

	seg := NewSegment(crew, "Search planetside location",
		NewOperation(crew, "Land Shuttle 2"),
		NewOperation(crew, "Recon Area 2"),
		NewOperation(crew, "Attack Vilage 2"),
		NewOperation(crew, "Extract Force 2"),
	)
	seg.Resolve()
	fmt.Println(seg.Report())

}
