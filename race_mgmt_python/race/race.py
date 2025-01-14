# Class to facillitate races
# Author: Cameron Denton
# Created: 1/14/2025
# Last Modified: 1/14/2025

# Class for running races
class Race:

    # Constructor that takes a list of vehicles and the number of iterations the race is run
    def __init__(self, vehicles, iterations):
        self._vehicles = vehicles
        self._iterations = iterations
    
    # Run the race
    def race(self):
        # Initialize all vehicles
        for vehicle in self._vehicles:
            print(f"{vehicle.Name()}:{vehicle.Type()} is starting")
            start_speed = vehicle.start()
            print(f"Start speed: {start_speed}")
            vehicle.Update_Position(start_speed)
        
        # Run the race for the specified iterations
        for i in range(self._iterations):
            print(f"\nIteration {i+1}")
            for vehicle in self._vehicles:
                print(f"{vehicle.Name()}:{vehicle.Type()} is driving")
                drive_speed = vehicle.drive()
                print(f"Drive speed: {drive_speed}")
                vehicle.Update_Position(drive_speed)
            
            # Determine and print the standings
            standings = sorted(self._vehicles, key=lambda v: v.Position(), reverse=True)
            print("\nStandings:")
            for idx, vehicle in enumerate(standings[:3], start=1):  # Get top 3 vehicles
                print(f"{idx} - {vehicle.Name()} ({vehicle.Type()}): Position {vehicle.Position()}")
            
            print("===========================")

        # Final results
        print("\nFinal Results:")
        final_standings = sorted(self._vehicles, key=lambda v: v.Position(), reverse=True)
        for idx, vehicle in enumerate(final_standings, start=1):
            print(f"{idx} - {vehicle.Name()} ({vehicle.Type()}): Final Position {vehicle.Position()}")
