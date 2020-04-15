package main

import (
	"testing"
)

func Test_1_4(t *testing.T) {
	x, y := 1, 4
	expected := 2
	given := hammingDistance(x, y)
	if given != expected {
		t.Errorf("%d given but %d expected for [%d, %d]", given, expected, x, y)
	}
}

func Test_1_2(t *testing.T) {
	x, y := 1, 2
	expected := 2
	given := hammingDistance(x, y)
	if given != expected {
		t.Errorf("%d given but %d expected for [%d, %d]", given, expected, x, y)
	}
}

func Test_0_2(t *testing.T) {
	x, y := 0, 2
	expected := 1
	given := hammingDistance(x, y)
	if given != expected {
		t.Errorf("%d given but %d expected for [%d, %d]", given, expected, x, y)
	}
}

func Test_8_8(t *testing.T) {
	x, y := 8, 8
	expected := 0
	given := hammingDistance(x, y)
	if given != expected {
		t.Errorf("%d given but %d expected for [%d, %d]", given, expected, x, y)
	}
}

func Test_5_8(t *testing.T) {
	x, y := 5, 8
	expected := 3
	given := hammingDistance(x, y)
	if given != expected {
		t.Errorf("%d given but %d expected for [%d, %d]", given, expected, x, y)
	}
}
