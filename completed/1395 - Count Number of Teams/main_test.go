package main

import (
	"testing"
)

func TestPerp(t *testing.T) {
	input := []int{2, 5, 3, 4, 1}

	expected := 3
	given := numTeams(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestSec(t *testing.T) {
	input := []int{2, 1, 3}

	expected := 0
	given := numTeams(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestThird(t *testing.T) {
	input := []int{1, 2, 3, 4}

	expected := 4
	given := numTeams(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}
