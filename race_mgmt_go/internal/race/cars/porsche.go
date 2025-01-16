// Porsche car implementation
// Author: Cameron Denton
// Created: 01/16/2025
// Last Modified: 01/16/2025

package cars

import (
	"fmt"
	interfaces "race_mgmt_go/internal/race/interfaces"
	random "race_mgmt_go/internal/utils"
)

// Porsche represents a Porsche car type, embedding BaseVehicle.
type Porsche struct {
	interfaces.BaseVehicle
}

// NewPorsche creates and returns a new Porsche car instance with the provided attributes.
func NewPorsche(name string, startSpeed int, maxSpeed int, accelRate int) *Porsche {
	return &Porsche{
		BaseVehicle: interfaces.BaseVehicle{
			VehicleType: "Car",
			Name:        name,
			StartSpeed:  startSpeed,
			MaxSpeed:    maxSpeed,
			AccelRate:   accelRate,
			Position:    0,
		},
	}
}

// Start calculates the initial speed of the Porsche car,
// with a chance to apply a slipstream boost for increased speed.
func (d *Porsche) Start() int {
	startSpeed := d.Start_Speed()
	fmt.Println(d.Get_Name(), d.Type(), "is starting!")

	executeSlipstream := random.RandomChance(60)

	if executeSlipstream {
		slipstreamBoost := d.Slipstream_Boost()
		fmt.Println("Executing slipstream! Boosting speed by", slipstreamBoost)
		startSpeed += slipstreamBoost
	}

	return startSpeed
}

// Drive calculates the Porsche car's current speed during the race,
// considering its position, acceleration rate, and drift boost.
func (d *Porsche) Drive() int {
	currentPosition := d.Get_Position()
	maxSpeed := d.Max_Speed()
	accelerationRate := d.Acceleration_Rate()

	curSpeed := 0

	executeDriftBoost := random.RandomBool()
	if executeDriftBoost {
		curSpeed += d.Drift_Boost()
	}

	curSpeed += currentPosition * accelerationRate
	return random.Min(curSpeed, maxSpeed)
}

// Slipstream_Boost returns a random speed boost value for performing a slipstream.
func (d *Porsche) Slipstream_Boost() int {
	return random.RandomInt(5, 12)
}

// Drift_Boost returns a random speed boost value for applying drift effects.
func (d *Porsche) Drift_Boost() int {
	return random.RandomInt(8, 18)
}
