# This file contains the main function to run the race
# Author: Cameron Denton
# Created: 1/14/2025
# Last Modified: 1/14/2025

# Import the vehicles
from race.vehicles.bikes.ducati import Ducati
from race.vehicles.bikes.kawasaki import Kawasaki
from race.vehicles.cars.ferrari import Ferrari
from race.vehicles.cars.lamborghini import Lamborghini
from race.vehicles.cars.porsche import Porsche
from race.race import Race

# Main function that runs the race for 10 iterations
def main():
    vehicles = [
        Ducati(vehicle_type="Ducati",name="Demon", start_speed=5, max_speed=200, acceleration_rate=10, position=0),
        Kawasaki(vehicle_type="Kawasaki",name="Silver Fire" ,start_speed=7, max_speed=180, acceleration_rate=8, position=0),
        Ferrari(vehicle_type="Ferrari",name="Angel", start_speed=3, max_speed=300, acceleration_rate=15, position=0),
        Lamborghini(vehicle_type="Lamborghini",name="Lightning", start_speed=4, max_speed=280, acceleration_rate=12, position=0),
        Porsche(vehicle_type="Porsche", name="Sonic",start_speed=4, max_speed=250, acceleration_rate=11, position=0)
    ]
    raceObject = Race(vehicles=vehicles, iterations=10)
    raceObject.race()

# Call to main function
if __name__ == "__main__":
    main()

