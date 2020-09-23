package absfactory

//Bike - Отличается от базы кузовом
type Bike interface {
	GetBikeType() int
}

type SportsBike struct{}

func (*SportsBike) NumSeats() int {
	return 1
}

func (*SportsBike) NumWheels() int {
	return 2
}

func (*SportsBike) BikeType() int {
	return SportsBikeType
}

type CruiseBike struct{}

func (*CruiseBike) NumSeats() int {
	return 2
}

func (*CruiseBike) NumWheels() int {
	return 2
}

func (*CruiseBike) BikeType() int {
	return CruiseBikeType
}
