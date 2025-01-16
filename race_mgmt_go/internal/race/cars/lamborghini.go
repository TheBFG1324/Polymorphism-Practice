// Lamborghini car implementation
// Author: Cameron Denton
// Created: 01/16/2025
// Last Modified: 01/16/2025

package cars

import (
	"fmt"
	interfaces "race_mgmt_go/internal/race/interfaces"
	random "race_mgmt_go/internal/utils"
)

// Lamborghini represents a Lamborghini car type, embedding BaseVehicle.
type Lamborghini struct {
	interfaces.BaseVehicle
}

// NewLamborghini creates and returns a new Lamborghini car instance with the provided attributes.
func NewLamborghini(name string, startSpeed int, maxSpeed int, accelRate int) *Lamborghini {
	return &Lamborghini{
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

// Start calculates the initial speed of the Lamborghini car,
// with a chance to apply a slipstream boost for increased speed.
func (d *Lamborghini) Start() int {
	startSpeed := d.Start_Speed()
	fmt.Println(d.Get_Name(), d.Type(), "is starting!")

	executeSlipstream := random.RandomChance(20)

	if executeSlipstream {
		slipstreamBoost := d.Slipstream_Boost()
		fmt.Println("Executing slipstream! Boosting speed by", slipstreamBoost)
		startSpeed += slipstreamBoost
	}

	return startSpeed
}

// Drive calculates the Lamborghini car's current speed during the race,
// considering its position, acceleration rate, and drift boost.
func (d *Lamborghini) Drive() int {
	currentPosition := d.Get_Position()
	maxSpeed := d.Max_Speed()
	accelerationRate := d.Acceleration_Rate()

	curSpeed := 0

	executeDriftBoost := random.RandomChance(30)
	if executeDriftBoost {
		curSpeed += d.Drift_Boost()
	}

	curSpeed += currentPosition * accelerationRate
	return random.Min(curSpeed, maxSpeed)
}

// Slipstream_Boost returns a random speed boost value for performing a slipstream.
func (d *Lamborghini) Slipstream_Boost() int {
	return random.RandomInt(10, 20)
}

// Drift_Boost returns a random speed boost value for applying drift effects.
func (d *Lamborghini) Drift_Boost() int {
	return random.RandomInt(15, 25)
}
