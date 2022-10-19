from abc import ABCMeta, abstractmethod

class Water:
    def __init__(self, state):
        self.setState(state)
        self.setTemperature(25)
    def setState(self, state):
        self._state = state
    def getState(self):
        return self._state
    def changeState(self, nextState):
        if (self.getState()):
            print("From ", self.getState().getName(), " to ", nextState.getName())
        else:
            print("State initialization ", state.getName())
        self.setState(nextState)
    def getTemperature(self):
        return self._temperature
    def setTemperature(self, temperature):
        self._temperature = temperature
        if (self.getTemperature() <= 0):
            self.changeState(SolidState("solid"))
        elif (self.getTemperature() <= 100):
            self.changeState(LiquidState("liquid"))
        else:
            self.changeState(GaseousState("Gas"))
    def riseTemperature(self, step):
        self.setTemperature(self.getTemperature() + step)
    def reduceTemperature(self, step):
        self.setTemperature(self.getTemperature() - step)
    def behavior(self):
        self.getState().behavior(self)

class State(metaclass=ABCMeta):
    def __init__(self, name):
        self.setName(name)
    def getName(self):
        return self._name
    def setName(self, name):
        self._name = name
    @abstractmethod
    def behavior(self, water):
        pass

class SolidState(State):
    def __init__(self, name):
        super().__init__(name)
    def behavior(self, water):
        print(self.getName(), str(water.getTemperature()))

class LiquidState(State):
    def __init__(self, name):
        super().__init__(name)
    def behavior(self, water):
        print(self.getName(), str(water.getTemperature()))

class GaseousState(State):
    def __init__(self, name):
        super().__init__(name)
    def behavior(self, water):
        print(self.getName(), str(water.getTemperature()))

def testState():
    water = Water(LiquidState("liquid"))
    steps = [0, -50, 30, 200]
    for step in steps:
        if step >= 0:
            water.riseTemperature(step)
        else:
            water.reduceTemperature(step)
        water.behavior()

testState()