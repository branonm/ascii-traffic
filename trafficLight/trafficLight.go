package trafficLight

import "context"

type TrafficLight interface {
	Run(c context.Context)
}
