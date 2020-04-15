package main

import (
	"testing"
)

func TestTwoA(t *testing.T) {
	input := "PALAP"
	expected := false
	given := checkRecord(input)

	if given != expected {
		t.Errorf("%v given but %v expected for %s", given, expected, input)
	}
}

func TestMultipleTwoL(t *testing.T) {
	input := "PLLALLP"
	expected := true
	given := checkRecord(input)

	if given != expected {
		t.Errorf("%v given but %v expected for %s", given, expected, input)
	}
}

func TestThreeL(t *testing.T) {
	input := "PLLLAP"
	expected := false
	given := checkRecord(input)

	if given != expected {
		t.Errorf("%v given but %v expected for %s", given, expected, input)
	}
}

func TestEmpty(t *testing.T) {
	input := ""
	expected := true
	given := checkRecord(input)

	if given != expected {
		t.Errorf("%v given but %v expected for %s", given, expected, input)
	}
}

func TestOne(t *testing.T) {
	input := "PPALLP"
	expected := true
	given := checkRecord(input)

	if given != expected {
		t.Errorf("%v given but %v expected for %s", given, expected, input)
	}
}

func TestTwo(t *testing.T) {
	input := "PPALLL"
	expected := false
	given := checkRecord(input)

	if given != expected {
		t.Errorf("%v given but %v expected for %s", given, expected, input)
	}
}
