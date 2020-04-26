package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	expected := []int{1, 4, 2, 7, 5, 3, 8, 6, 9}
	given := findDiagonalOrder(input)

	if len(given) != len(expected) {
		t.Errorf("%v given but %v expected for %v", given, expected, input)
		t.Errorf("Lengths of given and expected are different!")
		return
	}

	for i, v := range given {
		if v != expected[i] {
			t.Errorf("%v given but %v expected for %v", given, expected, input)
			return
		}
	}
}

func TestSecond(t *testing.T) {
	input := [][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}

	expected := []int{1, 6, 2, 8, 7, 3, 9, 4, 12, 10, 5, 13, 11, 14, 15, 16}
	given := findDiagonalOrder(input)

	if len(given) != len(expected) {
		t.Errorf("%v given but %v expected for %v", given, expected, input)
		t.Errorf("Lengths of given and expected are different!")
		return
	}

	for i, v := range given {
		if v != expected[i] {
			t.Errorf("%v given but %v expected for %v", given, expected, input)
			return
		}
	}
}

func TestThird(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4}, {5, 6, 7}, {8}, {9, 10, 11}}
	// 1  2  3
	// 4
	// 5  6  7
	// 8
	// 9  10 11

	expected := []int{1, 4, 2, 5, 3, 8, 6, 9, 7, 10, 11}
	given := findDiagonalOrder(input)

	if len(given) != len(expected) {
		t.Errorf("%v given but %v expected for %v", given, expected, input)
		t.Errorf("Lengths of given and expected are different!")
		return
	}

	for i, v := range given {
		if v != expected[i] {
			t.Errorf("%v given but %v expected for %v", given, expected, input)
			return
		}
	}
}
