package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	input := []int{-3, 2, -3, 4, 2}

	expected := 5
	given := minStartValue(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestSecond(t *testing.T) {
	input := []int{1, 2}

	expected := 1
	given := minStartValue(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestThird(t *testing.T) {
	input := []int{1, -2, -3}

	expected := 5
	given := minStartValue(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}

func TestEmpty(t *testing.T) {
	input := []int{}

	expected := 1
	given := minStartValue(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}
