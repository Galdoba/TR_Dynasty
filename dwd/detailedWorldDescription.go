package dwd

import (
	"fmt"

	"github.com/Galdoba/TR_Dynasty/pkg/dice"
	"github.com/Galdoba/TR_Dynasty/pkg/profile/uwp"
)

type DetailedWorldData struct {
	dicepool dice.Dicepool
	//Astronomical:
	Stellar      string
	PositionCode string
	UWP          uwp.UWP
	//Size Related Details
	BasicWorldType            string // (star/planet/asteroid/GasGigant)
	Diametr                   float64
	DiameterAsteroidPlanetoid string
	Density                   float64 //is expressed in standard densities (Terra=1, or 5.517 grams per cubic centimeter)
	DensityType               string
	UWPsize                   int
	Mass                      float64 // (Terra=1)
	Gravity                   float64 // (Terra=1)
	BeltZones                 string  // "N-10 M-20 C-70"
	BeltOrbitWidth            float64 // AU
	BeltProfileNotation       string  // "0.1/100, N-10 M-20 C-70, 0.22 au"
	PlanetOrbitalPeriod       float64 // local Year
	StellarMass               float64 // (Sol=1)
	OrbitalDistance           float64 // au
	OrbitalPeriod             float64 // local day
	SateliteOrbitalPeriod     float64 // (Lunar month=1)
	SateliteOrbitalDistance   float64 // Mm
	RotationPeriod            float64 // duration
	AxialTilt                 float64 // ГРАДУСЫ
	OrbitalEccentricity       float64
	SeismicStressFactor       int
	//Atmo Related Details
	AtmosphericComposition     string
	SurfaceAtmosphericPressure float64 // atm
	NativeLife                 bool
	//Hydro Related Details
	HydrographicPercentage  int
	HydrographycComposition string
	Volcanoes               int
}

func GenerateDetailedWorldData(parentstar string, positionCode string, uwpStr string, basicWorldType string, nomenaSeed string) DetailedWorldData {
	dwd := DetailedWorldData{}
	dwd.dicepool = *dice.New().SetSeed(nomenaSeed)
	dwd.PositionCode = positionCode
	dwd.Stellar = parentstar
	switch basicWorldType {
	case "Star":
		dwd.BasicWorldType = basicWorldType
	case "Large Gas Gigant", "Small Gas Gigant", "Ice Gigant":
		dwd.BasicWorldType = "Gas Gigant"
	case "Ring System":
		dwd.UWP = *uwp.New("YR00000-0")
	default:
		dwd.UWP = *uwp.New(uwpStr)
		dwd.BasicWorldType = "Planet"
		if dwd.UWP.Size().Value() == 0 {
			dwd.BasicWorldType = "Asteroid"
		}
	}
	switch dwd.BasicWorldType {
	case "Planet":

	}
	//	dwd.BasicWorldType = basicWorldType
	return dwd
}

func (detwd *DetailedWorldData) PrintData() {
	switch detwd.BasicWorldType {
	case "Star":
		fmt.Println("Basic Type: ", detwd.BasicWorldType)
		fmt.Println("Class: ", detwd.Stellar)
	case "Gas Gigant":
		fmt.Println("Basic Type: ", detwd.BasicWorldType)
		fmt.Println("Parent Star: ", detwd.Stellar)
		fmt.Println("Position: ", detwd.PositionCode)
	case "Planet":
		fmt.Println("Basic Type: ", detwd.BasicWorldType)
		fmt.Println("Parent Star: ", detwd.Stellar)
		fmt.Println("Position: ", detwd.PositionCode)
		fmt.Println("UWP: ", detwd.UWP.String())
	case "Asteroid":
		fmt.Println("Basic Type: ", detwd.BasicWorldType)
		fmt.Println("Parent Star: ", detwd.Stellar)
		fmt.Println("Position: ", detwd.PositionCode)
		fmt.Println("UWP: ", detwd.UWP.String())
	case "Ring System":
		fmt.Println("Basic Type: ", detwd.BasicWorldType)
		fmt.Println("Parent Star: ", detwd.Stellar)
		fmt.Println("Position: ", detwd.PositionCode)
		fmt.Println("UWP: ", detwd.UWP.String())
	}

}

func (detwd *DetailedWorldData) rollPlanetDiameter() int {
	return 5000
}
