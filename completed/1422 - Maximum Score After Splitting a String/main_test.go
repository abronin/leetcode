package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	input := "011101"

	expected := 5
	given := maxScore(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %s", given, expected, input)
	}
}

func TestSecond(t *testing.T) {
	input := "00111"

	expected := 5
	given := maxScore(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %s", given, expected, input)
	}
}

func TestThird(t *testing.T) {
	input := "1111"

	expected := 3
	given := maxScore(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %s", given, expected, input)
	}
}
