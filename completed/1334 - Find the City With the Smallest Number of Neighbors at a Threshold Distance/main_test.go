package main

import (
	"testing"
)

func TestMaxIDOfMinNormal(t *testing.T) {
	input := map[int]int{0: 10, 1: 5, 2: 8, 3: 5, 4: 10}
	n := 5

	expected := 3
	given := maxIDOfMinCount(n, input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestMaxIDOfMinCountZeros(t *testing.T) {
	input := map[int]int{0: 0, 1: 5, 2: 0, 3: 5, 4: 10}
	n := 5

	expected := 2
	given := maxIDOfMinCount(n, input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestFirst(t *testing.T) {
	input := [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}
	n := 4
	distanceThreshold := 4

	expected := 3
	given := findTheCity(n, input, distanceThreshold)

	if given != expected {
		t.Errorf("%d given but %d expected for %d", given, expected, input)
	}
}

func TestSecond(t *testing.T) {
	input := [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}
	n := 5
	distanceThreshold := 2

	expected := 0
	given := findTheCity(n, input, distanceThreshold)

	if given != expected {
		t.Errorf("%d given but %d expected for %d", given, expected, input)
	}
}
