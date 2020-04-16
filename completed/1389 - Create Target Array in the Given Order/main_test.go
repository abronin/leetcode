package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4}
	index := []int{0, 1, 2, 2, 1}

	expected := []int{0, 4, 1, 3, 2}
	given := createTargetArray(nums, index)

	for i, v := range given {
		if v != expected[i] {
			t.Errorf("%v given but %v expected for %v, %v", given, expected, nums, index)
		}
	}
}

func TestSecond(t *testing.T) {
	nums := []int{1, 2, 3, 4, 0}
	index := []int{0, 1, 2, 3, 0}

	expected := []int{0, 1, 2, 3, 4}
	given := createTargetArray(nums, index)

	for i, v := range given {
		if v != expected[i] {
			t.Errorf("%v given but %v expected for %v, %v", given, expected, nums, index)
		}
	}
}

func TestSimple(t *testing.T) {
	nums := []int{1}
	index := []int{0}

	expected := []int{1}
	given := createTargetArray(nums, index)

	for i, v := range given {
		if v != expected[i] {
			t.Errorf("%v given but %v expected for %v, %v", given, expected, nums, index)
		}
	}
}
