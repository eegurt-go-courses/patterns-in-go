package main

import "fmt"

func main() {
	var trafficLight TrafficLight = TrafficLight{color: "Green", timer: 15}

	var driver1 Driver = Driver{name: "Mark", trafficLight: &trafficLight}
	var driver2 Driver = Driver{name: "Sam", trafficLight: &trafficLight}
	var pedestrian1 Pedestrian = Pedestrian{name: "David", trafficLight: &trafficLight}

	trafficLight.add(&driver1)
	trafficLight.add(&driver2)
	trafficLight.add(&pedestrian1)

	trafficLight.notify()

	// 5 seconds later
	fmt.Printf("%v and %v have passed through.\n", driver1.name, driver2.name)
	trafficLight.timer = 10

	trafficLight.remove(&driver1)
	trafficLight.remove(&driver2)

	trafficLight.notify()
}
