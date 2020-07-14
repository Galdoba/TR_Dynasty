package SpaceShip

// type LivingSpace struct {
// 	quarterStatus map[string][]string
// }

const (
	stateroomLow  = "Low"
	stateroomStd  = "Standard"
	stateroomHigh = "High"
	stateroomLux  = "Luxury"
)

// func NewLivingSpace(stateRooms [4]int) *LivingSpace {
// 	ls := &LivingSpace{}
// 	ls.quarterStatus = make(map[string][]string)
// 	for i := range stateRooms {
// 		for q := 0; q < stateRooms[i]; q++ {
// 			switch i {
// 			default:
// 				panic("Can't Create New Living Space")
// 			case 0:
// 				ls.quarterStatus[stateroomLow] = append(ls.quarterStatus[stateroomLow], "Free")
// 			case 1:
// 				ls.quarterStatus[stateroomStd] = append(ls.quarterStatus[stateroomStd], "Free")
// 			case 2:
// 				ls.quarterStatus[stateroomHigh] = append(ls.quarterStatus[stateroomHigh], "Free")
// 			case 3:
// 				ls.quarterStatus[stateroomLux] = append(ls.quarterStatus[stateroomLux], "Free")
// 			}
// 		}
// 	}
// 	return ls
// }

// /*
// В каюты вписываются экипаж и пассажиры.
// Приоритет:
// 	Экипаж
// 		Std
// 		Double Std - думать в последнюю очередь
// 		Low
// 	Пассажиры
// 		Lux
// 		High
// 		Std
// 		Low
// */

// type Assigner interface {
// 	//	Assign(string, string) bool
// 	//	Unsign(string, string) bool
// }

// type Person struct {
// 	pRole string
// }

// func NewPerson(pRole string) *Person {
// 	return &Person{pRole}
// }

// func (p *Person) quarterPriority() (priority []string) {
// 	switch p.pRole {
// 	default:
// 		priority = []string{stateroomStd, stateroomLow, stateroomHigh, stateroomLux}
// 	case "High Passenger":
// 		priority = []string{stateroomLux, stateroomHigh}
// 	case "Middle Passenger":
// 		priority = []string{stateroomStd}
// 	case "Basic Passenger":
// 		priority = []string{stateroomStd}
// 	case "Low Passenger":
// 		priority = []string{stateroomLow}
// 	}
// 	return priority
// }

// func (p *Person) canShareQuarter() bool {
// 	switch p.pRole {
// 	default:
// 		return true
// 	case "High Passenger", "Middle Passenger", "Low Passenger":
// 		return false
// 	}
// }

// func (ls *LivingSpace) Assigned(person *Person) bool {
// 	priority := person.quarterPriority()
// 	for i := range priority {
// 		if allquarters, ok := ls.quarterStatus[priority[i]]; ok {
// 			for k, quarterStatus := range allquarters {
// 				if quarterStatus == "Free" {
// 					ls.quarterStatus[priority[i]][k] = person.pRole
// 					return true
// 				}
// 				if person.canShareQuarter() && quarterStatus == person.pRole && priority[i] == stateroomStd {
// 					ls.quarterStatus[priority[i]][k] = "2x " + person.pRole
// 					return true
// 				}
// 			}
// 		}
// 	}
// 	return true
// }

// func (ls *LivingSpace) Info() {
// 	for i, val := range ls.quarterStatus[stateroomLux] {
// 		fmt.Println("Lux", i+1, val)
// 	}
// 	for i, val := range ls.quarterStatus[stateroomHigh] {
// 		fmt.Println("Hi ", i+1, val)
// 	}
// 	for i, val := range ls.quarterStatus[stateroomStd] {
// 		fmt.Println("Std", i+1, val)
// 	}
// 	for i, val := range ls.quarterStatus[stateroomLow] {
// 		fmt.Println("Low", i+1, val)
// 	}

// }
