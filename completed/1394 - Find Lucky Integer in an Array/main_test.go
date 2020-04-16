package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	input := []int{2, 2, 3, 4}

	expected := 2
	given := findLucky(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestSecond(t *testing.T) {
	input := []int{1, 2, 2, 3, 3, 3}

	expected := 3
	given := findLucky(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestThird(t *testing.T) {
	input := []int{2, 2, 2, 3, 3}

	expected := -1
	given := findLucky(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestForth(t *testing.T) {
	input := []int{5}

	expected := -1
	given := findLucky(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestSeventh(t *testing.T) {
	input := []int{7, 7, 7, 7, 7, 7, 7}

	expected := 7
	given := findLucky(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestOneOnly(t *testing.T) {
	input := []int{1}

	expected := 1
	given := findLucky(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}
