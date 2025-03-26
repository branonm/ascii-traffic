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

type TrafficLight struct {
	redLit    time.Duration
	yellowLit time.Duration
	greenLit  time.Duration
}

func (t *TrafficLight) clearScreen() error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// TODO: pull lit times from context
func (t *TrafficLight) Run(c context.Context) error {
done:
	for {
		select {
		case <-c.Done():
			break done
		default:
			if err := t.clearScreen(); err != nil {
				fmt.Println(err)
			}
			fmt.Println(redLight)
			fmt.Println("\nPress CTRL + x for settings")
			time.Sleep(t.redLit)
			if err := t.clearScreen(); err != nil {
				fmt.Println(err)
			}
			fmt.Println(yellowLight)
			fmt.Println("\nPress CTRL + x for settings")
			time.Sleep(t.yellowLit)
			if err := t.clearScreen(); err != nil {
				fmt.Println(err)
			}
			fmt.Println(greenLight)
			time.Sleep(t.greenLit)
		}
	}
	return nil
}

func NewTrafficLight(redLit, yellowLit, greenLit time.Duration) (*TrafficLight, error) {

	if redLit < 5*time.Second {
		return nil, fmt.Errorf("red Light duration must be 5 seconds or more")
	}

	if redLit > 10*time.Minute {
		return nil, fmt.Errorf("red Light duration must be 10 minutes or less")
	}

	if yellowLit < 5*time.Second {
		return nil, fmt.Errorf("yellow Light duration must be 5 seconds or more")
	}

	if yellowLit > 30*time.Second {
		return nil, fmt.Errorf("yellow Light duration must be 30 seconds or less")
	}

	if greenLit < 5*time.Second {
		return nil, fmt.Errorf("green Light duration must be 5 seconds or more")
	}

	if greenLit > 10*time.Minute {
		return nil, fmt.Errorf("green Light duration must be 10 minutes or less")
	}

	return &TrafficLight{
		redLit:    redLit,
		yellowLit: yellowLit,
		greenLit:  greenLit,
	}, nil
}
