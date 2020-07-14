package Astrogation

import (
	"fmt"
	"math"

	"github.com/Galdoba/TR_Dynasty/worldBuilder"
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

func JumpPointDistance(world *world.World) (megaMeters float64) {
	/*
	   start = JumpShadowRadius
	   end = Planetradius

	*/
	star := world.HomeStar()
	starJSOrbit := JumpDriveLimitStar(star)
	planetOrbit := world.Orbit()
	planetJSDistance := float64(world.Stats()["Size"]*100) * 1.6
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

//TravelTime - distance считается в Мегаметрах (тясяча км), thrust - среднее ускореение которое может поддерживать пилот
//на протяжении перелета. Возвращает количество часов требуемое для перелета.
// TODO: рассмотреть необходимость и возможность ввести тест меру эффекта от теста 'Pilot(Routine, DEX)'
//или механики Risk/Reward в Merchant Prince.
func TravelTime(distance float64, thrust float64) (travelHours float64) {
	travelHours = math.Sqrt(distance / thrust / 32.4) // SQRT(Distance (MegaMeters) / Acceleration (G or Trust) / 32.4) -- result Time (Hours) -- 32400 - это вес при 1G
	travelHours = utils.RoundFloat64(travelHours, 1)
	return travelHours
}
