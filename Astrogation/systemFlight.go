package Astrogation

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Galdoba/TR_Dynasty/TrvCore"
	"github.com/Galdoba/TR_Dynasty/world"
	"github.com/Galdoba/utils"
)

type Astrogator interface {
	Test() string
}

type astrogatorObj struct {
	objType string
}

type planetaryBody interface {
	JumpPointDistance() float64
}

/*
Astrogator.CalculateManuver([start point], [end point], [manuverType], [thrust]) (time) - сложно((

*/

func (ast *astrogatorObj) Test() string {
	return "I am Astrogator"
}

func NewAstrogator() *astrogatorObj {
	astr := &astrogatorObj{}
	return astr
}

//JumpPointToOrbit - возвращает расстояние между точкой прыжка (100D) и стандартной орбитой (10D)
//расстояние измеряется в МегаМетрах (1 Мм = 1000 Км)
func JumpPointToOrbit(planetaryDiametrKM int) float64 {
	return float64(planetaryDiametrKM/1000) * 90
}

//JumpPointFromObject - возвращает расстояние от точки прыжка до обьекта (100 Диаметров)
//расстояние измеряется в МегаМетрах (1 Мм = 1000 Км)
func JumpPointFromObject(diameter float64) float64 {
	return utils.RoundFloat64(diameter*100.0, 3)
}

func JumpPointDistance(world *world.World) (megaMeters float64) {
	/*
	   start = JumpShadowRadius
	   end = Planetradius

	*/
	star := world.HomeStar()
	starJSOrbit := JumpDriveLimitStar(star)
	planetOrbit := world.Orbit()
	planetJSDistance := float64(TrvCore.EhexToDigit(world.PlanetaryData("Size"))*100) * 1.6
	if planetJSDistance == 0 {
		planetJSDistance = 400 * 1.6
	}
	if starJSOrbit > planetOrbit {
		au := OrbitToAU(starJSOrbit) - OrbitToAU(planetOrbit)
		planetJSDistance = AUnitsToMegameters(au)
		fmt.Println("DEBUG:", au, planetJSDistance, starJSOrbit)
	}
	//fmt.Println(planetJSDistance, starJSOrbit, planetOrbit)
	return planetJSDistance
}

func JumpPointDistanceExtended(planetOrbitAU float64, planetSizeKm int, starDiameter float64) float64 {
	starShadow := utils.RoundFloat64((starDiameter*13927.2)/AU2Megameters, 2)
	planetShadow := utils.RoundFloat64(float64(planetSizeKm*100*1000)/AU2Megameters, 2)
	jumpPoint := planetShadow
	if starShadow > planetShadow+planetOrbitAU {
		jumpPoint = starShadow
	}
	jumpPoint = utils.RoundFloat64(jumpPoint, 2)
	return jumpPoint
}

func StarJumpShadowAU(starDiameter float64) float64 {
	//1392720 km = diameter of Sun (1 Sol)
	starShadow := utils.RoundFloat64((starDiameter*1392.7)/AU2Megameters*100, 2)

	return starShadow
}

/*
sunD := 13927.7
sS = 0.76
pO = 0.64
pS = 0.12
jP = pO + pS
if sS > jP {
jP = sS
}

*/

//TravelTime - distance считается в Мегаметрах (тясяча км), thrust - среднее ускореение которое может поддерживать пилот
//на протяжении перелета. Возвращает количество часов требуемое для перелета.
// TODO: рассмотреть необходимость и возможность ввести тест меру эффекта от теста 'Pilot(Routine, DEX)'
//или механики Risk/Reward в Merchant Prince.
func TravelTime(distance float64, thrust float64) (travelHours float64) {
	travelHours = math.Sqrt(distance / thrust / 32.4) // SQRT(Distance (MegaMeters) / Acceleration (G or Trust) / 32.4) -- result Time (Hours) -- 32400 - это вес при 1G
	travelHours = utils.RoundFloat64(travelHours, 1)
	return travelHours
}

//TravelTime - distance считается в Мегаметрах (тясяча км), thrust - среднее ускореение которое может поддерживать пилот
//на протяжении перелета. Возвращает количество часов требуемое для перелета.
// TODO: рассмотреть необходимость и возможность ввести тест меру эффекта от теста 'Pilot(Routine, DEX)'
//или механики Risk/Reward в Merchant Prince.
func TravelTimeStr(distance float64, thrust float64) string {
	tH := TravelTime(distance, thrust)
	return hoursToTimeStr(tH)
}

//TravelTimeAU - distance считается в AU (Астрономическая единица), thrust - среднее ускореение которое может поддерживать пилот
//на протяжении перелета. Возвращает количество часов требуемое для перелета.
// TODO: рассмотреть необходимость и возможность ввести тест меру эффекта от теста 'Pilot(Routine, DEX)'
//или механики Risk/Reward в Merchant Prince.
func TravelTimeAU(au float64, thrust float64) (travelHours float64) {
	distance := au * 149597.8707                      //149597.8707
	travelHours = math.Sqrt(distance / thrust / 32.4) // SQRT(Distance (MegaMeters) / Acceleration (G or Trust) / 32.4) -- result Time (Hours) -- 32400 - это вес при 1G
	travelHours = utils.RoundFloat64(travelHours, 1)
	return travelHours
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////

//TravelTimeOrbitalArrival - возвращает колличество часов требуемое для фиксации орбиты с произвольного расстояния
func TravelTimeOrbitalArrival(planetSizeCode string, diameters int, trust float64) float64 {
	planetTravelDistance := float64(TrvCore.EhexToDigit(planetSizeCode)*(diameters-10)) * 1.6
	if planetTravelDistance == 0 {
		planetTravelDistance = 400 * 1.6 * float64(diameters-10)
	}
	planetTravelDistance = utils.RoundFloat64(planetTravelDistance, 1)
	return TravelTime(planetTravelDistance, trust)
}

func hoursToTimeStr(hours float64) string {
	dd := int(hours) / 24
	hh := int(hours) % 24
	mm := int((hours - utils.RoundFloat64(hours, 0)) * 60)
	//fmt.Println(hours, utils.RoundFloat64(hours, 0))
	if mm < 0 {
		//hh++
		mm += 60
	}
	return strconv.Itoa(dd) + " days " + strconv.Itoa(hh) + " hours " + strconv.Itoa(mm) + " minutes"
}
