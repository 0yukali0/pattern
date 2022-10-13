# Observer pattern
Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.
There is one 'observable' called subject and many 'observers' called 'listeners'. 

# class diagram
```
- - - - - - - - - - - - - -
	Observable
- - - - - - - - - - - - - -
- obeservers: ObserverList
- - - - - - - - - - - - - -
+ addObserver(Observer)
+ 
```
