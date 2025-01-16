// Race management implementation
// Author: Cameron Denton
// Created: 01/14/2025
// Last Modified: 01/14/2025

package race

import (
	"fmt"
	interfaces "race_mgmt_go/internal/race/interfaces"
	"sort"
)

// Race runs a simulation with the provided vehicles for a set number of iterations.
func Race(vehicles []interfaces.Vehicle, iterations int) {
	// Initialize all vehicles
	for _, vehicle := range vehicles {
		startSpeed := vehicle.Start()
		fmt.Printf("Start speed: %d\n", startSpeed)
		vehicle.Update_Position(startSpeed)
	}

	// Run the race for the specified iterations
	for i := 1; i <= iterations; i++ {
		fmt.Printf("\nIteration %d\n", i)

		for _, vehicle := range vehicles {
			fmt.Printf("%s (%s) is driving\n", vehicle.Get_Name(), vehicle.Type())
			driveSpeed := vehicle.Drive()
			fmt.Printf("Drive speed: %d\n", driveSpeed)
			vehicle.Update_Position(vehicle.Get_Position() + driveSpeed)
		}

		// Determine and print the standings
		sort.SliceStable(vehicles, func(i, j int) bool {
			return vehicles[i].Get_Position() > vehicles[j].Get_Position()
		})

		fmt.Println("\nStandings:")
		for idx, vehicle := range vehicles[:3] { // Display top 3 vehicles
			fmt.Printf("%d - %s (%s): Position %d\n", idx+1, vehicle.Get_Name(), vehicle.Type(), vehicle.Get_Position())
		}

		fmt.Println("===========================")
	}

	// Final results
	fmt.Println("\nFinal Results:")
	sort.SliceStable(vehicles, func(i, j int) bool {
		return vehicles[i].Get_Position() > vehicles[j].Get_Position()
	})
	for idx, vehicle := range vehicles {
		fmt.Printf("%d - %s (%s): Final Position %d\n", idx+1, vehicle.Get_Name(), vehicle.Type(), vehicle.Get_Position())
	}
}
