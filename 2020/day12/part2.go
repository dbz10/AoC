package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// Part 2 was kind of fun

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	ourShip := waypointFollowingShip{0.0, 0.0, waypoint{10.0, 1.0}}
	for scanner.Scan() {
		l := scanner.Text()
		instruction := string(l[0])
		num, err := strconv.Atoi(l[1:])
		n := float64(num)
		if err != nil {
			log.Fatal(err)
		}

		switch instruction {
		case "N":
			ourShip.w.yRel += n
		case "E":
			ourShip.w.xRel += n
		case "S":
			ourShip.w.yRel -= n
		case "W":
			ourShip.w.xRel -= n
		case "R":
			ourShip.w.Rotate(-n)
		case "L":
			ourShip.w.Rotate(n)
		case "F":
			ourShip.Forward(n)
		default:
			log.Fatalf("Got an unexpected instruction %s", l)
		}
	}

	fmt.Printf("Part One: Final position is %f %f, for a Manhattan Distance of %f.\n",
		ourShip.x, ourShip.y, math.Abs(float64(ourShip.x))+math.Abs(float64(ourShip.y)))

}

type waypointFollowingShip struct {
	x float64
	y float64
	w waypoint
}

type waypoint struct {
	xRel float64
	yRel float64
}

func (s *waypointFollowingShip) Forward(n float64) {

	s.x += n * s.w.xRel
	s.y += n * s.w.yRel
}

func (w *waypoint) Rotate(d float64) {
	w.xRel, w.yRel = w.xRel*math.Cos(d*math.Pi/180.0)-w.yRel*math.Sin(d*math.Pi/180), w.xRel*math.Sin(d*math.Pi/180)+w.yRel*math.Cos(d*math.Pi/180.0)
}
