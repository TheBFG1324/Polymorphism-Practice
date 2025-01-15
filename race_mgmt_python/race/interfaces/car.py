# This file contains the abstract class Car which is the parent class for all Cars
# Author: Cameron Denton
# Created: 1/14/2025
# Last Modified: 1/14/2025

# Import the ABC and abstractmethod from the abc module
from abc import abstractmethod
from race.interfaces.vehicle import Vehicle

# Abstract class Car inherits from Vehicle
class Car(Vehicle):
    
    # Abstract method to perform a drift boost
    @abstractmethod
    def Drift_Boost(self):
        pass
    
    # Abstract method to perform an aerodynamic boost
    @abstractmethod
    def Slipstreaming_Boost(self):
        pass
