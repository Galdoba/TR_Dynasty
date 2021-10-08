package survey

import (
	"fmt"
	"testing"

	"github.com/Galdoba/utils"
)

func TestParcing(t *testing.T) {
	lines := utils.LinesFromTXT("c:\\Users\\Public\\TrvData\\formattedData.txt")
	for i, input := range lines {
		if i < 29460 && i != 0 {
			continue
		}
		ssd := Parse(input)
		fmt.Println(ssd)

		if i > 29480 {
			return
		}
	}
}
