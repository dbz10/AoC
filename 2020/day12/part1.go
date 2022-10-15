package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	ourShip := ship{0, 0, 0.0}
	for scanner.Scan() {
		l := scanner.Text()
		instruction := string(l[0])
		num, err := strconv.Atoi(l[1:])
		if err != nil {
			log.Fatal(err)
		}

		switch instruction {
		case "N":
			ourShip.y += num
		case "E":
			ourShip.x += num
		case "S":
			ourShip.y -= num
		case "W":
			ourShip.x -= num
		case "R":
			ourShip.bearing -= float64(num)
		case "L":
			ourShip.bearing += float64(num)
		case "F":
			ourShip.Forward(num)
		default:
			log.Fatalf("Got an unexpected instruction %s", l)
		}
	}

	fmt.Printf("Part One: Final position is %d %d, for a Manhattan Distance of %d.\n",
		ourShip.x, ourShip.y, int(math.Abs(float64(ourShip.x))+math.Abs(float64(ourShip.y))))

}

type ship struct {
	x       int
	y       int
	bearing float64
}
