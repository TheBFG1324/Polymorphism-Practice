// Main package to initialize and run the race simulation
// Author: Cameron Denton
// Created: 01/16/2025
// Last Modified: 01/16/2025

package main

import (
	race "race_mgmt_go/internal/race"
	bikes "race_mgmt_go/internal/race/bikes"
	cars "race_mgmt_go/internal/race/cars"
	interfaces "race_mgmt_go/internal/race/interfaces"
)

func main() {
	// Initialize a list of vehicles participating in the race
	vehicles := []interfaces.Vehicle{
		bikes.NewKawasaki("Ninja", 10, 100, 2),
		bikes.NewDucati("Panigale", 15, 120, 3),
		cars.NewFerrari("F8 Tributo", 20, 150, 4),
		cars.NewLamborghini("Aventador", 25, 200, 5),
		cars.NewPorsche("911 GT3", 30, 180, 6),
	}

	// Run the race with the list of vehicles for 10 iterations
	race.Race(vehicles, 10)
}
