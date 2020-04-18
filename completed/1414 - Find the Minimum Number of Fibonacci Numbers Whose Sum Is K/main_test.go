package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	input := 7

	expected := 2
	given := findMinFibonacciNumbers(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %d", given, expected, input)
	}
}

func TestThird(t *testing.T) {
	input := 19

	expected := 3
	given := findMinFibonacciNumbers(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %d", given, expected, input)
	}
}
