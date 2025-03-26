package main

import (
	"bufio"
	"comcast/trafficLight"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"
)

func main() {

	redLit, yellowLit, greenLit := getLitFromInput()

	ryg, err := trafficLight.NewTrafficLight(
		time.Duration(redLit)*time.Second,
		time.Duration(yellowLit)*time.Second,
		time.Duration(greenLit)*time.Second,
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	go func(t trafficLight.TrafficLight, c context.Context) {
		t.Run(c)
		return
	}(ryg, ctx) // passing *TrafficLightASCII as a TrafficLight

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, os.Interrupt)

	switch <-sigChan {
	case os.Interrupt:
		ctx.Done()
		fmt.Println("Exiting Traffic Light Simulator")
	default:
		fmt.Println("Received unknown signal, exiting")
	}

}

// getLitFromInput prompts the user for the length each color light should be lit for
func getLitFromInput() (int, int, int) {
	var (
		redLit    int
		yellowLit int
		greenLit  int
	)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("How long should the red light be lit in seconds: ")
		redLitString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("failed to read")
			continue
		}

		redLitString = strings.TrimSpace(redLitString)

		redLit, err = strconv.Atoi(redLitString)

		if err != nil {
			fmt.Println(fmt.Sprintf("failed to convert %s to integer", redLitString))
			continue
		} else {
			break
		}
	}

	reader.Reset(os.Stdin)

	for {
		fmt.Print("How long should the yellow light be lit in seconds: ")
		yellowLitString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("failed to read")
			continue
		}

		yellowLitString = strings.TrimSpace(yellowLitString)

		yellowLit, err = strconv.Atoi(yellowLitString)

		if err != nil {
			fmt.Println(fmt.Sprintf("failed to convert %s to integer", yellowLitString))
			continue
		} else {
			break
		}
	}

	reader.Reset(os.Stdin)

	for {
		fmt.Print("How long should the green light be lit in seconds: ")
		greenLitString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("failed to read")
			continue
		}

		greenLitString = strings.TrimSpace(greenLitString)

		greenLit, err = strconv.Atoi(greenLitString)

		if err != nil {
			fmt.Println(fmt.Sprintf("failed to convert %s to integer", greenLitString))
			continue
		} else {
			break
		}
	}

	return redLit, yellowLit, greenLit
}
