// Ducati bike implementation
// Author: Cameron Denton
// Created: 01/16/2025
// Last Modified: 01/16/2025

package bikes

import (
	"fmt"
	interfaces "race_mgmt_go/internal/race/interfaces"
	random "race_mgmt_go/internal/utils"
)

type Ducati struct {
	interfaces.BaseVehicle
}

// NewDucati creates and returns a new Ducati bike instance with the provided attributes.
func NewDucati(name string, startSpeed int, maxSpeed int, accelRate int) *Ducati {
	return &Ducati{
		BaseVehicle: interfaces.BaseVehicle{
			VehicleType: "Bike",
			Name:        name,
			StartSpeed:  startSpeed,
			MaxSpeed:    maxSpeed,
			AccelRate:   accelRate,
			Position:    0,
		},
	}
}

// Start calculates the initial speed of the Ducati bike,
// with a chance to apply a wheelie boost for increased speed.
func (d *Ducati) Start() int {
	startSpeed := d.Start_Speed()
	fmt.Println(d.Get_Name(), d.Type(), "is starting!")

	executeWheelie := random.RandomBool()

	if executeWheelie {
		wheelieBoost := d.Wheelie_Boost()
		fmt.Println("Executing wheelie! Boosting speed by", wheelieBoost)
		startSpeed += wheelieBoost
	}

	return startSpeed
}

// Drive calculates the Ducati bike's current speed during the race,
// considering its position, acceleration rate, and aerodynamic boost.
func (d *Ducati) Drive() int {
	currentPosition := d.Get_Position()
	maxSpeed := d.Max_Speed()
	accelerationRate := d.Acceleration_Rate()

	curSpeed := 0

	executeAerodynamicBoost := random.RandomBool()
	if executeAerodynamicBoost {
		curSpeed += d.Aerodynamic_Boost()
	}

	curSpeed += currentPosition * accelerationRate
	return random.Min(curSpeed, maxSpeed)
}

// Wheelie_Boost returns a random speed boost value for performing a wheelie.
func (d *Ducati) Wheelie_Boost() int {
	return random.RandomInt(1, 10)
}

// Aerodynamic_Boost returns a random speed boost value for applying aerodynamic effects.
func (d *Ducati) Aerodynamic_Boost() int {
	return random.RandomInt(1, 5)
}
