package main

import (
	"comcast/trafficLight"
	"context"
	"log"
	"time"
)

func main() {

	ryg, err := trafficLight.NewTrafficLight(10*time.Second, 3*time.Second, 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	go func(t *trafficLight.TrafficLight, c context.Context) {
		err := t.Run(c)
		if err != nil {
			log.Fatal(err)
		}
		return
	}(ryg, ctx)

}
