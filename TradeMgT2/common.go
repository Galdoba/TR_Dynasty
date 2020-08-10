package trademgt2

import "fmt"

var PassageMap map[int]int

const (
	passageHigh   = 0
	passageMiddle = 1
	passageBasic  = 2
	passageLow    = 3
	freight       = 4
	j1            = 10
	j2            = 20
	j3            = 30
	j4            = 40
	j5            = 50
	j6            = 60
)

func Init() {
	PassageCostMap := make(map[int]int)
	//J1 Passage
	PassageCostMap[j1+passageHigh] = 8500
	PassageCostMap[j2+passageHigh] = 12000
	PassageCostMap[j3+passageHigh] = 20000
	PassageCostMap[j4+passageHigh] = 41000
	PassageCostMap[j5+passageHigh] = 45000
	PassageCostMap[j6+passageHigh] = 47000
	PassageCostMap[j1+passageMiddle] = 6200
	PassageCostMap[j2+passageMiddle] = 9000
	PassageCostMap[j3+passageMiddle] = 15000
	PassageCostMap[j4+passageMiddle] = 31000
	PassageCostMap[j5+passageMiddle] = 34000
	PassageCostMap[j6+passageMiddle] = 350000
	PassageCostMap[j1+passageBasic] = 2200
	PassageCostMap[j2+passageBasic] = 2900
	PassageCostMap[j3+passageBasic] = 4400
	PassageCostMap[j4+passageBasic] = 8600
	PassageCostMap[j5+passageBasic] = 9400
	PassageCostMap[j6+passageBasic] = 93000
	PassageCostMap[j1+passageLow] = 700
	PassageCostMap[j2+passageLow] = 1300
	PassageCostMap[j3+passageLow] = 2200
	PassageCostMap[j4+passageLow] = 4300
	PassageCostMap[j5+passageLow] = 13000
	PassageCostMap[j6+passageLow] = 96000
	PassageCostMap[j1+freight] = 1000
	PassageCostMap[j2+freight] = 1600
	PassageCostMap[j3+freight] = 3000
	PassageCostMap[j4+freight] = 7000
	PassageCostMap[j5+freight] = 7700
	PassageCostMap[j6+freight] = 86000
	PassageMap = PassageCostMap
}

func PassageFreightCost(passageType int, route []int) (sum int) {
	fmt.Println(PassageMap)
	for i := range route {
		fmt.Println(sum, route[i], passageType, route[i]*10+passageType, PassageMap[route[i]*10+passageType])
		sum += PassageMap[route[i]*10+passageType]
	}
	return sum
}
