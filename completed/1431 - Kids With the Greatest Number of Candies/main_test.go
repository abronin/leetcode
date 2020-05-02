package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	input := 7

	expected := 2
	given := 3

	if given != expected {
		t.Errorf("%d given but %d expected for %d", given, expected, input)
	}
}

func TestArrays(t *testing.T) {
	input := 7

	expected := []int{1}
	given := []int{2}

	if len(given) != len(expected) {
		t.Errorf("%v given but %v expected for %d", given, expected, input)
		t.Errorf("Lengths of given and expected are different!")
		return
	}

	for i, v := range given {
		if v != expected[i] {
			t.Errorf("%v given but %v expected for %d", given, expected, input)
			return
		}
	}
}
