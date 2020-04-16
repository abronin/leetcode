package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	arr1 := []int{4, 5, 8}
	arr2 := []int{10, 9, 1, 8}
	d := 2

	expected := 2
	given := findTheDistanceValue(arr1, arr2, d)

	if given != expected {
		t.Errorf("%d given but %d expected for %v %v, %d", given, expected, arr1, arr2, d)
	}
}

func TestSecond(t *testing.T) {
	arr1 := []int{1, 4, 2, 3}
	arr2 := []int{-4, -3, 6, 10, 20, 30}
	d := 3

	expected := 2
	given := findTheDistanceValue(arr1, arr2, d)

	if given != expected {
		t.Errorf("%d given but %d expected for %v %v, %d", given, expected, arr1, arr2, d)
	}
}
