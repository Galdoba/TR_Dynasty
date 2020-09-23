package absfactory

import (
	"errors"
	"fmt"
)

const (
	LuxuryCarType  = 1
	FamilyCarType  = 2
	SportsBikeType = 1
	CruiseBikeType = 2
)

//Vehicle - база, то есть все общее
type Vehicle interface {
	NumWheels() int
	NumSeats() int
}

type CarFactory struct{}

func (c *CarFactory) Build(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle type %d unrecognised\n", v))
	}
}

type BikeFactory struct{}

func (m *BikeFactory) Build(v int) (Vehicle, error) {
	switch v {
	case SportsBikeType:
		return new(SportsBike), nil
	case CruiseBikeType:
		return new(CruiseBike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle type %d unrecognised\n", v))
	}
}
