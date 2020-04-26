package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 1}
	k := 3

	expected := 12
	given := maxScore(input, k)

	if given != expected {
		t.Errorf("%d given but %d expected for %v, %d", given, expected, input, k)
	}
}

func TestSecond(t *testing.T) {
	input := []int{2, 2, 2}
	k := 2

	expected := 4
	given := maxScore(input, k)

	if given != expected {
		t.Errorf("%d given but %d expected for %v, %d", given, expected, input, k)
	}
}

func TestThird(t *testing.T) {
	input := []int{9, 7, 7, 9, 7, 7, 9}
	k := 7

	expected := 55
	given := maxScore(input, k)

	if given != expected {
		t.Errorf("%d given but %d expected for %v, %d", given, expected, input, k)
	}
}

func TestForth(t *testing.T) {
	input := []int{1, 1000, 1}
	k := 1

	expected := 1
	given := maxScore(input, k)

	if given != expected {
		t.Errorf("%d given but %d expected for %v, %d", given, expected, input, k)
	}
}

func TestFifth(t *testing.T) {
	input := []int{1, 79, 80, 1, 1, 1, 200, 1}
	k := 3

	expected := 202
	given := maxScore(input, k)

	if given != expected {
		t.Errorf("%d given but %d expected for %v, %d", given, expected, input, k)
	}
}
