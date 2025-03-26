// Package trafficLight implements a simulation of a traffic light
package trafficLight

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const redLight = `
         .----'-'----.
         |   .===.   |
         |  /RRRRR\  |
         |  \RRRRR/  |
         |   .===.   |
         |   .===.   |
         |  /     \  |
         |  \     /  |
         |   .===.   |
         |   .===.   |
         |  /     \  |
         |  \     /  |
         |   .===.   |
         '-._______.-'`

const yellowLight = `
         .----'-'----.
         |   .===.   |
         |  /     \  |
         |  \     /  |
         |   .===.   |
         |   .===.   |
         |  /YYYYY\  |
         |  \YYYYY/  |
         |   .===.   |
         |   .===.   |
         |  /     \  |
         |  \     /  |
         |   .===.   |
         '-._______.-'`

const greenLight = `
         .----'-'----.
         |   .===.   |
         |  /     \  |
         |  \     /  |
         |   .===.   |
         |   .===.   |
         |  /     \  |
         |  \     /  |
         |   .===.   |
         |   .===.   |
         |  /GGGGG\  |
         |  \GGGGG/  |
         |   .===.   |
         '-._______.-'`

// A TrafficLight is an abstraction of a real world traffic light
// with the member values being the default lit times
type TrafficLight struct {
	redLit    time.Duration
	yellowLit time.Duration
	greenLit  time.Duration
}

// clearScreen blanks to CLI screen to prepare for drawing the next traffic light image
// This currently works only for UNIX based shells
func (t *TrafficLight) clearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// Run runs the traffic light simulation. It will run until interupted via the
// given Context, or it hits an unrecoverable error
func (t *TrafficLight) Run(c context.Context) (runErr error) {

done:
	for {
		select {
		case <-c.Done():
			runErr = fmt.Errorf("Exiting Traffic Light Simulator")
			break done
		default:
			redLit, yellowLit, greenLit := t.retrieveLitFromContextOrDefault(c)
			if err := t.clearScreen(); err != nil {
				runErr = fmt.Errorf("recieved error : %s, exiting", err)
				break done
			}
			fmt.Println(redLight)
			fmt.Println("\nPress CTRL + x for settings")
			time.Sleep(redLit)
			if err := t.clearScreen(); err != nil {
				runErr = fmt.Errorf("recieved error : %s, exiting", err)
				break done
			}
			fmt.Println(yellowLight)
			fmt.Println("\nPress CTRL + x for settings")
			time.Sleep(yellowLit)
			if err := t.clearScreen(); err != nil {
				runErr = fmt.Errorf("recieved error : %s, exiting", err)
				break done
			}
			fmt.Println(greenLight)
			time.Sleep(greenLit)
		}
	}
	return
}

// retrieveLitFromContextOrDefault takes the context and attempts to pull lit times for red, green and yellow.
// If it finds the values in the context it will validate the values are in spec. If the values aren't found in
// the context or context value doesn't evaluate then the defaults passed to the TrafficLight constructor are used
func (t *TrafficLight) retrieveLitFromContextOrDefault(c context.Context) (redLit, yellowLit, greenLit time.Duration) {
	var (
		retrievedFromContext bool
		ok                   bool
	)

	if redLit, ok = c.Value("red-lit").(time.Duration); !ok {
		redLit = t.redLit
	} else {
		//log.Info("retrieved red-lit from context")
		retrievedFromContext = true
	}

	if yellowLit, ok = c.Value("yellow-lit").(time.Duration); !ok {
		yellowLit = t.yellowLit
	} else {
		//log.Info("retrieved yellow-lit from context")
		retrievedFromContext = true
	}

	if greenLit, ok = c.Value("green-lit").(time.Duration); !ok {
		greenLit = t.greenLit
	} else {
		//log.Info("retrieved green-lit from context")
		retrievedFromContext = true
	}

	// If at least one value is retrieved from Context then re-validate
	if retrievedFromContext {
		err := validateLitTimes(redLit, yellowLit, greenLit)
		if err != nil {
			//log.Warning(fmt.Sprintf("error validating lit time from context : %s, using defaults", err.Error()))
			redLit = t.redLit
			yellowLit = t.yellowLit
			greenLit = t.greenLit
		}
	}

	return
}

// NewTrafficLight takes 3 durations as parameters. Each parameter specifics how long
// a particular color remains lit. The function validates the lit times against
// hard coded defaults. If the durations are in spec it will return a *TrafficLight
// or an error if the validation fails
func NewTrafficLight(redLit, yellowLit, greenLit time.Duration) (*TrafficLight, error) {

	if err := validateLitTimes(redLit, yellowLit, greenLit); err != nil {
		return nil, err
	}

	return &TrafficLight{
		redLit:    redLit,
		yellowLit: yellowLit,
		greenLit:  greenLit,
	}, nil
}

// validateLitTimes determines if the 3 given durations are within spec and
// returns a populated error if a value doesn't work
func validateLitTimes(redLit, yellowLit, greenLit time.Duration) error {
	if redLit < 5*time.Second {
		return fmt.Errorf("red Light duration must be 5 seconds or more")
	}

	if redLit > 10*time.Minute {
		return fmt.Errorf("red Light duration must be 10 minutes or less")
	}

	if yellowLit < 5*time.Second {
		return fmt.Errorf("yellow Light duration must be 5 seconds or more")
	}

	if yellowLit > 30*time.Second {
		return fmt.Errorf("yellow Light duration must be 30 seconds or less")
	}

	if greenLit < 5*time.Second {
		return fmt.Errorf("green Light duration must be 5 seconds or more")
	}

	if greenLit > 10*time.Minute {
		return fmt.Errorf("green Light duration must be 10 minutes or less")
	}

	return nil
}
