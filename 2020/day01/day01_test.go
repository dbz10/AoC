package main

import "testing"

func TestFindPair(t *testing.T) {
	var inputs = []int{1721, 979, 366, 299, 675, 1456}

	var gotV1, gotV2, err = findPair(inputs, 2020)

	if err != nil {
		t.Error("Shoud not have failed")
	}
	if (gotV1 != 1721) && (gotV2 != 299) {
		t.Errorf("Expected 1721 and 299, but got %d, %d", gotV1, gotV2)
	}
}

func TestFindTriple(t *testing.T) {
	var inputs = []int{1721, 979, 366, 299, 675, 1456}
	var triplet, err = findThree(inputs, 2020)
	if err != nil {
		t.Error("Shoud not have failed")
	}
	if triplet != [3]int{979, 366, 675} {
		t.Errorf("Expected [979, 366, 675], but got %d", triplet)
	}
}
