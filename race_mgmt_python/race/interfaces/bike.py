# This file contains the abstract class Bike which is the parent class for all bikes
# Author: Cameron Denton
# Created: 1/14/2025
# Last Modified: 1/14/2025

# Import the ABC and abstractmethod from the abc module
from abc import abstractmethod
from interfaces.vehicle import Vehicle  # Assuming the Vehicle class is in this path

# Abstract class Bike inherits from Vehicle
class Bike(Vehicle):
    
    # Abstract method to perform a wheelie boost
    @abstractmethod
    def Wheelie_Boost(self):
        pass
    
    # Abstract method to perform an aerodynamic boost
    @abstractmethod
    def Aerodynamic_Boost(self):
        pass
