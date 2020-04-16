package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	lo := 12
	hi := 15
	k := 2

	expected := 13
	given := getKth(lo, hi, k)

	if given != expected {
		t.Errorf("%d given but %d expected for %d, %d, %d", given, expected, lo, hi, k)
	}
}

func TestSec(t *testing.T) {
	lo := 1
	hi := 1
	k := 1

	expected := 1
	given := getKth(lo, hi, k)

	if given != expected {
		t.Errorf("%d given but %d expected for %d, %d, %d", given, expected, lo, hi, k)
	}
}

func TestThird(t *testing.T) {
	lo := 7
	hi := 11
	k := 4

	expected := 7
	given := getKth(lo, hi, k)

	if given != expected {
		t.Errorf("%d given but %d expected for %d, %d, %d", given, expected, lo, hi, k)
	}
}

func TestForth(t *testing.T) {
	lo := 10
	hi := 20
	k := 5

	expected := 13
	given := getKth(lo, hi, k)

	if given != expected {
		t.Errorf("%d given but %d expected for %d, %d, %d", given, expected, lo, hi, k)
	}
}
