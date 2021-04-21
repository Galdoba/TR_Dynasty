package asset

type asset struct {
	name              string
	description       string
	usage             string
	numericalValues   []int
	numericalValuesFl []float64
	list1             []string
}

type Asset interface {
	Valid() asset
}

func (a *asset) Valid() asset {
	return *a
}

func (a *asset) Name() string {
	return a.name
}
func (a *asset) Description() string {
	return a.description
}
func (a *asset) NumericalValues() []int {
	return a.numericalValues
}
func (a *asset) NumericalValuesFl() []float64 {
	return a.numericalValuesFl
}
func (a *asset) List1() []string {
	return a.list1
}
