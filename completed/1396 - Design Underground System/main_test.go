package main

import (
	"testing"
)

func TestPerp(t *testing.T) {
	input := [][]int{{1, 2, 1, 2, 1, 2, 1}}

	expected := false
	given := true
	// given := hasValidPath(input)

	if given != expected {
		t.Errorf("%v given but %v expected for %v", given, expected, input)
	}
}
