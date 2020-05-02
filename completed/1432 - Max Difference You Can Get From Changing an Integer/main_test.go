package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	input := 123456

	expected := 820000
	given := maxDiff(input)

	if given != expected {
		t.Errorf("%d given but %d expected for %d", given, expected, input)
	}
}
