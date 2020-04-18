package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	arr := []int{1, 3, 4, 8}
	q := [][]int{
		{0, 1},
		{1, 2},
		{0, 3},
		{3, 3},
	}

	expected := []int{2, 7, 14, 8}
	given := xorQueries(arr, q)

	for i, v := range expected {
		if given[i] != v {
			t.Errorf("%v given but %v expected for %v %v", given, expected, arr, q)
		}
	}
}
