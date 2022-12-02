package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	opponentShapeDecoding := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}

	selfShapeDecoding := map[string]string{
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}

	outcomeDecoding := map[string]string{
		"X": "Loss",
		"Y": "Draw",
		"Z": "Win",
	}

	shapeScore := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	outcomeScore := map[string]int{
		"Loss": 0,
		"Draw": 3,
		"Win":  6,
	}

	var points1 int
	var points2 int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		opp := string(scanner.Text()[0])
		self := string(scanner.Text()[2])
		outcome1 := game(selfShapeDecoding[self], opponentShapeDecoding[opp])
		derivedShape := invertGame(opponentShapeDecoding[opp], outcomeDecoding[self])

		points1 += shapeScore[selfShapeDecoding[self]] + outcomeScore[outcome1]
		points2 += shapeScore[derivedShape] + outcomeScore[outcomeDecoding[self]]

	}

	fmt.Printf("Part One: Scored %d points according to the encrypted strategy guide\n", points1)
	fmt.Printf("Part Two: Scored %d points according to the encrypted strategy guide\n", points2)
}

func game(self, opp string) string {
	switch [2]string{self, opp} {
	case [2]string{"Rock", "Scissors"}:
		return "Win"
	case [2]string{"Paper", "Rock"}:
		return "Win"
	case [2]string{"Scissors", "Paper"}:
		return "Win"
	case [2]string{"Rock", "Rock"}:
		return "Draw"
	case [2]string{"Paper", "Paper"}:
		return "Draw"
	case [2]string{"Scissors", "Scissors"}:
		return "Draw"
	case [2]string{"Rock", "Paper"}:
		return "Loss"
	case [2]string{"Paper", "Scissors"}:
		return "Loss"
	case [2]string{"Scissors", "Rock"}:
		return "Loss"
	default:
		return "Huh?"
	}
}

func invertGame(opp, desiredOutcome string) string {
	switch [2]string{opp, desiredOutcome} {
	case [2]string{"Rock", "Win"}:
		return "Paper"
	case [2]string{"Paper", "Win"}:
		return "Scissors"
	case [2]string{"Scissors", "Win"}:
		return "Rock"
	case [2]string{"Rock", "Draw"}:
		return "Rock"
	case [2]string{"Paper", "Draw"}:
		return "Paper"
	case [2]string{"Scissors", "Draw"}:
		return "Scissors"
	case [2]string{"Rock", "Loss"}:
		return "Scissors"
	case [2]string{"Paper", "Loss"}:
		return "Rock"
	case [2]string{"Scissors", "Loss"}:
		return "Paper"
	default:
		return "Huh?"
	}
}
