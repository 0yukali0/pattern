from abc import ABCMeta, abstractmethod

class Context(metaclass=ABCMeta):
    def __init__(self):
        self._states = []
        self._curState = None
        self._stateInfo = 0
    def addState(self, state):
        if (state not in self._states):
            self._states.append(state)
    def changeState(self, state):
        if (state is None):
            return False
        if (self.getState() is None):
            print("State initialization", state.getName())
        else:
            print("From", self.getState().getName(), "to", state.getName())
        self.addState(state)
        self._setState(state)
        return True
    def getState(self):
        return self._curState
    def _setState(self, state):
        self._curState = state
    def _setStateInfo(self, stateInfo):
        self._stateInfo = stateInfo
        for state in self._states:
            if (state.isMatch(stateInfo)):
                self.changeState(state)
    def _getStateInfo(self):
        return self._stateInfo

class State(metaclass=ABCMeta):
    def __init__(self, name):
        self._setName(name)
    def getName(self):
        return self._name
    def _setName(self, name):
        self._name = name
    def isMatch(self, stateInfo):
        return False
    @abstractmethod
    def behavior(self, context):
        pass