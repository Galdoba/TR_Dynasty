package survey

import (
	"fmt"
	"testing"

	"github.com/Galdoba/utils"
)

func TestParcing(t *testing.T) {
	// f, err := os.Create("c:\\Users\\Public\\TrvData\\cleanedData.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	wwritenn := 0
	lines := utils.LinesFromTXT("c:\\Users\\Public\\TrvData\\cleanedData.txt")
	lenLines := len(lines) - 2
	errFound := 0
	errMap := make(map[string]int)
	dataMap := make(map[string]int)
	toTable := []*SecondSurveyData{}
	for i, input := range lines {

		fmt.Printf("checking world data: %v/%v (errors found: %v) - worlds Written: %v\r", i-1, lenLines, errFound, wwritenn)

		if i < 2 {
			continue
		}

		ssd := Parse(input)
		dataMap[ssd.Allegiance]++
		//block := true
		if !ssd.containsErrors() { //&& !block {
			wwritenn++
			//cleaned := strings.ReplaceAll(ssd.String(), "   ", "|")
			//f.WriteString("|" + cleaned + "\n")

			if ssd.Allegiance != "XXXX" {
				toTable = append(toTable, ssd)
			}
		}
		//errFound++
		//fmt.Println(ssd)
		//dataMap[ssd.Allegiance]++
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

	// for k, v := range dataMap {
	// 	//if v > 0 {
	// 	fmt.Println(k, ":", v)
	// 	//}
	// }

	fmt.Println("\n----------------------------------------")
	for _, lns := range ListOf(toTable) {
		fmt.Println(lns)
	}

}
