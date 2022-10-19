from framework import Context, State

class Water(Context):
    def __init__(self):
        super().__init__()
        registerDefaultStates()
        self.setTemperature(25)
    def __init__(self, state):
        self.registerDefaultStates()
    def registerDefaultStates(self):
        self.addState(SolidState("solid"))
        self.addState(LiquidState("liquid"))
        self.addState(GaseousState("gas"))
    def getTemperature(self):
        return self._getStateInfo()
    def setTemperature(self, temperature):
        self._setStateInfo(temperature)
    def riseTemperature(self, step):
        self.setTemperature(self.getTemperature() + step)
    def reduceTemperature(self, step):
        self.setTemperature(self.getTemperature() - step)
    def behavior(self):
        state = self.getState()
        if (isinstance(state, State)):
            state.behavior(self)


def singleton(cls, *args, **kwargs):
    instance = {}
    def _singleton(*args, **kwargs):
        if cls not in instance:
            instance[cls] = cls(*args, **kwargs)
        return instance[cls]
    return _singleton

@singleton
class SolidState(State):
    def __init__(self, name):
        super().__init__(name)
    def isMatch(self, stateInfo):
        return stateInfo < 0
    def behavior(self, context):
        print(self.getName(), str(context.getTemperature()))

@singleton
class LiquidState(State):
    def __init__(self, name):
        super().__init__(name)
    def isMatch(self, stateInfo):
        return (stateInfo < 100 and stateInfo > 0)
    def behavior(self, context):
        print(self.getName(), str(context.getTemperature()))

@singleton
class GaseousState(State):
    def __init__(self, name):
        super().__init__(name)
    def isMatch(self, stateInfo):
        return stateInfo >= 100
    def behavior(self, context):
        print(self.getName(), str(context.getTemperature()))

def testState():
    water = Water()
    steps = [0, -50, 30, 200]
    for step in steps:
        if step >= 0:
            water.riseTemperature(step)
        else:
            water.reduceTemperature(step)
        water.behavior()

testState()