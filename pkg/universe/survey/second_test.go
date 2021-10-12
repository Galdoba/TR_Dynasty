package survey

import (
	"fmt"
	"testing"

	"github.com/Galdoba/utils"
)

func TestParcing(t *testing.T) {

	lines := utils.LinesFromTXT("c:\\Users\\Public\\TrvData\\formattedData.txt")
	//lenLines := len(lines) - 2
	errFound := 0
	errMap := make(map[string]int)
	dataMap := make(map[string]int)
	for i, input := range lines {
		//fmt.Printf("checking world data: %v/%v (errors found: %v)\r", i-1, lenLines, errFound)
		if i < 2 {
			continue
		}

		ssd := Parse(input)
		dataMap[ssd.MW_UWP]++
		if !ssd.containsErrors() {
			continue
		}
		errFound++
		//fmt.Println(ssd)
		for _, err := range ssd.errors {
			if err != nil {
				//fmt.Println(err.Error())
				errFound++
				errMap[err.Error()]++
			}
		}

		// if i > 29480 {
		// 	return
		// }
	}
	fmt.Println("\n----------------------------------------")
	for k, v := range errMap {
		//if !calculations.UWPvalid(k) {
		fmt.Println(k, ":", v)
		//}
	}
	fmt.Println("\n----------------------------------------")

}