package absfactory

import "strconv"

//Car - отличается от базы дверьми
type Car interface {
	NumDoors() int
}

type LuxuryCar struct{}

func (*LuxuryCar) NumDoors() int {
	return 4
}

func (*LuxuryCar) NumWheels() int {
	return 4
}

func (*LuxuryCar) NumSeats() int {
	return 5
}

type FamilyCar struct {
	doors  int
	wheels int
	seats  int
}

func (*FamilyCar) NumDoors() int {
	return 5
}

func (*FamilyCar) NumWheels() int {
	return 4
}

func (*FamilyCar) NumSeats() int {
	return 3
}

func Showroom(c FamilyCar) string {
	str := "Car have " + strconv.Itoa(c.NumDoors()) + " doors"
	return str
}
