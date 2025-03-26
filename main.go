package main

import (
	"comcast/trafficLight"
)

func main() {

	ryg := trafficLight.TrafficLight{}

	err := ryg.Run()
	if err != nil {
		return
	}
}
