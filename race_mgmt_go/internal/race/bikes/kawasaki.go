// Kawasaki bike implementation
// Author: Cameron Denton
// Created: 01/16/2025
// Last Modified: 01/16/2025

package bikes

import (
	"fmt"
	interfaces "race_mgmt_go/internal/race/interfaces"
	random "race_mgmt_go/internal/utils"
)

type Kawasaki struct {
	interfaces.BaseVehicle
}

// NewKawasaki creates and returns a new Kawasaki bike instance.
func NewKawasaki(name string, startSpeed int, maxSpeed int, accelRate int) *Kawasaki {
	return &Kawasaki{
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

// Start calculates the initial speed of the Kawasaki bike,
// optionally applying a wheelie boost with a 33% chance.
func (d *Kawasaki) Start() int {
	startSpeed := d.Start_Speed()
	fmt.Println(d.Get_Name(), d.Type(), "is starting!")

	executeWheelie := random.RandomChance(33)

	if executeWheelie {
		wheelieBoost := d.Wheelie_Boost()
		fmt.Println("Executing wheelie! Boosting speed by", wheelieBoost)
		startSpeed += wheelieBoost
	}

	return startSpeed
}

// Drive calculates the current speed of the Kawasaki bike
// based on its position, acceleration rate, and a 25% chance
// to apply an aerodynamic boost.
func (d *Kawasaki) Drive() int {
	currentPosition := d.Get_Position()
	maxSpeed := d.Max_Speed()
	accelerationRate := d.Acceleration_Rate()

	curSpeed := 0

	executeAerodynamicBoost := random.RandomChance(25)
	if executeAerodynamicBoost {
		curSpeed += d.Aerodynamic_Boost()
	}

	curSpeed += currentPosition * accelerationRate
	return random.Min(curSpeed, maxSpeed)
}

// Wheelie_Boost returns a random boost value when performing a wheelie.
func (k *Kawasaki) Wheelie_Boost() int {
	return random.RandomInt(5, 15)
}

// Aerodynamic_Boost returns a random boost value when applying aerodynamic enhancements.
func (k *Kawasaki) Aerodynamic_Boost() int {
	return random.RandomInt(3, 10)
}
