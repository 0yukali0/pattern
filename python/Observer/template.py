#!/usr/bin/python3
from abc import ABCMeta, abstractclassmethod, abstractmethod

class Observer(metaclass=ABCMeta):
    @abstractmethod
    def update(self, observable, object):
        pass

class Observable:
    def __init__(self):
        self._observers = []

    def AddObserver(self, observer):
        self._observers.append(observer)
    
    def RemoveObserver(self, observer):
        self._observers.remove(observer)

    def NotifyObservers(self, object=0):
        for observer in self._observers:
            observer.update()
