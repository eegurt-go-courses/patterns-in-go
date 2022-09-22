package main

import "fmt"

type IObserver interface {
	update()
}

type Driver struct {
	name         string
	trafficLight *TrafficLight
	povColor     string
}

type Pedestrian struct {
	name         string
	trafficLight *TrafficLight
	povColor     string
}

func (d *Driver) update() {
	d.povColor = d.trafficLight.color

	fmt.Printf("%v, the color is %v for you. %v seconds left.\n",
		d.name, d.povColor, d.trafficLight.timer)
}

func (p *Pedestrian) update() {
	switch p.trafficLight.color {
	case "Green":
		p.povColor = "Red"
	case "Red":
		p.povColor = "Green"
	}

	fmt.Printf("%v, the color is %v for you. %v seconds left.\n",
		p.name, p.povColor, p.trafficLight.timer)
}
