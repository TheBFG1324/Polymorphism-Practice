package race

// Vehicle interface defines the required methods that all vehicles must implement.
type Vehicle interface {
	Type() string
	Name() string
	Start_Speed() int
	Max_Speed() int
	Acceleration_Rate() int
	Position() int
	Update_Position(int) int
	Start() int
	Drive() int
}

// BaseVehicle struct provides a base implementation for common fields and methods
// shared by all vehicle types.
type BaseVehicle struct {
	vehicleType string
	name        string
	startSpeed  int
	maxSpeed    int
	accelRate   int
	position    int
}

// Type returns the type of the vehicle (e.g., "Car", "Bike").
func (v *BaseVehicle) Type() string {
	return v.vehicleType
}

// Name returns the name of the vehicle (e.g., "Ferrari Enzo").
func (v *BaseVehicle) Name() string {
	return v.name
}

// Start_Speed returns the initial starting speed of the vehicle.
func (v *BaseVehicle) Start_Speed() int {
	return v.startSpeed
}

// Max_Speed returns the maximum speed the vehicle can achieve.
func (v *BaseVehicle) Max_Speed() int {
	return v.maxSpeed
}

// Acceleration_Rate returns the rate at which the vehicle accelerates.
func (v *BaseVehicle) Acceleration_Rate() int {
	return v.accelRate
}

// Position returns the current position of the vehicle in the race.
func (v *BaseVehicle) Position() int {
	return v.position
}

// Update_Position updates the vehicle's position by adding the given distance
// and returns the updated position.
func (v *BaseVehicle) Update_Position(distance int) int {
	v.position += distance
	return v.position
}
