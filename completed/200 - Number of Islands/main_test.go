package main

import (
	"testing"
)

func TestFirst(t *testing.T) {
	input := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	intsToBytes(input)
	expected := 1
	given := numIslands(input)

	if given != expected {
		t.Errorf("%d expected but %d given for %v", expected, given, input)
	}
}

func TestSecond(t *testing.T) {
	input := [][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	}
	intsToBytes(input)
	expected := 3
	given := numIslands(input)

	if given != expected {
		t.Errorf("%d expected but %d given for %v", expected, given, input)
	}
}
func TestCheck(t *testing.T) {
	if false {
		t.Errorf("Test of the test")
	}
}
