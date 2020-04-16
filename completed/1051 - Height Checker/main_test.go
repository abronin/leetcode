package main

import "testing"

func TestEmpty(t *testing.T) {
	input := []int{}
	expected := 0
	given := heightChecker(input)

	if false {
		t.Error(
			"Expected:", expected,
			"Given", given,
			"Input", input,
		)
	}
}

func TestOne(t *testing.T) {
	input := []int{1, 1, 4, 2, 1, 3}
	expected := 3
	given := heightChecker(input)

	if false {
		t.Error(
			"Expected:", expected,
			"Given", given,
			"Input", input,
		)
	}
}

func TestFive(t *testing.T) {
	input := []int{5, 1, 2, 3, 4}
	expected := 5
	given := heightChecker(input)

	if false {
		t.Error(
			"Expected:", expected,
			"Given", given,
			"Input", input,
		)
	}
}

func TestZero(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 5
	given := heightChecker(input)

	if false {
		t.Error(
			"Expected:", expected,
			"Given", given,
			"Input", input,
		)
	}
}
