package main

import (
	"testing"
)

func TestCorrectOne(t *testing.T) {
	input := "a0b1c2"

	expected := "a0b1c2"
	given := reformat(input)

	if given != expected {
		t.Errorf("%s given but %s expected for %s", given, expected, input)
	}
}

func TestCorrectTwo(t *testing.T) {
	input := "covid2019"

	expected := "c2o0v1i9d"
	given := reformat(input)

	if given != expected {
		t.Errorf("%s given but %s expected for %s", given, expected, input)
	}
}

func TestCorrectThree(t *testing.T) {
	input := "ab123"

	expected := "1a2b3"
	given := reformat(input)

	if given != expected {
		t.Errorf("%s given but %s expected for %s", given, expected, input)
	}
}

func TestDigits(t *testing.T) {
	input := "1229857369"

	expected := ""
	given := reformat(input)

	if given != expected {
		t.Errorf("%s given but %s expected for %s", given, expected, input)
	}
}

func TestLeetcode(t *testing.T) {
	input := "leetcode"

	expected := ""
	given := reformat(input)

	if given != expected {
		t.Errorf("%s given but %s expected for %s", given, expected, input)
	}
}
