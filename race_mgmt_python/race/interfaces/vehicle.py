# This file contains the abstract class Vehicle which is the parent class for all cars and bikes
# Author: Cameron Denton
# Created: 1/14/2025
# Last Modified: 1/14/2025

# Import the ABC and abstractmethod from the abc module
from abc import ABC, abstractmethod

# Create the abstract class Vehicle
class Vehicle(ABC):

    # Constructor for the Vehicle class
    def __init__(self, vehicle_type, name, start_speed, max_speed, acceleration_rate, position):
        self._type = vehicle_type
        self._name = name
        self._start_speed = start_speed
        self._max_speed = max_speed
        self._acceleration_rate = acceleration_rate
        self._position = position

    # Abstract method to start the vehicle
    @abstractmethod
    def start(self):
        pass

    # Abstract method to stop the vehicle
    @abstractmethod
    def drive(self):
        pass

    # Returns the Vehicle type
    def Type(self):
        return self._type
    
    # Returns the Vehicle name
    def Name(self):
        return self._name
    
    # Returns the Vehicle start speed
    def Start_Speed(self):
        return self._start_speed
    
    # Returns the Vehicle start speed
    def Max_Speed(self):
        return self._max_speed
    
    # Returns the Vehicle max speed
    def Acceleration_Rate(self):
        return self._acceleration_rate

    # Returns the Vehicle position
    def Position(self):
        return self._position
    
    # Updates the Vehicle position
    def Update_Position(self, position):
        self._position = position
    