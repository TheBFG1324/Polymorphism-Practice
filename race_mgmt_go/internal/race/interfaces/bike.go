package race

// Bike interface extends the Vehicle interface to include bike-specific behaviors.
type Bike interface {
	Vehicle
	Wheelie_Boost() int
	Aerodynamic_Boost() int
}
