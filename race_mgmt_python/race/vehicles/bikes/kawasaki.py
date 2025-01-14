# Class for Kawasaki bikes
# Author: Cameron Denton
# Created: 1/14/2025
# Last Modified: 1/14/2025

# Import needed interface and random module
from interfaces.bike import Bike
from random import randint

# Class for Kawasaki bikes
class Kawasaki(Bike):
    
    # Returns a wheelie boost integer
    def Wheelie_Boost(self):
        return randint(5, 15)
    
    # Returns an aerodynamic boost integer
    def Aerodynamic_Boost(self):
        return randint(3, 10)
    
    # Returns the start speed integer
    def start(self):
        start_speed = self.Start_Speed()
        print(f"Start speed: {start_speed}")

        execute_wheelie = True if randint(1, 3) == 1 else False
        
        if execute_wheelie:
            wheelie_boost = self.Wheelie_Boost()
            print(f"Executing wheelie for a boost of {wheelie_boost} km/h")
            start_speed += wheelie_boost
        
        return start_speed
    
    # Returns the drive speed integer
    def drive(self):
        cur_position = self.Position()
        max_speed = self.Max_Speed()
        acceleration_rate = self.Acceleration_Rate()

        cur_speed = 0

        execute_aerodynamic_boost = True if randint(1, 4) == 1 else False
        if execute_aerodynamic_boost:
            aerodynamic_boost = self.Aerodynamic_Boost()
            print(f"Executing aerodynamic boost of {aerodynamic_boost} km/h")
            cur_speed += aerodynamic_boost

        cur_speed += cur_position * acceleration_rate

        return min(cur_speed, max_speed)
