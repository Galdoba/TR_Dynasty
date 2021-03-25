package pbc

/*
Traveller creation is carried out in the following stages:
• Create and assign characteristics
• Choose a background package
• Choose a career package
• Finalise the Traveller




*/
const (
	PackageTypeBACKGROUND = 0
	PackageTypeCAREER     = 1
)

type Package struct {
	name      string
	pType     int //0 = background
	charBonus [6]int
	
}
