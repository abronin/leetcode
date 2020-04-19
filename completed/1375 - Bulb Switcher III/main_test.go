package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	input := []int{2, 1, 3, 5, 4}

	expected := 3
	given := numTimesAllBlue(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestSecond(t *testing.T) {
	input := []int{3, 2, 4, 1, 5}

	expected := 2
	given := numTimesAllBlue(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestThird(t *testing.T) {
	input := []int{4, 1, 2, 3}

	expected := 1
	given := numTimesAllBlue(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestForth(t *testing.T) {
	input := []int{2, 1, 4, 3, 6, 5}

	expected := 3
	given := numTimesAllBlue(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestFifth(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}

	expected := 6
	given := numTimesAllBlue(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}
