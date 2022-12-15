package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type coord struct {
	x, y int
}

func manhattan(c1, c2 coord) int {
	return abs(c1.x-c2.x) + abs(c1.y-c2.y)
}

type sensor struct {
	loc           coord
	closestBeacon coord
	md            int
}

func main() {
	contents, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	sensors := []sensor{}
	for _, line := range strings.Split(string(contents), "\n") {
		sensors = append(sensors, parseSensor(line))
	}

	var beaconlessOnLine int

	sort.Slice(sensors,
		func(i, j int) bool {
			return sensors[i].loc.x-sensors[i].md < sensors[j].loc.x-sensors[j].md
		})
	xMin := sensors[0].loc.x - sensors[0].md
	sort.Slice(sensors,
		func(i, j int) bool {
			return sensors[i].loc.x+sensors[i].md > sensors[j].loc.x+sensors[j].md
		})
	xMax := sensors[0].loc.x + sensors[0].md
	for x := xMin; x <= xMax; x++ {
		if mustBeBeaconless(coord{x, 2000000}, sensors) {
			beaconlessOnLine++
		}
	}
	fmt.Printf("Part One: %d points cannot contain beacons on the line scanned\n", beaconlessOnLine)

	// Part Two
	// The vacant spot must be at the border of an exclusion zone.
	// So perhaps we can "just" walk along the borders

	allSensorLocations := []coord{}
	allBeaconLocations := []coord{}
	for _, s := range sensors {
		allSensorLocations = append(allSensorLocations, s.loc)
		allBeaconLocations = append(allBeaconLocations, s.closestBeacon)
	}

	sensors = []sensor{}
	for _, line := range strings.Split(string(contents), "\n") {
		sensors = append(sensors, parseSensor(line))
	}

	var tuningFrequency int
outer:
	for _, sensor := range sensors {
		sx := sensor.loc.x
		sy := sensor.loc.y
		boundary := sensor.md + 1
		for dx := -boundary; dx <= boundary; dx++ {
			for _, dy := range []int{-boundary + abs(dx), boundary - abs(dx)} {
				c := coord{sx + dx, sy + dy}
				if c.x < 0 || c.x > 4000000 {
					continue
				}
				if c.y < 0 || c.y > 4000000 {
					continue
				}
				if !mustBeBeaconless(c, sensors) && !contains(allSensorLocations, c) && !contains(allBeaconLocations, c) {
					tuningFrequency = 4000000*c.x + c.y
					break outer
				}
			}
		}

	}

	fmt.Printf("Part Two: The distress beacon has tuning frequency %d\n", tuningFrequency)

}

func parseSensor(line string) sensor {
	var sensorX, sensorY, beaconX, beaconY int
	fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
	return sensor{
		loc:           coord{sensorX, sensorY},
		closestBeacon: coord{beaconX, beaconY},
		md:            manhattan(coord{sensorX, sensorY}, coord{beaconX, beaconY}),
	}
}

func BeaconlessPoints(sensors []sensor) map[coord]struct{} {
	// The problem states that two beacons are never the same distance
	// from a sensor, however one beacon can be the closest beacon to
	// more than one sensors. So basically I think we want to
	// maintain a hashset of all of the spaces within the
	// envelop of the nearest beacon to each sensor.

	// OK! it seems this N^2 approach cannot handle the actual input
	// But may as well include my failure
	beaconlessPoints := map[coord]struct{}{}
	for _, sensor := range sensors {
		md := manhattan(sensor.loc, sensor.closestBeacon)
		for dx := -md; dx <= md; dx++ {
			for dy := -(md - abs(dx)); dy <= (md - abs(dx)); dy++ {
				cprime := coord{
					sensor.loc.x + dx,
					sensor.loc.y + dy,
				}
				if cprime != sensor.closestBeacon {
					beaconlessPoints[cprime] = struct{}{}
				}
			}
		}
	}
	return beaconlessPoints
}

func mustBeBeaconless(c coord, ss []sensor) bool {
	// If the point is within the exclusion zone of any sensor,
	// it must be beaconless.
	allBeacons := []coord{}
	for _, s := range ss {
		allBeacons = append(allBeacons, s.closestBeacon)
	}
	for _, sensor := range ss {
		if manhattan(c, sensor.loc) <= sensor.md && !contains(allBeacons, c) {
			return true
		}
	}
	return false
}

func contains(cc []coord, c coord) bool {
	for _, v := range cc {
		if c == v {
			return true
		}
	}
	return false
}

func printExclusionZone(sensors []sensor) {
	beaconlessPoints := BeaconlessPoints(sensors)
	keys := make([]coord, 0, len(beaconlessPoints))

	allBeacons := []coord{}
	for _, s := range sensors {
		allBeacons = append(allBeacons, s.closestBeacon)
	}
	allSensors := []coord{}
	for _, s := range sensors {
		allSensors = append(allSensors, s.loc)
	}

	for c := range beaconlessPoints {
		keys = append(keys, c)
	}
	xMin := keys[0].x
	xMax := keys[0].x
	yMin := keys[0].y
	yMax := keys[0].y
	for p := range beaconlessPoints {
		if p.x <= xMin {
			xMin = p.x
		}
		if p.x >= xMax {
			xMax = p.x
		}
		if p.y <= yMin {
			yMin = p.y
		}
		if p.y >= yMax {
			yMax = p.y
		}
	}

	xMin = 0
	xMax = 20
	yMin = 0
	yMax = 20

	fmt.Print(" ")
	for x := xMin; x <= xMax; x++ {
		if x%5 == 0 {
			fmt.Print(x)
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")

	for y := yMin; y <= yMax; y++ {
		if y < 0 {
			fmt.Print(y)
		} else if y < 10 {
			fmt.Printf("%d ", y)
		} else {
			fmt.Print(y)
		}
		for x := xMin; x <= xMax; x++ {
			if contains(allBeacons, coord{x, y}) {
				fmt.Print("+")
			} else if contains(allSensors, coord{x, y}) {
				fmt.Print("S")
			} else {
				_, exists := beaconlessPoints[coord{x, y}]
				if exists {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Print("\n")
	}
}

func abs(i int) int {
	return int(math.Abs(float64(i)))
}
