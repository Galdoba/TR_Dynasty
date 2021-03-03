package dwd

import "github.com/Galdoba/TR_Dynasty/pkg/profile/uwp"

type DetailedWorldData struct {
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
