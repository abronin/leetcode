package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	n := 7
	edges := [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}
	time := 1
	target := 7

	expected := float64(1) / float64(3)
	given := frogPosition(n, edges, time, target)

	if given != expected {
		t.Errorf("%f given but %f expected for %d", given, expected, n)
	}
}

func TestSecond(t *testing.T) {
	n := 9
	edges := [][]int{{2, 1}, {3, 1}, {4, 2}, {5, 3}, {6, 5}, {7, 4}, {8, 7}, {9, 7}}
	time := 1
	target := 8

	expected := float64(0)
	given := frogPosition(n, edges, time, target)

	if given != expected {
		t.Errorf("%f given but %f expected for %d", given, expected, n)
	}
}

func TestThird(t *testing.T) {
	n := 18
	edges := [][]int{
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 3},
		{6, 1},
		{7, 1},
		{8, 1},
		{9, 1},
		{10, 9},
		{11, 5},
		{12, 11},
		{13, 12},
		{14, 11},
		{15, 3},
		{16, 2},
		{17, 12},
		{18, 10}}
	time := 14
	target := 17

	expected := float64(0.01042)
	given := frogPosition(n, edges, time, target)

	if given != expected {
		t.Errorf("%f given but %f expected for %d", given, expected, n)
	}
}
