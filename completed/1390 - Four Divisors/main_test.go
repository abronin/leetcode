package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	input := []int{21, 4, 7}
	expected := 32
	given := sumFourDivisors(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %v", given, expected, input)
	}
}
