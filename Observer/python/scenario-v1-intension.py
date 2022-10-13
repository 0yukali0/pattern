#!/usr/bin/python3
from  abc import ABC, ABCMeta, abstractclassmethod, abstractmethod

class WaterHeater:
    def __init__(self):
        self._observers = []
        self._temperature = 25
    
    def GetTemperature(self):
        return self._temperature

    def SetTemperature(self, temperature):
        self._temperature = temperature
        print("Updated temperature: " + str(self.GetTemperature()))
        self.notifies()

    def AddObserver(self, observer):
        self._observers.append(observer)

    def notifies(self):
        for subscriber in self._observers:
            subscriber.update(self)

class Observer(metaclass=ABCMeta):
    @abstractmethod
    def update(self, waterHeater):
        pass

class WashingMode(Observer):
    def update(self, waterHeater):
        if waterHeater.GetTemperature() >= 50 and waterHeater.GetTemperature() < 70:
            print("Warm water is ready and you can take a shower.")

class DrinkingMode(Observer):
    def update(self, waterHeater):
        if waterHeater.GetTemperature() >= 100:
            print("Hot water is ready and you can drink the water.")

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