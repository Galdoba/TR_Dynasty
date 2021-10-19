package survey

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Galdoba/utils"
)

func TestParcing(t *testing.T) {
	f, err := os.Create("c:\\Users\\Public\\TrvData\\cleanedData.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	wwritenn := 0
	lines := utils.LinesFromTXT("c:\\Users\\Public\\TrvData\\formattedData.txt")
	lenLines := len(lines) - 2
	errFound := 0
	errMap := make(map[string]int)
	dataMap := make(map[string]int)
	toTable := []*SecondSurveyData{}
	for i, input := range lines {
		if dataMap[input] != 0 {
			errFound++
			errMap["Duplicated input"]++
			continue
		}
		fmt.Printf("checking world data: %v/%v (errors found: %v) - worlds Written: %v\r", i-1, lenLines, errFound, wwritenn)

		if i < 2 {
			continue
		}

		ssd := Parse(input)
		dataMap[input]++
		if !ssd.containsErrors() {
			wwritenn++
			cleaned := strings.ReplaceAll(ssd.String(), "   ", "|")
			f.WriteString("|" + cleaned + "\n")
			continue
		}
		//errFound++
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
		if v > 1 {
			fmt.Println(k, ":", v)
		}
	}

	fmt.Println("\n----------------------------------------")
	for _, lns := range ListOf(toTable) {
		fmt.Println(lns)
	}

}
