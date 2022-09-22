package main

type IObservable interface {
	add(observer *IObserver)
	remove(observer *IObserver)
	notify()
}

type TrafficLight struct {
	// Consider the color is for driver
	color     string
	timer     int
	observers []IObserver
}

func (tl *TrafficLight) add(o IObserver) {
	tl.observers = append(tl.observers, o)
}

func (tl *TrafficLight) remove(o IObserver) {
	var i = indexOf(o, tl.observers)
	var slice []IObserver
	slice = append(tl.observers[:i], tl.observers[i+1:]...)
	tl.observers = slice
}

func (tl *TrafficLight) notify() {
	for _, observer := range tl.observers {
		observer.update()
	}
}

func indexOf(element IObserver, slice []IObserver) int {
	for index, value := range slice {
		if element == value {
			return index
		}
	}
	return -1
}
