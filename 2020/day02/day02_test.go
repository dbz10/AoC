package main

import "testing"

func TestPartOne(t *testing.T) {
	if !partOne("abaa", "a", 1, 3) {
		t.Fail()
	}

	if partOne("abcdef", "a", 2, 3) {
		t.Fail()
	}

	if !partOne("dddddd", "d", 1, 7) {
		t.Fail()
	}

	if partOne("dddd", "d", 1, 3) {
		t.Fail()
	}

}

func TestPartTwo(t *testing.T) {
	if !partTwo("abaa", "a", 1, 2) {
		t.Fail()
	}

	if partTwo("abcdef", "a", 2, 3) {
		t.Fail()
	}

	if partTwo("dddddd", "d", 1, 3) {
		t.Fail()
	}

	if !partTwo("aaad", "d", 1, 4) {
		t.Fail()
	}

}
