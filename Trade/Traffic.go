package trade

import (
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/convert"
	"github.com/Galdoba/utils"
)

//SeekPassengers - возвращает колличество пассажиров разного ранга
//в любом направлении
func SeekPassengers(planet *world.World, npcEffect int) (low int, basic int, middle int, high int) {
	dm := npcEffect + passengerDM(planet)
	lowDie := passengerTraffic(dm + 1)
	basicDie := passengerTraffic(dm)
	middleDie := passengerTraffic(dm)
	highDie := passengerTraffic(dm - 4)
	low = utils.RollDice(lowDie + "d6")
	basic = utils.RollDice(basicDie + "d6")
	middle = utils.RollDice(middleDie + "d6")
	high = utils.RollDice(highDie + "d6")
	return low, basic, middle, high
}

func passengerDM(planet *world.World) int {
	dm := 0
	stats := planet.Stats()
	if stats[planetStatPops] < 2 {
		dm = dm - 4
	}
	if stats[planetStatPops] == 6 || stats[planetStatPops] == 7 {
		dm = dm + 1
	}
	if stats[planetStatPops] > 7 {
		dm = dm + 3
	}
	switch planet.StarPort() {
	case "A":
		dm = dm + 2
	case "B":
		dm = dm + 1
	case "E":
		dm = dm - 1
	case "X":
		dm = dm - 3
	}
	if utils.ListContains(planet.TradeCodes(), travelCodeAmber) {
		dm = dm + 1
	}
	if utils.ListContains(planet.TradeCodes(), travelCodeRed) {
		dm = dm - 4
	}
	return dm
}

func passengerTraffic(dm int) string {
	r := utils.RollDice("2d6")
	r = r + dm
	res := 0
	switch r {
	default:
		if r < 2 {
			res = 0
		}
		if r > 19 {
			res = 10
		}
	case 1, 2:
		res = 1
	case 4, 5, 6:
		res = 2
	case 7, 8, 9, 10:
		res = 3
	case 11, 12, 13:
		res = 4
	case 14, 15:
		res = 5
	case 16:
		res = 6
	case 17:
		res = 7
	case 18:
		res = 8
	case 19:
		res = 9
	}
	return convert.ItoS(res)
}

//SeekFreight -
func SeekFreight(planet *world.World, npcEffect int) (major int, minor int, incedental int) {
	trafficDM := trafficDM(planet)
	majorDie := freightTraffic(trafficDM - 4)
	minorDie := freightTraffic(trafficDM)
	incedentalDie := freightTraffic(trafficDM + 2)
	major = utils.RollDice(majorDie + "d6")
	minor = utils.RollDice(minorDie + "d6")
	incedental = utils.RollDice(incedentalDie + "d6")
	return major, minor, incedental
}

func trafficDM(planet *world.World) int {
	stats := planet.Stats()
	dm := 0
	if stats[planetStatPops] < 2 {
		dm = dm - 4
	}
	if utils.InRange(stats[planetStatPops], 6, 7) {
		dm = dm + 2
	}
	if stats[planetStatPops] > 7 {
		dm = dm + 4
	}
	switch planet.StarPort() {
	case "A":
		dm = dm + 2
	case "B":
		dm = dm + 1
	case "E":
		dm = dm - 1
	case "X":
		dm = dm - 2
	}
	if stats[planetStatTech] < 7 {
		dm = dm - 1
	}
	if stats[planetStatTech] > 8 {
		dm = dm + 2
	}
	if utils.ListContains(planet.TradeCodes(), travelCodeAmber) {
		dm = dm - 2
	}
	if utils.ListContains(planet.TradeCodes(), travelCodeRed) {
		dm = dm - 6
	}
	return dm
}

func freightTraffic(dm int) string {
	r := utils.RollDice("2d6")
	r = r + dm
	res := 0
	switch r {
	default:
		if r < 2 {
			res = 0
		}
		if r > 19 {
			res = 10
		}
	case 1, 2:
		res = 1
	case 4, 5:
		res = 2
	case 6, 7, 8:
		res = 3
	case 9, 10, 11:
		res = 4
	case 12, 13, 14:
		res = 5
	case 15, 16:
		res = 6
	case 17:
		res = 7
	case 18:
		res = 8
	case 19:
		res = 9
	}
	return convert.ItoS(res)
}

func mailDM(planet *world.World, npcEffect int) int {
	dm := 0
	freightDm := trafficDM(planet)
	if freightDm < -9 {
		dm = dm - 2
	}
	if utils.InRange(freightDm, -9, -5) {
		dm = dm - 1
	}
	if utils.InRange(freightDm, 5, 9) {
		dm = dm + 1
	}
	if freightDm > 9 {
		dm = dm + 2
	}
	return dm + npcEffect
}

//SeekMail -
func SeekMail(planet *world.World, npcEffect int) (mail int) {
	mailDM := mailDM(planet, npcEffect)
	mContainers := 0
	if utils.RollDice("2d6", mailDM) > 11 {
		mContainers = utils.RollDice("d6")
	}
	return mContainers
}

/*

utils.SetSeed(15)
	world1 := worldBuilder.BuildPlanetUPP("A678689-B")
	fmt.Println(world1.UPP())
	low, basic, middle, high := TradeAnalitica.SeekPassengers(world1, 0)
	fmt.Println("")
	fmt.Println("Passenger Traffic ")
	fmt.Println("Low              : ", low)
	fmt.Println("Basic            : ", basic)
	fmt.Println("Middle           : ", middle)
	fmt.Println("High             : ", high)
	fmt.Println("")
	maj, min, inc := TradeAnalitica.SeekFreight(world1, 0)
	fmt.Println("Freight Traffic   ")
	fmt.Println("Major Lots       : ", maj)
	fmt.Println("Minor Lots       : ", min)
	fmt.Println("Incedental Lots  : ", inc)
	fmt.Println("")
	mail := TradeAnalitica.SeekMail(world1, 0)
	fmt.Println("Mail Traffic   ")
	fmt.Println("Mail Containers  : ", mail)

*/
