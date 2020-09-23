package absfactory

import (
	"errors"
	"fmt"
)

const (
	CarFactoryType  = 1
	BikeFactoryType = 2
)

//VehicleFactory - Строит Фабрики
type VehicleFactory interface {
	Build(v int) (Vehicle, error)
}

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d is not recognized"))
	case CarFactoryType:
		return new(CarFactory), nil
	case BikeFactoryType:
		return new(BikeFactory), nil
	}
}
