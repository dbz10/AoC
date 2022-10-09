package main

import (
	"testing"
)

func TestWrapIndex(t *testing.T) {
	if wrapIndex(1, 3) != 1 {
		t.Errorf("Expected %d, got %d", 1, wrapIndex(4, 3))
	}
	if wrapIndex(3, 3) != 0 {
		t.Errorf("Expected %d, got %d", 0, wrapIndex(4, 3))
	}

	if wrapIndex(4, 3) != 1 {
		t.Errorf("Expected %d, got %d", 1, wrapIndex(4, 3))
	}
}

func TestPartOne(t *testing.T) {
	testContent, err := readFile("day03_test_input.txt")
	if err != nil {
		t.Fatal(err)
	}

	if testContent[0][0] != "." {
		t.Errorf("Expected %s, got %s", ".", testContent[0][0])
	}

	want := 7
	got := partOne(testContent)

	if got != want {
		t.Errorf("Wanted %d trees encountered, got %d", want, got)
	}
}

func TestPartTwo(t *testing.T) {
	testContent, err := readFile("day03_test_input.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := 336

	var treesEncounteredPartTwo int = 1

	for _, slope := range []scooterSlope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2}} {
		treesEncounteredPartTwo *= partTwo(testContent, slope.xStep, slope.yStep)
	}
	got := treesEncounteredPartTwo

	if got != want {
		t.Errorf("Wanted %d trees encountered, got %d", want, got)
	}
}
