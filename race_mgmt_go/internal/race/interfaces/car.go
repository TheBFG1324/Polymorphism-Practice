package interfaces

// Car interface extends the Vehicle interface to include car-specific behaviors.
type Car interface {
	Vehicle
	Drift_Boost() int
	Slipstream_Boost() int
}
