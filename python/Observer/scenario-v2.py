#!/usr/bin/python3
from abc import ABC, ABCMeta, abstractclassmethod, abstractmethod
from template import Observer, Observable

class WaterHeater(Observable):
    def __init__(self):
        super().__init__()
        self._temperature = 25
    
    def GetTemperature(self):
        return self._temperature

    def SetTemperature(self, temperature):
        self._temperature = temperature
        print("Updated temperature: " + str(self.GetTemperature()))
        self.notifyObservers()

#  The Observable already give these following function. 
#    
#    def AddObserver(self, observer):
#        self._observers.append(observer)

#    def notifies(self):
#        for subscriber in self._observers:
#            subscriber.update(self)

#  This class has been moved to the template.py
#
#class Observer(metaclass=ABCMeta):
#    @abstractmethod
#    def update(self, waterHeater):
#        pass

class WashingMode(Observer):
    def update(self, observable, object):
        if isinstance(obsevable, WaterHeater) and observable.GetTemperature() < 70 and observable.GetTemperature() > 50:
            print("Hot water is ready for you to take a shower.")

class DrinkingMode(Observer):
    def update(self, observable, object):
        if isinstance(observable, WaterHeater) and observable.GetTemperture() >= 100:
            print("Hot water is ready for you to drink.")

def main():
    # Create objects
    heater = WaterHeater()
    washingObser = WashingMode()
    drinkingObser = DrinkingMode()

    # Registerion of observers
    heater.AddObserver(washingObser)
    heater.AddObserver(drinkingObser)

    # cases
    heater.SetTemperature(40)
    heater.SetTemperature(60)
    heater.SetTemperature(100)


if __name__ == "__main__":
    main()
