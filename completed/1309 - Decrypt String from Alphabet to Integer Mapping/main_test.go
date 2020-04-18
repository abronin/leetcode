package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	input := "10#11#12"

	expected := "jkab"
	given := freqAlphabets(input)

	if given != expected {
		t.Errorf("%s given but %s expected for %s", given, expected, input)
	}
}
