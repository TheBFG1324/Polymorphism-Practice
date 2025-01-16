package main

import (
	race "race_mgmt_go/internal/race"
	bikes "race_mgmt_go/internal/race/bikes"
	cars "race_mgmt_go/internal/race/cars"
	interfaces "race_mgmt_go/internal/race/interfaces"
)

func main() {
	vehicles := []interfaces.Vehicle{
		bikes.NewKawasaki("Ninja", 10, 100, 2),
		bikes.NewDucati("Panigale", 15, 120, 3),
		cars.NewFerrari("F8 Tributo", 20, 150, 4),
		cars.NewLamborghini("Aventador", 25, 200, 5),
		cars.NewPorsche("911 GT3", 30, 180, 6),
	}
	race.Race(vehicles, 10)
}
