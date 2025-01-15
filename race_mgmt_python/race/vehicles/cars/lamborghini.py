# Class for Lamborghini cars
# Author: Cameron Denton
# Created: 1/14/2025
# Last Modified: 1/14/2025

# Import needed interface and random module
from race.interfaces.car import Car
from random import randint

# Class for Lamborghini cars
class Lamborghini(Car):
    
    # Returns a drift boost integer
    def Drift_Boost(self):
        return randint(15, 25)
    
    # Returns a slipstreaming boost integer
    def Slipstreaming_Boost(self):
        return randint(10, 20)
    
    # Returns the start speed integer
    def start(self):
        start_speed = self.Start_Speed()
        print(f"Start speed: {start_speed}")

        execute_drift = True if randint(1, 5) <= 2 else False
        
        if execute_drift:
            drift_boost = self.Drift_Boost()
            print(f"Executing drift boost of {drift_boost} km/h")
            start_speed += drift_boost
        
        return start_speed
    
    # Returns the drive speed integer
    def drive(self):
        cur_position = self.Position()
        max_speed = self.Max_Speed()
        acceleration_rate = self.Acceleration_Rate()

        cur_speed = 0

        execute_slipstreaming = True if randint(1, 10) <= 3 else False
        if execute_slipstreaming:
            slipstreaming_boost = self.Slipstreaming_Boost()
            print(f"Executing slipstreaming boost of {slipstreaming_boost} km/h")
            cur_speed += slipstreaming_boost

        cur_speed += cur_position * acceleration_rate

        return min(cur_speed, max_speed)
